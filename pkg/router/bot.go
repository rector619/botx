package router

import (
	"fmt"

	"github.com/SineChat/bot-ms/external/request"
	"github.com/SineChat/bot-ms/pkg/controller/bot"
	"github.com/SineChat/bot-ms/pkg/middleware"
	"github.com/SineChat/bot-ms/pkg/repository/storage/mongodb"
	"github.com/SineChat/bot-ms/utility"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func Bot(r *gin.Engine, ApiVersion string, validator *validator.Validate, db *mongodb.Database, logger *utility.Logger) *gin.Engine {
	extReq := request.ExternalRequest{Logger: logger, Test: false}
	botC := bot.Controller{Db: db, Validator: validator, Logger: logger, ExtReq: extReq}

	botUrl := r.Group(fmt.Sprintf("%v", ApiVersion))
	{
		botUrl.GET("/webhook/:bot_id", botC.HandleWebhook)
		botUrl.POST("/webhook/:bot_id", botC.HandleWebhook)
	}

	botApiUrl := r.Group(fmt.Sprintf("%v/bot", ApiVersion), middleware.Authorize(db, extReq, middleware.ApiPrivateType, middleware.ApiPublicType))
	{
		botApiUrl.POST("/create", botC.CreateBot)
		botApiUrl.GET("/get/:bot_id", botC.GetBot)
		botApiUrl.GET("/get-all", botC.GetAllBots) // paginated
		botApiUrl.GET("/get-webhook-data/:bot_id", botC.GetWebhookData)
		botApiUrl.GET("/get/:bot_id/connections", botC.GetBotConnections)
		botApiUrl.POST("/deploy-to-connection", botC.DeployBotToConnection)
		botApiUrl.POST("/remove-from-connection", botC.RemoveBotFromConnection)
	}

	connectionApiUrl := r.Group(fmt.Sprintf("%v/connection", ApiVersion), middleware.Authorize(db, extReq, middleware.ApiPrivateType, middleware.ApiPublicType))
	{
		connectionApiUrl.POST("/add", botC.ConnectPlatform)
		connectionApiUrl.PATCH("/update/:connection_id", botC.UpdatePlatformConnection)
	}

	actionApiUrl := r.Group(fmt.Sprintf("%v/bot/action", ApiVersion), middleware.Authorize(db, extReq, middleware.ApiPrivateType, middleware.ApiPublicType))
	{
		actionApiUrl.POST("/add", botC.AddAction)
		// actionApiUrl.PATCH("/move", botC.MoveAction)
		// actionApiUrl.PATCH("/update", botC.UpdateAction)
		actionApiUrl.DELETE("/delete/:action_id", botC.DeleteAction)
	}

	return r
}
