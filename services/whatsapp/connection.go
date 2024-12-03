package whatsapp

import (
	"fmt"
	"net/http"

	"github.com/SineChat/bot-ms/external/external_models"
	"github.com/SineChat/bot-ms/internal/models"
	"github.com/SineChat/bot-ms/pkg/repository/storage/mongodb"
)

func ConnectWhatsappService(db *mongodb.Database, req models.CreateConnectionReq, accessToken external_models.AccessToken) (models.Connection, int, error) {
	if req.NumberID == 0 {
		return models.Connection{}, http.StatusBadRequest, fmt.Errorf("provide phone number id from whatsapp in number_id")
	}

	if req.PlatformToken == "" {
		return models.Connection{}, http.StatusBadRequest, fmt.Errorf("provide authentication token from whatsapp in platform_token")
	}

	connection := models.Connection{
		NumberID:      req.NumberID,
		Type:          models.WhatsApp,
		PlatformToken: req.PlatformToken,
		MarkAsRead:    req.MarkAsRead,
		AccountID:     accessToken.AccountID,
		IsDeployed:    false,
	}

	err := connection.GetByNumberID(db)
	if err == nil {
		return models.Connection{}, http.StatusBadRequest, fmt.Errorf("number id provided is in use by another connection")
	}

	err = connection.CreateConnection(db)
	if err != nil {
		return models.Connection{}, http.StatusInternalServerError, err
	}

	return connection, http.StatusOK, nil
}

func UpdateWhatsappConnectionService(db *mongodb.Database, conn models.Connection, req models.UpdateConnectionReq) (models.Connection, int, error) {
	if req.NumberID != 0 && req.NumberID != conn.NumberID {
		conn.NumberID = req.NumberID
		err := conn.GetByNumberID(db)
		if err == nil {
			return models.Connection{}, http.StatusBadRequest, fmt.Errorf("number id provided is in use by another connection")
		}
	}

	if req.PlatformToken != "" {
		conn.PlatformToken = req.PlatformToken
	}

	if req.MarkAsRead != nil {
		conn.MarkAsRead = *req.MarkAsRead
	}

	err := conn.UpdateAll(db)
	if err != nil {
		return models.Connection{}, http.StatusInternalServerError, err
	}

	return conn, http.StatusOK, nil
}
