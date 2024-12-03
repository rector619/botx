package bot

import (
	"fmt"
	"net/http"

	"github.com/SineChat/bot-ms/internal/models"
	"github.com/SineChat/bot-ms/pkg/repository/storage/mongodb"
	"github.com/SineChat/bot-ms/services/bot"
	"github.com/SineChat/bot-ms/utility"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (base *Controller) CreateBot(c *gin.Context) {

	var (
		req = models.CreateBotReq{}
	)

	err := c.ShouldBind(&req)
	if err != nil {
		rd := utility.BuildErrorResponse(http.StatusBadRequest, "error", "Failed to parse request body", err, nil)
		c.JSON(http.StatusBadRequest, rd)
		return
	}

	err = base.Validator.Struct(&req)
	if err != nil {
		rd := utility.BuildErrorResponse(http.StatusBadRequest, "error", "Validation failed", utility.ValidationResponse(err, base.Validator), nil)
		c.JSON(http.StatusBadRequest, rd)
		return
	}

	vr := mongodb.ValidateRequestM{Logger: base.Logger, Test: false}
	err = vr.ValidateRequest(req)
	if err != nil {
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

	bot, code, err := bot.CreateBotService(base.ExtReq, base.Db, req, *models.MyAccessToken)
	if err != nil {
		rd := utility.BuildErrorResponse(code, "error", err.Error(), err, nil)
		c.JSON(code, rd)
		return
	}

	rd := utility.BuildSuccessResponse(http.StatusCreated, "successfully created bot", bot)
	c.JSON(http.StatusCreated, rd)
}

func (base *Controller) GetBot(c *gin.Context) {

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

	bot, code, err := bot.GetBotService(base.ExtReq, base.Db, botId, *models.MyAccessToken)
	if err != nil {
		rd := utility.BuildErrorResponse(code, "error", err.Error(), err, nil)
		c.JSON(code, rd)
		return
	}

	rd := utility.BuildSuccessResponse(http.StatusOK, "successfully created bot", bot)
	c.JSON(http.StatusOK, rd)
}

func (base *Controller) GetBotConnections(c *gin.Context) {

	var (
		botIdStr  = c.Param("bot_id")
		paginator = mongodb.GetPagination(c)
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

	connections, pagination, code, err := bot.GetBotConnectionsService(base.ExtReq, base.Db, botId, *models.MyAccessToken, paginator)
	if err != nil {
		rd := utility.BuildErrorResponse(code, "error", err.Error(), err, nil)
		c.JSON(code, rd)
		return
	}

	rd := utility.BuildSuccessResponse(http.StatusOK, "successful", connections, pagination)
	c.JSON(http.StatusOK, rd)
}

// get all bots
func (base *Controller) GetAllBots(c *gin.Context) {

	var (
		paginator = mongodb.GetPagination(c)
	)

	if models.MyAccessToken == nil {
		msg := "error retrieving authenticated user"
		rd := utility.BuildErrorResponse(http.StatusBadRequest, "error", msg, fmt.Errorf(msg), nil)
		c.JSON(http.StatusBadRequest, rd)
		return
	}

	bots, pagination, code, err := bot.GetAllBotsService(base.ExtReq, base.Db, *models.MyAccessToken, paginator)
	if err != nil {
		rd := utility.BuildErrorResponse(code, "error", err.Error(), err, nil)
		c.JSON(code, rd)
		return
	}

	rd := utility.BuildSuccessResponse(http.StatusOK, "successful", bots, pagination)
	c.JSON(http.StatusOK, rd)
}
