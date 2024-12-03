package bot

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"

	"github.com/SineChat/bot-ms/internal/models"
	"github.com/SineChat/bot-ms/pkg/repository/storage/mongodb"
	"github.com/SineChat/bot-ms/services/bot"
	"github.com/SineChat/bot-ms/utility"
	"github.com/gin-gonic/gin"
)

func (base *Controller) AddAction(c *gin.Context) {
	req := models.AddActionReq{}

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

	action, code, err := bot.AddActionService(base.ExtReq, base.Db, req, *models.MyAccessToken)
	if err != nil {
		rd := utility.BuildErrorResponse(code, "error", err.Error(), err, nil)
		c.JSON(code, rd)
		return
	}

	rd := utility.BuildSuccessResponse(http.StatusOK, "successful", action)
	c.JSON(http.StatusOK, rd)
}

// delete action using action id as a parameter
func (base *Controller) DeleteAction(c *gin.Context) {

	actionIDStr := c.Param("action_id")

	actionID, err := primitive.ObjectIDFromHex(actionIDStr)
	if err != nil {
		err = fmt.Errorf("invalid action id:  %v", err.Error())
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

	code, err := bot.DeleteActionService(base.ExtReq, base.Db, actionID, *models.MyAccessToken, true)
	if err != nil {
		rd := utility.BuildErrorResponse(code, "error", err.Error(), err, nil)
		c.JSON(code, rd)
		return
	}

	rd := utility.BuildSuccessResponse(http.StatusOK, "successful", nil)
	c.JSON(http.StatusOK, rd)
}
