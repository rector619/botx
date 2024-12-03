package whatsapp

import (
	"fmt"

	"github.com/SineChat/bot-ms/external/external_models"
	"github.com/SineChat/bot-ms/external/request"
	"github.com/SineChat/bot-ms/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	version                 = "17.0"
	MessagingProduct        = "whatsapp"
	IndividualRecipientType = "individual"
)

type Whatsapp struct {
	ID          primitive.ObjectID                         `bson:"_id,omitempty" json:"id"`
	NumberID    int                                        `bson:"number_id" json:"number_id"`           //Your phone number id provided by WhatsApp cloud
	Type        models.IntegrationType                     `bson:"type" json:"type"`                     // type of bot
	AccessToken string                                     `bson:"platform_token" json:"platform_token"` // plaform accesstoken
	MarkAsRead  bool                                       `bson:"mark_as_read" json:"mark_as_read"`     // Use to set whether incoming messages should be marked as read Default is True
	BaseUrl     string                                     `bson:"base_url" json:"base_url"`
	MessageUrl  string                                     `bson:"message_url" json:"message_url"`
	MediaUrl    string                                     `bson:"media_url" json:"media_url"`
	Message     external_models.WhatsappSendMessageRequest `bson:"message" json:"message"`
	ExtReq      request.ExternalRequest                    `bson:"-" json:"-"`
	Error       error                                      `bson:"-" json:"-"`
}

func Init(bot models.Bot, conn models.Connection, extReq request.ExternalRequest) (*Whatsapp, error) {
	if conn.Type != models.WhatsApp {
		return nil, fmt.Errorf("this is not a whatsapp connection")
	}

	baseUrl := fmt.Sprintf("https://graph.facebook.com/v%v/%v", version, conn.NumberID)

	return &Whatsapp{
		ID:          conn.ID,
		NumberID:    conn.NumberID,
		Type:        conn.Type,
		AccessToken: conn.PlatformToken,
		MarkAsRead:  conn.MarkAsRead,
		BaseUrl:     baseUrl,
		MessageUrl:  fmt.Sprintf("%v/messages", baseUrl),
		MediaUrl:    fmt.Sprintf("%v/media", baseUrl),
		ExtReq:      extReq,
	}, nil
}

func (w *Whatsapp) SendMessage() (external_models.WhatsappSendMessageResponse, error) {
	if w.Error != nil {
		return external_models.WhatsappSendMessageResponse{}, w.Error
	}

	respBodyInterface, err := w.ExtReq.SendExternalRequest(request.WhatsAppSendMessage, external_models.WhatsappSendMessagePreRequest{
		Url:                w.MessageUrl,
		Token:              w.AccessToken,
		SendMessageRequest: w.Message,
	})
	if err != nil {
		return external_models.WhatsappSendMessageResponse{}, err
	}

	respBody, ok := respBodyInterface.(external_models.WhatsappSendMessageResponse)
	if !ok {
		return external_models.WhatsappSendMessageResponse{}, fmt.Errorf("invalid response interface")
	}

	return respBody, nil
}
