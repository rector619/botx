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

func DeployBotToConnectionService(extReq request.ExternalRequest, db *mongodb.Database, req models.DeployBotToConnectionReq, accessToken external_models.AccessToken) (int, error) {
	var (
		bot  = models.Bot{ID: req.BotID}
		conn = models.Connection{ID: req.ConnectionID}
	)

	err := bot.GetByID(db)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	err = conn.GetByID(db)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	if bot.AccountID != conn.AccountID || bot.AccountID != accessToken.AccountID || conn.AccountID != accessToken.AccountID {
		return http.StatusBadRequest, fmt.Errorf("access denied")
	}

	if bot.Action.IsZero() {
		return http.StatusBadRequest, fmt.Errorf("cannot deploy an empty bot")
	}

	if !conn.BotID.IsZero() {
		return http.StatusBadRequest, fmt.Errorf("already connected to a bot")
	}

	conn.BotID = bot.ID
	conn.IsDeployed = true
	err = conn.UpdateAll(db)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func RemoveBotFromConnectionService(extReq request.ExternalRequest, db *mongodb.Database, req models.DeployBotToConnectionReq, accessToken external_models.AccessToken) (int, error) {
	var (
		bot  = models.Bot{ID: req.BotID}
		conn = models.Connection{ID: req.ConnectionID}
	)

	err := bot.GetByID(db)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	err = conn.GetByID(db)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	if bot.AccountID != conn.AccountID || bot.AccountID != accessToken.AccountID || conn.AccountID != accessToken.AccountID {
		return http.StatusBadRequest, fmt.Errorf("access denied")
	}

	if conn.BotID != bot.ID {
		return http.StatusBadRequest, fmt.Errorf("the connection is not deployed for this bot")
	}

	conn.BotID = primitive.NilObjectID
	conn.IsDeployed = false
	err = conn.UpdateAll(db)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
