package bot

import (
	"fmt"
	"net/http"

	"github.com/SineChat/bot-ms/external/external_models"
	"github.com/SineChat/bot-ms/external/request"
	"github.com/SineChat/bot-ms/internal/models"
	"github.com/SineChat/bot-ms/pkg/repository/storage/mongodb"
	"github.com/SineChat/bot-ms/utility"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateBotService(extReq request.ExternalRequest, db *mongodb.Database, req models.CreateBotReq, accessToken external_models.AccessToken) (models.Bot, int, error) {
	if req.Name == "" {
		req.Name = fmt.Sprintf("BOT_%v", utility.RandomString(12))
	}

	secret := utility.RandomString(20)

	bot := models.Bot{
		AccountID: accessToken.AccountID,
		Name:      req.Name,
		Secret:    secret,
	}

	err := bot.CreateBot(db)
	if err != nil {
		return bot, http.StatusInternalServerError, err
	}

	return bot, http.StatusCreated, nil
}

func GetBotService(extReq request.ExternalRequest, db *mongodb.Database, botID primitive.ObjectID, accessToken external_models.AccessToken) (models.Bot, int, error) {
	var (
		bot = models.Bot{ID: botID}
	)

	err := bot.GetByID(db)
	if err != nil {
		return models.Bot{}, http.StatusInternalServerError, fmt.Errorf("bot not found %v", err.Error())
	}

	if bot.AccountID != accessToken.AccountID {
		return models.Bot{}, http.StatusInternalServerError, fmt.Errorf("access denied")
	}

	if !bot.Action.IsZero() {
		actions, err := getAllActionsFromBase(extReq, db, bot.Action)
		if err != nil {
			return models.Bot{}, http.StatusInternalServerError, err
		}

		bot.Template = actions
	}

	return bot, http.StatusOK, nil
}

func getAllActionsFromBase(extReq request.ExternalRequest, db *mongodb.Database, baseActionID primitive.ObjectID) (models.Action, error) {
	var (
		baseAction = models.Action{ID: baseActionID}
	)

	err := baseAction.GetByID(db)
	if err != nil {
		return baseAction, err
	}

	if len(baseAction.ActionsIDs) > 0 {
		baseAction.Actions, _ = getActionsWithData(extReq, db, baseAction.ActionsIDs)
	}

	return baseAction, nil
}

func getActionsWithData(extReq request.ExternalRequest, db *mongodb.Database, ids map[int]primitive.ObjectID) (map[int]models.Action, error) {
	var (
		actions  = map[int]models.Action{}
		position = 0
	)

	for _, id := range ids {
		action := models.Action{ID: id}
		err := action.GetByID(db)
		if err != nil {
			extReq.Logger.Error(fmt.Sprintf("error getting action with id %v: Err %v", id, err.Error()))
			continue
		}

		if len(action.ActionsIDs) > 0 {
			action.Actions, _ = getActionsWithData(extReq, db, action.ActionsIDs)
		}
		actions[position] = action
		position += 1
	}

	return actions, nil
}

func GetBotConnectionsService(extReq request.ExternalRequest, db *mongodb.Database, botID primitive.ObjectID, accessToken external_models.AccessToken, paginator mongodb.Pagination) ([]models.Connection, mongodb.PaginationResponse, int, error) {
	var (
		bot = models.Bot{ID: botID}
	)

	err := bot.GetByID(db)
	if err != nil {
		return []models.Connection{}, mongodb.PaginationResponse{}, http.StatusInternalServerError, fmt.Errorf("bot not found %v", err.Error())
	}

	if bot.AccountID != accessToken.AccountID {
		return []models.Connection{}, mongodb.PaginationResponse{}, http.StatusInternalServerError, fmt.Errorf("access denied")
	}

	connection := models.Connection{BotID: botID}
	connections, pagination, err := connection.GetConnections(db, paginator)
	if err != nil {
		return []models.Connection{}, mongodb.PaginationResponse{}, http.StatusInternalServerError, err
	}

	return connections, pagination, http.StatusOK, nil
}

// get all bots service
func GetAllBotsService(extReq request.ExternalRequest, db *mongodb.Database, accessToken external_models.AccessToken, paginator mongodb.Pagination) ([]models.Bot, mongodb.PaginationResponse, int, error) {

	bot := models.Bot{}
	bots, pagination, err := bot.GetAllBots(db, paginator)
	if err != nil {
		return nil, mongodb.PaginationResponse{}, http.StatusInternalServerError, err
	}
	return bots, pagination, http.StatusOK, nil
}
