package models

import (
	"fmt"
	"time"

	"github.com/SineChat/bot-ms/pkg/repository/storage/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (Connection) CollectionName() string {
	return "connections"
}

type Connection struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	NumberID      int                `bson:"number_id" index:"number_id" json:"number_id"` //Your phone number id provided by WhatsApp cloud
	AccountID     primitive.ObjectID `bson:"account_id" index:"account_id" json:"account_id"`
	BotID         primitive.ObjectID `bson:"bot_id" index:"bot_id" json:"bot_id"`
	Type          IntegrationType    `bson:"type" json:"type"`                     // type of bot
	PlatformToken string             `bson:"platform_token" json:"platform_token"` // plaform accesstoken
	MarkAsRead    bool               `bson:"mark_as_read" json:"mark_as_read"`     // Use to set whether incoming messages should be marked as read Default is False
	IsDeployed    bool               `bson:"is_deployed" json:"is_deployed"`
	DeletedAt     time.Time          `bson:"deleted_at" json:"-"`
	CreatedAt     time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt     time.Time          `bson:"updated_at" json:"updated_at"`
}

type CreateConnectionReq struct {
	NumberID      int             `json:"number_id"`
	Type          IntegrationType `json:"type" validate:"required"`
	PlatformToken string          `json:"platform_token"`
	MarkAsRead    bool            `json:"mark_as_read"`
}

type DeployBotToConnectionReq struct {
	BotID        primitive.ObjectID `json:"bot_id" mgvalidate:"notexists=bot$bots$_id"  validate:"required"`
	ConnectionID primitive.ObjectID `json:"connection_id" mgvalidate:"notexists=bot$connections$_id"  validate:"required"`
}

type UpdateConnectionReq struct {
	NumberID      int    `json:"number_id"`
	PlatformToken string `json:"platform_token"`
	MarkAsRead    *bool  `json:"mark_as_read"`
}

func (c *Connection) CreateConnection(db *mongodb.Database) error {
	err := db.CreateOneRecord(&c)
	if err != nil {
		return fmt.Errorf("connection creation failed: %v", err.Error())
	}
	return nil
}

func (c *Connection) GetByID(db *mongodb.Database) error {
	err := db.SelectOneFromDb(&c, bson.M{"_id": c.ID})

	if err != nil {
		return err
	}
	return nil
}

func (c *Connection) GetByNumberID(db *mongodb.Database) error {
	err := db.SelectOneFromDb(&c, bson.M{"number_id": c.NumberID})

	if err != nil {
		return err
	}
	return nil
}

func (c *Connection) GetConnections(db *mongodb.Database, paginator mongodb.Pagination) ([]Connection, mongodb.PaginationResponse, error) {

	var (
		result []Connection
		filter = bson.M{}
	)

	if !c.BotID.IsZero() {
		filter["bot_id"] = c.BotID
	}

	pagination, err := db.SelectPaginatedFromDb("created_at", &c, filter, &result, paginator)

	if err != nil {
		return result, pagination, err
	}
	return result, pagination, nil
}

func (c *Connection) UpdateAll(db *mongodb.Database) error {
	err := db.SaveAllFields(&c)
	if err != nil {
		return fmt.Errorf("Connection update failed: %v", err.Error())
	}
	return nil
}
