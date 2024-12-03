package bot

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/SineChat/bot-ms/external/external_models"
	"github.com/SineChat/bot-ms/external/request"
	"github.com/SineChat/bot-ms/internal/config"
	"github.com/SineChat/bot-ms/internal/models"
	"github.com/SineChat/bot-ms/pkg/repository/storage/mongodb"
	"github.com/SineChat/bot-ms/services/whatsapp"
	"github.com/SineChat/bot-ms/utility"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func HandleWebhookService(c *gin.Context, extReq request.ExternalRequest, db *mongodb.Database, requestBody []byte) (interface{}, int, error) {

	resp, integration, err := GetPlatform(c, db, requestBody)
	if err != nil {
		extReq.Logger.Error(fmt.Sprintf("webhook log error for %v %v", integration, err.Error()))
		return resp, http.StatusInternalServerError, err
	}

	logWebhookData(extReq, db, c.Param("bot_id"), integration, requestBody)

	switch integration {
	case models.WhatsApp:
		err = whatsapp.HandleWebhookService(c, extReq, db, requestBody)
	default:
		err = fmt.Errorf("%v is not implemented", integration)
	}

	if err != nil {
		extReq.Logger.Error(fmt.Sprintf("webhook log error for %v %v", integration, err.Error()))
		return resp, http.StatusInternalServerError, err
	}

	return resp, http.StatusOK, nil
}

func GetPlatform(c *gin.Context, db *mongodb.Database, requestBody []byte) (interface{}, models.IntegrationType, error) {
	var (
		mapData = map[string]interface{}{}
	)

	if len(requestBody) > 0 {
		err := json.Unmarshal(requestBody, &mapData)
		if err != nil {
			return nil, models.IntegrationType("unknown"), err
		}
	}

	whatsappObject, whatsappObjectOk := mapData["object"]

	if c.Query("hub.challenge") != "" && c.Query("hub.mode") != "" && c.Query("hub.verify_token") != "" {
		challenge, _ := strconv.Atoi((c.Query("hub.challenge")))
		botId, err := primitive.ObjectIDFromHex(c.Param("bot_id"))
		if err != nil {
			err = fmt.Errorf("invalid bot id:  %v", err.Error())
			rd := utility.BuildErrorResponse(http.StatusBadRequest, "error", err.Error(), err, nil)
			c.JSON(http.StatusBadRequest, rd)
			return nil, models.WhatsApp, err
		}

		bot := models.Bot{ID: botId}
		err = bot.GetByID(db)
		if err != nil {
			return nil, models.WhatsApp, err
		}

		if bot.Secret != c.Query("hub.verify_token") {
			return nil, models.WhatsApp, fmt.Errorf("access denied")
		}

		return challenge, models.WhatsApp, nil
	} else if whatsappObjectStr, ok := whatsappObject.(string); whatsappObjectOk && ok && strings.EqualFold(whatsappObjectStr, "whatsapp_business_account") {
		return nil, models.WhatsApp, nil
	} else {
		return nil, models.IntegrationType("not implemented"), nil
	}

}

func logWebhookData(extReq request.ExternalRequest, db *mongodb.Database, botID string, integration models.IntegrationType, requestBody []byte) error {
	botId, _ := primitive.ObjectIDFromHex(botID)
	extReq.Logger.Info(fmt.Sprintf("webhook log info for %v %v", integration, string(requestBody)))
	webhookLog := models.WebhookLog{
		BotID:    botId,
		Platform: integration,
		Payload:  string(requestBody),
	}
	err := webhookLog.CreateWebhookLog(db)
	if err != nil {
		return err
	}
	return nil
}

func GetWebhookDataService(extReq request.ExternalRequest, db *mongodb.Database, botId primitive.ObjectID, accessToken external_models.AccessToken) (models.BotWebhookData, int, error) {
	var (
		bot    = models.Bot{ID: botId}
		appUrl = config.GetConfig().App.Url
	)

	err := bot.GetByID(db)
	if err != nil {
		return models.BotWebhookData{}, http.StatusBadRequest, fmt.Errorf("bot not found %v", err.Error())
	}

	if bot.Secret == "" {
		bot.Secret = utility.RandomString(20)
		err = bot.UpdateAll(db)
		if err != nil {
			return models.BotWebhookData{}, http.StatusInternalServerError, err
		}
	}

	if accessToken.AccountID != bot.AccountID {
		return models.BotWebhookData{}, http.StatusBadRequest, fmt.Errorf("access denied")
	}

	return models.BotWebhookData{
		Secret:     bot.Secret,
		WebhookUrl: fmt.Sprintf("%v/v1/webhook/%v", appUrl, bot.ID.Hex()),
	}, http.StatusOK, nil
}
