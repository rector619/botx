package bot

import (
	"fmt"
	"net/http"

	"github.com/SineChat/bot-ms/external/external_models"
	"github.com/SineChat/bot-ms/external/request"
	"github.com/SineChat/bot-ms/internal/models"
	"github.com/SineChat/bot-ms/pkg/repository/storage/mongodb"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddActionService(extReq request.ExternalRequest, db *mongodb.Database, req models.AddActionReq, accessToken external_models.AccessToken) (models.Action, int, error) {
	var (
		bot             = models.Bot{ID: req.BotID}
		referenceAction = models.Action{ID: req.ActionID}
	)

	err := bot.GetByID(db)
	if err != nil {
		return models.Action{}, http.StatusInternalServerError, fmt.Errorf("bot not found %v", err.Error())
	}

	if !req.ActionID.IsZero() && bot.Action.IsZero() {
		return models.Action{}, http.StatusBadRequest, fmt.Errorf("this bot has no base  action you have to create that first. you can do this by not speciying action_id")
	}

	if req.ActionID.IsZero() && !bot.Action.IsZero() {
		return models.Action{}, http.StatusBadRequest, fmt.Errorf("this bot already has a base  action, you can  edit or remove it, or provide an action_id for this action")
	}

	if !req.ActionID.IsZero() && req.Position == nil {
		return models.Action{}, http.StatusBadRequest, fmt.Errorf("position must be specified when action id is provided")
	}

	if req.ActionID.IsZero() && req.DefaultResponse == "" {
		return models.Action{}, http.StatusBadRequest, fmt.Errorf("you have to provide default_response for base action")
	}

	if !req.ActionID.IsZero() {
		err := referenceAction.GetByID(db)
		if err != nil {
			return models.Action{}, http.StatusInternalServerError, fmt.Errorf("action not found %v", err.Error())
		}

		if referenceAction.BotID != bot.ID {
			return models.Action{}, http.StatusInternalServerError, fmt.Errorf("the provided action_id is not  related to the  bot_id")
		}
	}

	if err := ValidateCreateActionTypes(req); err != nil {
		return models.Action{}, http.StatusBadRequest, err
	}

	action := models.Action{
		Type:            req.Type,
		BotID:           bot.ID,
		Text:            req.Text,
		DefaultResponse: req.DefaultResponse,
	}

	if !req.ActionID.IsZero() {
		action.ActionID = req.ActionID
	}

	err = action.CreateAction(db)
	if err != nil {
		return models.Action{}, http.StatusInternalServerError, err
	}

	if req.ActionID.IsZero() && bot.Action.IsZero() {
		bot.Action = action.ID
		err := bot.UpdateAll(db)
		if err != nil {
			return models.Action{}, http.StatusInternalServerError, err
		}
	}

	if !req.ActionID.IsZero() && !bot.Action.IsZero() {
		referenceActionsIDsSlice := []primitive.ObjectID{}
		newActionsIDsMap := map[int]primitive.ObjectID{}
		currentPosition, placed := -1, false

		for i := 0; i < len(referenceAction.ActionsIDs); i++ {
			id, ok := referenceAction.ActionsIDs[i]
			if ok {
				referenceActionsIDsSlice = append(referenceActionsIDsSlice, id)
			}
		}

		for i := 0; i <= len(referenceActionsIDsSlice); i++ {
			currentPosition += 1
			if currentPosition == *req.Position && !placed {
				newActionsIDsMap[currentPosition], placed = action.ID, true
				currentPosition += 1
			} else if i == len(referenceActionsIDsSlice) && !placed {
				newActionsIDsMap[currentPosition], placed = action.ID, true
			}

			if i < len(referenceActionsIDsSlice) {
				newActionsIDsMap[currentPosition] = referenceActionsIDsSlice[i]
			}
		}

		referenceAction.ActionsIDs = newActionsIDsMap
		err := referenceAction.UpdateAll(db)
		if err != nil {
			return models.Action{}, http.StatusInternalServerError, err
		}
	}

	return action, http.StatusOK, nil
}

// delete action using action id as a parameter

func DeleteActionService(extReq request.ExternalRequest, db *mongodb.Database, actionID primitive.ObjectID, accessToken external_models.AccessToken, isFirstLayer bool) (int, error) {
	action := models.Action{ID: actionID}
	err := action.GetByID(db)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("action not found %v", err.Error())
	}

	bot := models.Bot{ID: action.BotID}
	err = bot.GetByID(db)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("bot not found %v", err.Error())
	}

	subActions := action.ActionsIDs

	// Skip this step if it is not the first layer
	// STEP: 1
	// check if the action_id field is set.
	// if it is, remove the actionID from the referenced Action ActionID's map

	if isFirstLayer {
		if !action.ActionID.IsZero() {
			referenceAction := models.Action{ID: action.ActionID}
			err := referenceAction.GetByID(db)
			if err != nil {
				return http.StatusInternalServerError, fmt.Errorf("referenceAction action not found %v", err.Error())
			}
			currentPosition := 0
			actionIDs := map[int]primitive.ObjectID{}

			//  loop through referenceAction.ActionsIDs and recreate the map without action.ID
			for _, action_id := range referenceAction.ActionsIDs {
				if action_id != action.ID {
					actionIDs[currentPosition] = action_id
					currentPosition = currentPosition + 1
				}
			}
			referenceAction.ActionsIDs = actionIDs

			// update the referenceAction.ActionIDs map
			err = referenceAction.UpdateAll(db)
			if err != nil {
				return http.StatusInternalServerError, fmt.Errorf("error updating action %v", err.Error())
			}
		} else {
			// if the action.ActionID is empty, that means it is a base action, then remove action from bot
			bot.Action = primitive.NilObjectID
			err := bot.UpdateAll(db)
			if err != nil {
				return http.StatusInternalServerError, fmt.Errorf("error updating bot %v", err.Error())
			}
		}

	}

	// STEP: 2
	// check if subActions map has data in it
	if len(subActions) > 0 {

		var actionIDsSlice []primitive.ObjectID // a slice to hold the actionID's from the action.ActionIDs map
		for _, action_id := range subActions {
			actionIDsSlice = append(actionIDsSlice, action_id)
		}

		// delete all the actions
		err = action.DeleteActionMany(db, actionIDsSlice)
		if err != nil {
			return http.StatusInternalServerError, fmt.Errorf("error deleting actions %v", err.Error())
		}
	}

	// STEP: 3
	// check it the action has a bot linked to it
	// if it is, set bot.Action to zero since we want to delete the action linked
	if !action.BotID.IsZero() {
		// get the bot
		bot := models.Bot{ID: action.BotID}
		err = bot.GetByID(db)
		if err != nil {
			return http.StatusInternalServerError, err
		}

		// set it to objectID(000000000000000000000000)
		bot.Action = primitive.NilObjectID
		// update
		err = bot.UpdateAll(db)
		if err != nil {
			return http.StatusInternalServerError, err
		}
	}

	// STEP: 4
	// finally, delete the action from the database じゃね
	err = action.DeleteAction(db)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("error deleting action %v", err.Error())

	}

	return http.StatusOK, nil
}
