package external_models

type WhatsappMessageType string
type WhatsappMessageInteractiveType string
type WhatsappMessageInteractiveHeaderType string
type WhatsappMessageInteractiveButtonType string
type WhatsappMessageHomeOrWOrk string

var (
	WhatsAppText        WhatsappMessageType = "text"
	WhatsAppReaction    WhatsappMessageType = "reaction"
	WhatsAppImage       WhatsappMessageType = "image"
	WhatsAppAudio       WhatsappMessageType = "audio"
	WhatsAppDocument    WhatsappMessageType = "document"
	WhatsAppSticker     WhatsappMessageType = "sticker"
	WhatsAppVideo       WhatsappMessageType = "video"
	WhatsAppLocation    WhatsappMessageType = "location"
	WhatsAppContacts    WhatsappMessageType = "contacts"
	WhatsAppInteractive WhatsappMessageType = "interactive"

	WhatsAppContactsHome WhatsappMessageHomeOrWOrk = "HOME"
	WhatsAppContactsWork WhatsappMessageHomeOrWOrk = "WORK"

	WhatsAppInteractiveList           WhatsappMessageInteractiveType = "list"
	WhatsAppInteractiveButton         WhatsappMessageInteractiveType = "button"
	WhatsAppInteractiveCatalogMessage WhatsappMessageInteractiveType = "catalog_message"
	WhatsAppInteractiveProduct        WhatsappMessageInteractiveType = "product"
	WhatsAppInteractiveProductList    WhatsappMessageInteractiveType = "product_list"

	WhatsAppInteractiveTextHeader WhatsappMessageInteractiveHeaderType = "text"

	WhatsAppInteractiveReplyButton WhatsappMessageInteractiveButtonType = "reply"
)

type WhatsappSendMessagePreRequest struct {
	Url                string                     `json:"url"`
	Token              string                     `json:"token"`
	SendMessageRequest WhatsappSendMessageRequest `json:"send_message_request"`
}

type WhatsappSendMessageRequest struct {
	MessagingProduct string                                     `json:"messaging_product"`
	RecipientType    string                                     `json:"recipient_type"`
	To               string                                     `json:"to"`
	Type             WhatsappMessageType                        `json:"type"`
	Context          *WhatsappSendMessageRequestContext         `json:"context,omitempty"`
	Text             *WhatsappSendMessageRequestTextType        `json:"text,omitempty"`
	Reaction         *WhatsappSendMessageRequestReactionType    `json:"reaction,omitempty"`
	Image            *WhatsappSendMessageRequestMediaType       `json:"image,omitempty"`
	Audio            *WhatsappSendMessageRequestMediaType       `json:"audio,omitempty"`
	Document         *WhatsappSendMessageRequestMediaType       `json:"document,omitempty"`
	Sticker          *WhatsappSendMessageRequestMediaType       `json:"sticker,omitempty"`
	Video            *WhatsappSendMessageRequestMediaType       `json:"video,omitempty"`
	Location         *WhatsappSendMessageRequestLocationType    `json:"location,omitempty"`
	Contacts         *[]WhatsappSendMessageRequestContactType   `json:"contacts,omitempty"`
	Interactive      *WhatsappSendMessageRequestInteractiveType `json:"interactive,omitempty"`
}

type WhatsappSendMessageRequestContext struct {
	MessageID string `json:"message_id"`
}

type WhatsappSendMessageRequestTextType struct {
	PreviewUrl bool   `json:"preview_url"`
	Body       string `json:"body"`
}

type WhatsappSendMessageRequestReactionType struct {
	MessageID string `json:"message_id"`
	Emoji     string `json:"emoji"`
}

type WhatsappSendMessageRequestMediaType struct {
	Link string `json:"link"`
	ID   string `json:"id"`
}

type WhatsappSendMessageRequestLocationType struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
	Name      string  `json:"name"`
	Address   string  `json:"address"`
}

type WhatsappSendMessageRequestContactType struct {
	Addresses []struct {
		Street      string                    `json:"street"`
		City        string                    `json:"city"`
		State       string                    `json:"state"`
		Zip         string                    `json:"zip"`
		Country     string                    `json:"country"`
		CountryCode string                    `json:"country_code"`
		Type        WhatsappMessageHomeOrWOrk `json:"type"` // HOME or WORK
	} `json:"addresses"`
	Birthday string `json:"birthday"` //YEAR_MONTH_DAY
	Emails   []struct {
		Email string                    `json:"email"`
		Type  WhatsappMessageHomeOrWOrk `json:"type"` // HOME or WORK
	} `json:"emails"`
	Name struct {
		FormattedName string `json:"formatted_name"`
		FirstName     string `json:"first_name"`
		LastName      string `json:"last_name"`
		MiddleName    string `json:"middle_name"`
		Suffix        string `json:"suffix"`
		Prefix        string `json:"prefix"`
	} `json:"name"`
	Org struct {
		Company    string `json:"company"`
		Department string `json:"department"`
		Title      string `json:"title"`
	} `json:"org"`
	Phones []struct {
		Phone string                    `json:"phone"`
		Type  WhatsappMessageHomeOrWOrk `json:"type"` // HOME or WORK
	} `json:"phones"`
	Urls []struct {
		Url  string                    `json:"url"`
		Type WhatsappMessageHomeOrWOrk `json:"type"` // HOME or WORK
	} `json:"urls"`
}

type WhatsappSendMessageRequestInteractiveType struct {
	Type   WhatsappMessageInteractiveType `json:"type"`
	Header struct {
		Type WhatsappMessageInteractiveHeaderType `json:"type"`
		Text string                               `json:"text"`
	} `json:"header"`
	Body struct {
		Text string `json:"text"`
	} `json:"body"`
	Footer struct {
		Text string `json:"text"`
	} `json:"footer"`
	Action struct {
		Button   string `json:"button,omitempty"`
		Sections *[]struct {
			Title string                                                 `json:"title"`
			Rows  *[]WhatsappSendMessageRequestInteractiveTypeSectionRow `json:"rows,omitempty"`
		} `json:"sections,omitempty"`
		Buttons *[]struct {
			Type  WhatsappMessageInteractiveButtonType `json:"type"`
			Reply struct {
				ID    string `json:"id"`
				Title string `json:"title"`
			} `json:"reply"`
		} `json:"buttons,omitempty"`
	} `json:"action"`
}

type WhatsappSendMessageRequestInteractiveTypeSectionRow struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type WhatsappSendMessageResponse struct {
	MessagingProduct string `json:"messaging_product"`
	Contacts         []struct {
		Input string `json:"input"`
		WaID  string `json:"wa_id"`
	} `json:"contacts"`
	Messages []struct {
		ID string `json:"id"`
	} `json:"messages"`
}
