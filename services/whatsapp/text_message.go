package whatsapp

import "github.com/SineChat/bot-ms/external/external_models"

func (w *Whatsapp) TextMessage(phoneNumber string, text string, previewUrl bool, messageID ...string) *Whatsapp {
	message := external_models.WhatsappSendMessageRequest{
		MessagingProduct: MessagingProduct,
		RecipientType:    IndividualRecipientType,
		To:               phoneNumber,
		Type:             external_models.WhatsAppText,
		Text: &external_models.WhatsappSendMessageRequestTextType{
			PreviewUrl: previewUrl,
			Body:       text,
		},
	}

	if len(messageID) > 0 {
		message.Context = &external_models.WhatsappSendMessageRequestContext{
			MessageID: messageID[0],
		}
	}

	w.Message = message
	return w
}
