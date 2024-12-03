package whatsapp

import (
	"fmt"

	"github.com/SineChat/bot-ms/external/external_models"
)

func (r *RequestObj) SendWhatsappMessage() (external_models.WhatsappSendMessageResponse, error) {

	var (
		outBoundResponse external_models.WhatsappSendMessageResponse
		logger           = r.Logger
		idata            = r.RequestData
	)

	data, ok := idata.(external_models.WhatsappSendMessagePreRequest)
	if !ok {
		logger.Error("send whatsapp messahe", idata, "request data format error")
		return outBoundResponse, fmt.Errorf("request data format error")
	}

	headers := map[string]string{
		"Authorization": fmt.Sprintf("Bearer %s", data.Token),
		"Content-Type":  "application/json",
	}

	logger.Info("send whatsapp messahe", data)
	err := r.getNewSendRequestObject(data.SendMessageRequest, headers, data.Url).SendRequest(&outBoundResponse)
	if err != nil {
		logger.Error("send whatsapp messahe", outBoundResponse, err.Error())
		return outBoundResponse, err
	}

	return outBoundResponse, nil
}
