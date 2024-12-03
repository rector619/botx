package whatsapp

import "github.com/SineChat/bot-ms/external/external_models"

type WhatsAppNotification struct {
	Object string                      `json:"object"`
	Entry  []WhatsAppNotificationEntry `json:"entry"`
}

type WhatsAppNotificationEntry struct {
	ID      string                       `json:"id"`
	Changes []WhatsAppNotificationChange `json:"changes"`
}

type WhatsAppNotificationChange struct {
	Value WhatsAppNotificationMessageValue `json:"value"`
	Field string                           `json:"field"`
}

type WhatsAppNotificationMessageValue struct {
	MessagingProduct string                        `json:"messaging_product"`
	Metadata         WhatsAppNotificationMetadata  `json:"metadata"`
	Contacts         []WhatsAppNotificationContact `json:"contacts"`
	Messages         []WhatsAppNotificationMessage `json:"messages"`
}

type WhatsAppNotificationMetadata struct {
	DisplayPhoneNumber string `json:"display_phone_number"`
	PhoneNumberID      string `json:"phone_number_id"`
}

type WhatsAppNotificationContact struct {
	Profile WhatsAppNotificationProfile `json:"profile"`
	WaID    string                      `json:"wa_id"`
}

type WhatsAppNotificationProfile struct {
	Name string `json:"name"`
}

type WhatsAppNotificationMessage struct {
	Context     WhatsAppNotificationContextMessage  `json:"context,omitempty"`
	From        string                              `json:"from,omitempty"`
	ID          string                              `json:"id,omitempty"`
	Timestamp   string                              `json:"timestamp,omitempty"`
	Type        external_models.WhatsappMessageType `json:"type,omitempty"`
	Image       WhatsAppNotificationMedia           `json:"image,omitempty"`
	Audio       WhatsAppNotificationMedia           `json:"audio,omitempty"`
	Video       WhatsAppNotificationMedia           `json:"video,omitempty"`
	Text        WhatsAppNotificationTextMessage     `json:"text,omitempty"`
	Location    WhatsAppNotificationLocation        `json:"location,omitempty"`
	Reaction    WhatsAppNotificationReaction        `json:"reaction,omitempty"`
	Contacts    WhatsAppNotificationContacts        `json:"contacts,omitempty"`
	Interactive WhatsAppNotificationInteractive     `json:"interactive,omitempty"`
}

type WhatsAppNotificationContextMessage struct {
	Forwarded bool   `json:"forwarded,omitempty"`
	From      string `json:"from,omitempty"`
	ID        string `json:"id,omitempty"`
}

type WhatsAppNotificationMedia struct {
	MimeType string `json:"mime_type,omitempty"`
	Sha256   string `json:"sha256,omitempty"`
	ID       string `json:"id,omitempty"`
	Voice    bool   `json:"voice,omitempty"`
}

type WhatsAppNotificationTextMessage struct {
	Body string `json:"body,omitempty"`
}

type WhatsAppNotificationLocation struct {
	Address   string  `json:"address,omitempty"`
	Latitude  float64 `json:"latitude,omitempty"`
	Longitude float64 `json:"longitude,omitempty"`
	Name      string  `json:"name,omitempty"`
}

type WhatsAppNotificationReaction struct {
	MessageID string `json:"message_id,omitempty"`
	Emoji     string `json:"emoji,omitempty"`
}

type WhatsAppNotificationContacts struct {
	Addresses    []WhatsAppNotificationAddress    `json:"addresses,omitempty"`
	Birthday     string                           `json:"birthday,omitempty"`
	Emails       []WhatsAppNotificationEmail      `json:"emails,omitempty"`
	Name         WhatsAppNotificationName         `json:"name,omitempty"`
	Organization WhatsAppNotificationOrganization `json:"org,omitempty"`
	Phones       []WhatsAppNotificationPhone      `json:"phones,omitempty"`
	URLs         []WhatsAppNotificationURL        `json:"urls,omitempty"`
}

type WhatsAppNotificationAddress struct {
	Street  string `json:"street,omitempty"`
	City    string `json:"city,omitempty"`
	State   string `json:"state,omitempty"`
	Zip     string `json:"zip,omitempty"`
	Country string `json:"country,omitempty"`
	Type    string `json:"type,omitempty"`
}

type WhatsAppNotificationEmail struct {
	Email string `json:"email,omitempty"`
	Type  string `json:"type,omitempty"`
}

type WhatsAppNotificationName struct {
	FirstName     string `json:"first_name,omitempty"`
	LastName      string `json:"last_name,omitempty"`
	FormattedName string `json:"formatted_name,omitempty"`
}

type WhatsAppNotificationOrganization struct {
	Company    string `json:"company,omitempty"`
	Department string `json:"department,omitempty"`
	Title      string `json:"title,omitempty"`
}

type WhatsAppNotificationPhone struct {
	Phone string `json:"phone,omitempty"`
	Type  string `json:"type,omitempty"`
}

type WhatsAppNotificationURL struct {
	URL  string `json:"url,omitempty"`
	Type string `json:"type,omitempty"`
}

type WhatsAppNotificationInteractive struct {
	Type        string                          `json:"type,omitempty"`
	ListReply   WhatsAppNotificationListReply   `json:"list_reply,omitempty"`
	ButtonReply WhatsAppNotificationButtonReply `json:"button_reply,omitempty"`
}

type WhatsAppNotificationListReply struct {
	ID          string `json:"id,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
}

type WhatsAppNotificationButtonReply struct {
	ID    string `json:"id,omitempty"`
	Title string `json:"title,omitempty"`
}
