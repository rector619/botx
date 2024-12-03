package whatsapp

import "github.com/SineChat/bot-ms/external/external_models"

func (w *Whatsapp) LocationMessage(phoneNumber string, longitude, latitude float64, name, address string) *Whatsapp {
	message := external_models.WhatsappSendMessageRequest{
		MessagingProduct: MessagingProduct,
		RecipientType:    IndividualRecipientType,
		To:               phoneNumber,
		Type:             external_models.WhatsAppLocation,
		Location: &external_models.WhatsappSendMessageRequestLocationType{
			Longitude: longitude,
			Latitude:  latitude,
			Name:      name,
			Address:   address,
		},
	}

	w.Message = message
	return w
}
