package whatsapp

import "github.com/SineChat/bot-ms/external/external_models"

func (w *Whatsapp) ContactMessage(phoneNumber string, contacts []external_models.WhatsappSendMessageRequestContactType) *Whatsapp {
	message := external_models.WhatsappSendMessageRequest{
		MessagingProduct: MessagingProduct,
		RecipientType:    IndividualRecipientType,
		To:               phoneNumber,
		Type:             external_models.WhatsAppContacts,
		Contacts:         &contacts,
	}

	w.Message = message
	return w
}
