package models

import (
	"fmt"
	"time"

	"github.com/SineChat/bot-ms/pkg/repository/storage/mongodb"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (WebhookLog) CollectionName() string {
	return "webhook_log"
}

type WebhookLog struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	BotID     primitive.ObjectID `bson:"bot_id,omitempty" json:"bot_id"`
	Platform  IntegrationType    `bson:"platform" json:"platform"`
	Payload   string             `bson:"response_payload" json:"response_payload"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
	DeletedAt time.Time          `bson:"deleted_at" json:"-"`
	Deleted   bool               `bson:"deleted" json:"-"`
}

func (w *WebhookLog) CreateWebhookLog(db *mongodb.Database) error {
	err := db.CreateOneRecord(&w)
	if err != nil {
		return fmt.Errorf("webhook log creation failed: %v", err.Error())
	}
	return nil
}
