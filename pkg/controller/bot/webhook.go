package bot

import (
	"fmt"
	"net/http"

	"github.com/SineChat/bot-ms/internal/models"
	"github.com/SineChat/bot-ms/services/bot"
	"github.com/SineChat/bot-ms/utility"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (base *Controller) HandleWebhook(c *gin.Context) {
	requestBody, err := c.GetRawData()
	if err != nil {
		base.ExtReq.Logger.Error("webhook log error", "Failed to read request body", err.Error())
	}

	response, code, err := bot.HandleWebhookService(c, base.ExtReq, base.Db, requestBody)
	if err != nil {
		rd := utility.BuildErrorResponse(code, "error", err.Error(), err, nil)
		c.JSON(code, rd)
		return
	}

	c.JSON(http.StatusOK, response)

}

func (base *Controller) GetWebhookData(c *gin.Context) {

	var (
		botIdStr = c.Param("bot_id")
	)

	botId, err := primitive.ObjectIDFromHex(botIdStr)
	if err != nil {
		err = fmt.Errorf("invalid bot id:  %v", err.Error())
		rd := utility.BuildErrorResponse(http.StatusBadRequest, "error", err.Error(), err, nil)
		c.JSON(http.StatusBadRequest, rd)
		return
	}

	if models.MyAccessToken == nil {
		msg := "error retrieving authenticated user"
		rd := utility.BuildErrorResponse(http.StatusBadRequest, "error", msg, fmt.Errorf(msg), nil)
		c.JSON(http.StatusBadRequest, rd)
		return
	}

	bot, code, err := bot.GetWebhookDataService(base.ExtReq, base.Db, botId, *models.MyAccessToken)
	if err != nil {
		rd := utility.BuildErrorResponse(code, "error", err.Error(), err, nil)
		c.JSON(code, rd)
		return
	}

	rd := utility.BuildSuccessResponse(http.StatusCreated, "successfully created bot", bot)
	c.JSON(http.StatusCreated, rd)
}
