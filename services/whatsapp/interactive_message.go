package whatsapp

import "github.com/SineChat/bot-ms/external/external_models"

type WhatsappInteractiveButton struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

type WhatsappInteractiveSection struct {
	Title string                          `json:"title"`
	Rows  []WhatsappInteractiveSectionRow `json:"rows"`
}
type WhatsappInteractiveSectionRow struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (w *Whatsapp) InteractiveButtonMessage(phoneNumber string, header, body, footer string, buttons []WhatsappInteractiveButton) *Whatsapp {
	message := external_models.WhatsappSendMessageRequest{
		MessagingProduct: MessagingProduct,
		RecipientType:    IndividualRecipientType,
		To:               phoneNumber,
		Type:             external_models.WhatsAppInteractive,
		Interactive: &external_models.WhatsappSendMessageRequestInteractiveType{
			Type: external_models.WhatsAppInteractiveButton,
			Header: struct {
				Type external_models.WhatsappMessageInteractiveHeaderType "json:\"type\""
				Text string                                               "json:\"text\""
			}{
				Type: external_models.WhatsAppInteractiveTextHeader,
				Text: header,
			},
			Body: struct {
				Text string "json:\"text\""
			}{
				Text: body,
			},
			Footer: struct {
				Text string "json:\"text\""
			}{
				Text: footer,
			},
		},
	}

	buttonsMade := make([]struct {
		Type  external_models.WhatsappMessageInteractiveButtonType "json:\"type\""
		Reply struct {
			ID    string "json:\"id\""
			Title string "json:\"title\""
		} "json:\"reply\""
	}, 0)

	for _, b := range buttons {
		buttonsMade = append(buttonsMade, struct {
			Type  external_models.WhatsappMessageInteractiveButtonType "json:\"type\""
			Reply struct {
				ID    string "json:\"id\""
				Title string "json:\"title\""
			} "json:\"reply\""
		}{
			Type: external_models.WhatsAppInteractiveReplyButton,
			Reply: struct {
				ID    string "json:\"id\""
				Title string "json:\"title\""
			}{
				ID:    b.ID,
				Title: b.Title,
			},
		})
	}

	message.Interactive.Action.Buttons = &buttonsMade

	w.Message = message
	return w
}

func (w *Whatsapp) InteractiveListMessage(phoneNumber string, header, body, footer, buttonText string, sections []WhatsappInteractiveSection) *Whatsapp {
	message := external_models.WhatsappSendMessageRequest{
		MessagingProduct: MessagingProduct,
		RecipientType:    IndividualRecipientType,
		To:               phoneNumber,
		Type:             external_models.WhatsAppInteractive,
		Interactive: &external_models.WhatsappSendMessageRequestInteractiveType{
			Type: external_models.WhatsAppInteractiveList,
			Header: struct {
				Type external_models.WhatsappMessageInteractiveHeaderType "json:\"type\""
				Text string                                               "json:\"text\""
			}{
				Type: external_models.WhatsAppInteractiveTextHeader,
				Text: header,
			},
			Body: struct {
				Text string "json:\"text\""
			}{
				Text: body,
			},
			Footer: struct {
				Text string "json:\"text\""
			}{
				Text: footer,
			},
			Action: struct {
				Button   string "json:\"button,omitempty\""
				Sections *[]struct {
					Title string                                                                 "json:\"title\""
					Rows  *[]external_models.WhatsappSendMessageRequestInteractiveTypeSectionRow "json:\"rows,omitempty\""
				} "json:\"sections,omitempty\""
				Buttons *[]struct {
					Type  external_models.WhatsappMessageInteractiveButtonType "json:\"type\""
					Reply struct {
						ID    string "json:\"id\""
						Title string "json:\"title\""
					} "json:\"reply\""
				} "json:\"buttons,omitempty\""
			}{
				Button: buttonText,
			},
		},
	}

	sectionsMade := make([]struct {
		Title string                                                                 "json:\"title\""
		Rows  *[]external_models.WhatsappSendMessageRequestInteractiveTypeSectionRow "json:\"rows,omitempty\""
	}, 0)

	for _, s := range sections {
		rows := []external_models.WhatsappSendMessageRequestInteractiveTypeSectionRow{}
		for _, r := range s.Rows {
			rows = append(rows, external_models.WhatsappSendMessageRequestInteractiveTypeSectionRow{
				ID:          r.ID,
				Title:       r.Title,
				Description: r.Description,
			})
		}

		sectionsMade = append(sectionsMade, struct {
			Title string                                                                 "json:\"title\""
			Rows  *[]external_models.WhatsappSendMessageRequestInteractiveTypeSectionRow "json:\"rows,omitempty\""
		}{
			Title: s.Title,
			Rows:  &rows,
		})

	}

	message.Interactive.Action.Sections = &sectionsMade

	w.Message = message
	return w
}
