package models

import (
	"fmt"
	"time"

	"github.com/SineChat/bot-ms/pkg/repository/storage/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (Bot) CollectionName() string {
	return "bots"
}

type Bot struct {
	ID        primitive.ObjectID   `bson:"_id,omitempty" json:"id"`
	AccountID primitive.ObjectID   `bson:"account_id" index:"account_id" json:"account_id"`
	Name      string               `bson:"name" json:"name"`
	Secret    string               `bson:"webhook_secret" json:"-"`
	Action    primitive.ObjectID   `bson:"template" json:"-"` // object id here is the Action.ID
	Template  Action               `bson:"-" json:"template"`
	Deployed  []primitive.ObjectID `bson:"deployed" json:"deployed"` // slice of   connections deployed to
	DeletedAt time.Time            `bson:"deleted_at" json:"-"`
	CreatedAt time.Time            `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time            `bson:"updated_at" json:"updated_at"`
	Deleted   bool                 `bson:"deleted" json:"-"`
}

type BotWebhookData struct {
	Secret     string `json:"secret"`
	WebhookUrl string `json:"webhook_url"`
}

type CreateBotReq struct {
	Name string `json:"name"`
}

func (b *Bot) CreateBot(db *mongodb.Database) error {
	err := db.CreateOneRecord(&b)
	if err != nil {
		return fmt.Errorf("bot creation failed: %v", err.Error())
	}
	return nil
}

func (b *Bot) GetByID(db *mongodb.Database) error {
	err := db.SelectOneFromDb(&b, bson.M{"_id": b.ID})

	if err != nil {
		return err
	}
	return nil
}

func (b *Bot) UpdateAll(db *mongodb.Database) error {
	err := db.SaveAllFields(&b)
	if err != nil {
		return fmt.Errorf("bot update failed: %v", err.Error())
	}
	return nil
}

// Get all bots
func (b *Bot) GetAllBots(db *mongodb.Database, paginator mongodb.Pagination) ([]Bot, mongodb.PaginationResponse, error) {

	var (
		result []Bot
		filter = bson.M{}
	)
	if !b.AccountID.IsZero() {
		filter["account_id"] = b.AccountID
	}

	pagination, err := db.SelectPaginatedFromDb("created_at", &b, filter, &result, paginator)
	if err != nil {
		return result, pagination, err
	}
	return result, pagination, nil
}
