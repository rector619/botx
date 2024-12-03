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

func (base *Controller) ConnectPlatform(c *gin.Context) {

	var (
		req = models.CreateConnectionReq{}
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

	connection, code, err := bot.ConnectPlatformService(base.ExtReq, base.Db, req, *models.MyAccessToken)
	if err != nil {
		rd := utility.BuildErrorResponse(code, "error", err.Error(), err, nil)
		c.JSON(code, rd)
		return
	}

	rd := utility.BuildSuccessResponse(http.StatusCreated, "successfully created connection", connection)
	c.JSON(http.StatusCreated, rd)
}

func (base *Controller) UpdatePlatformConnection(c *gin.Context) {

	var (
		req             = models.UpdateConnectionReq{}
		connectionIdStr = c.Param("connection_id")
	)

	connectionId, err := primitive.ObjectIDFromHex(connectionIdStr)
	if err != nil {
		err = fmt.Errorf("invalid connection id:  %v", err.Error())
		rd := utility.BuildErrorResponse(http.StatusBadRequest, "error", err.Error(), err, nil)
		c.JSON(http.StatusBadRequest, rd)
		return
	}

	err = c.ShouldBind(&req)
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

	connection, code, err := bot.UpdatePlatformConnectionService(base.ExtReq, base.Db, connectionId, req, *models.MyAccessToken)
	if err != nil {
		rd := utility.BuildErrorResponse(code, "error", err.Error(), err, nil)
		c.JSON(code, rd)
		return
	}

	rd := utility.BuildSuccessResponse(http.StatusOK, "successfully updated connection", connection)
	c.JSON(http.StatusOK, rd)
}
