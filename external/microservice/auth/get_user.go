package auth

import (
	"fmt"

	"github.com/SineChat/bot-ms/external/external_models"
	"github.com/SineChat/bot-ms/internal/config"
)

func (r *RequestObj) GetUser() (external_models.User, error) {

	var (
		appKey           = config.GetConfig().App.Key
		outBoundResponse external_models.GetUserResponse
		logger           = r.Logger
		idata            = r.RequestData
	)

	data, ok := idata.(external_models.GetUserRequestModel)
	if !ok {
		logger.Error("get user", idata, "request data format error")
		return outBoundResponse.Data.User, fmt.Errorf("request data format error")
	}

	headers := map[string]string{
		"Content-Type": "application/json",
		"app-key":      appKey,
	}

	logger.Info("get user", data)
	err := r.getNewSendRequestObject(data, headers, "").SendRequest(&outBoundResponse)
	if err != nil {
		logger.Error("get user", outBoundResponse, err.Error())
		return outBoundResponse.Data.User, err
	}
	logger.Info("get user", outBoundResponse)

	return outBoundResponse.Data.User, nil
}
