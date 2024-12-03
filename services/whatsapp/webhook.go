package whatsapp

import (
	"encoding/json"
	"fmt"

	"github.com/SineChat/bot-ms/external/request"
	"github.com/SineChat/bot-ms/pkg/repository/storage/mongodb"
	"github.com/gin-gonic/gin"
)

func HandleWebhookService(c *gin.Context, extReq request.ExternalRequest, db *mongodb.Database, requestBody []byte) error {
	if len(requestBody) < 1 {
		return nil
	}

	var data WhatsAppNotification
	err := json.Unmarshal(requestBody, &data)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}

	fmt.Printf("%+v\n", data)
	return nil
}
