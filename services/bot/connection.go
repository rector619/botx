package bot

import (
	"fmt"
	"net/http"

	"github.com/SineChat/bot-ms/external/external_models"
	"github.com/SineChat/bot-ms/external/request"
	"github.com/SineChat/bot-ms/internal/models"
	"github.com/SineChat/bot-ms/pkg/repository/storage/mongodb"
	"github.com/SineChat/bot-ms/services/whatsapp"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ConnectPlatformService(extReq request.ExternalRequest, db *mongodb.Database, req models.CreateConnectionReq, accessToken external_models.AccessToken) (models.Connection, int, error) {

	switch req.Type {
	case models.WhatsApp:
		return whatsapp.ConnectWhatsappService(db, req, accessToken)
	default:
		return models.Connection{}, http.StatusNotImplemented, fmt.Errorf("create for type %v is not implemented", req.Type)
	}
}

func UpdatePlatformConnectionService(extReq request.ExternalRequest, db *mongodb.Database, connectionId primitive.ObjectID, req models.UpdateConnectionReq, accessToken external_models.AccessToken) (models.Connection, int, error) {
	var (
		conn = models.Connection{ID: connectionId}
	)

	err := conn.GetByID(db)
	if err != nil {
		return models.Connection{}, http.StatusBadRequest, fmt.Errorf("connection not found %v", err.Error())
	}

	if conn.AccountID != accessToken.AccountID {
		return models.Connection{}, http.StatusBadRequest, fmt.Errorf("access denied")
	}

	switch conn.Type {
	case models.WhatsApp:
		return whatsapp.UpdateWhatsappConnectionService(db, conn, req)
	default:
		return models.Connection{}, http.StatusNotImplemented, fmt.Errorf("update for type %v is not implemented", conn.Type)
	}
}
