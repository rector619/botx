package whatsapp

import (
	"fmt"

	"github.com/SineChat/bot-ms/external/external_models"
)

type WhatsAppMediaType string
type WhatsAppFileReferenceType string

var (
	WhatsAppImage    WhatsAppMediaType = "image"
	WhatsAppAudio    WhatsAppMediaType = "audio"
	WhatsAppDocument WhatsAppMediaType = "document"
	WhatsAppSticker  WhatsAppMediaType = "sticker"
	WhatsAppVideo    WhatsAppMediaType = "video"

	WhatsappFileLink WhatsAppFileReferenceType = "link"
	WhatsappFileID   WhatsAppFileReferenceType = "id"
)

func (w *Whatsapp) MediaMessage(phoneNumber string, mediaType WhatsAppMediaType, media string, fileReferenceType WhatsAppFileReferenceType) *Whatsapp {
	message := external_models.WhatsappSendMessageRequest{
		MessagingProduct: MessagingProduct,
		RecipientType:    IndividualRecipientType,
		To:               phoneNumber,
	}

	mediaData := external_models.WhatsappSendMessageRequestMediaType{}

	switch fileReferenceType {
	case WhatsappFileLink:
		mediaData.Link = media
	case WhatsappFileID:
		mediaData.ID = media
	default:
		w.Error = fmt.Errorf("fileReferenceType %v not implemented", fileReferenceType)
		return w
	}

	switch mediaType {
	case WhatsAppImage:
		message.Type = external_models.WhatsAppImage
		message.Image = &mediaData
	case WhatsAppAudio:
		message.Type = external_models.WhatsAppAudio
		message.Audio = &mediaData
	case WhatsAppDocument:
		message.Type = external_models.WhatsAppDocument
		message.Document = &mediaData
	case WhatsAppSticker:
		message.Type = external_models.WhatsAppSticker
		message.Sticker = &mediaData
	case WhatsAppVideo:
		message.Type = external_models.WhatsAppVideo
		message.Video = &mediaData
	default:
		w.Error = fmt.Errorf("mediaType %v not implemented", mediaType)
	}

	w.Message = message
	return w
}
