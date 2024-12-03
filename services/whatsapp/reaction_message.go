package whatsapp

import "github.com/SineChat/bot-ms/external/external_models"

func (w *Whatsapp) ReactionMessage(phoneNumber string, emoji string, messageID string) *Whatsapp {
	message := external_models.WhatsappSendMessageRequest{
		MessagingProduct: MessagingProduct,
		RecipientType:    IndividualRecipientType,
		To:               phoneNumber,
		Type:             external_models.WhatsAppReaction,
		Reaction: &external_models.WhatsappSendMessageRequestReactionType{
			MessageID: messageID,
			Emoji:     emoji,
		},
	}

	w.Message = message
	return w
}
