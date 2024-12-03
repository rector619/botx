package models

import (
	"fmt"
	"time"

	"github.com/SineChat/bot-ms/pkg/repository/storage/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (Action) CollectionName() string {
	return "actions"
}

type ActionType string

var TextActionType ActionType = "text"

type Action struct {
	ID                   primitive.ObjectID         `bson:"_id,omitempty" json:"id"`
	Type                 ActionType                 `bson:"type" json:"type"`
	ActionID             primitive.ObjectID         `bson:"action_id" json:"action_id"`
	BotID                primitive.ObjectID         `bson:"bot_id" json:"bot_id"`
	Text                 *ActionTextType            `bson:"text,omitempty" json:"text,omitempty"`
	ExpectedResponseType string                     `bson:"expected_response_text" json:"expected_response_text"`
	ExpectedResponseList []string                   `bson:"expected_response_List" json:"expected_response_List"`
	ActionsIDs           map[int]primitive.ObjectID `bson:"actions_ids" json:"-"`
	Actions              map[int]Action             `bson:"-" json:"actions"`
	DefaultResponse      string                     `bson:"default_response" json:"default_response"`
	DeletedAt            time.Time                  `bson:"deleted_at" json:"-"`
	CreatedAt            time.Time                  `bson:"created_at" json:"created_at"`
	UpdatedAt            time.Time                  `bson:"updated_at" json:"updated_at"`
}

type ActionTextType struct {
	PreviewUrl *bool  `json:"preview_url"`
	Body       string `json:"body"`
}

type AddActionReq struct {
	Type            ActionType         `json:"type" validate:"required"`
	BotID           primitive.ObjectID `json:"bot_id" validate:"required"`
	ActionID        primitive.ObjectID `json:"action_id"`
	Position        *int               `json:"position"`
	DefaultResponse string             `json:"default_response"`
	Text            *ActionTextType    `json:"text"`
}

// delete action request
type DeleteActionReq struct {
	ID primitive.ObjectID `json:"id" validate:"required"`
}

func (b *Action) CreateAction(db *mongodb.Database) error {
	err := db.CreateOneRecord(&b)
	if err != nil {
		return fmt.Errorf("action creation failed: %v", err.Error())
	}
	return nil
}

func (a *Action) GetByID(db *mongodb.Database) error {
	err := db.SelectOneFromDb(&a, bson.M{"_id": a.ID})
	if err != nil {
		return err
	}
	return nil
}

func (a *Action) UpdateAll(db *mongodb.Database) error {
	err := db.SaveAllFields(&a)
	if err != nil {
		return fmt.Errorf("Action update failed: %v", err.Error())
	}
	return nil
}

func (a *Action) DeleteAction(db *mongodb.Database) error {
	filter := bson.M{"_id": a.ID}
	err := db.HardDeleteByFilter(&a, filter)
	if err != nil {
		return err
	}

	return nil
}

func (a *Action) DeleteActionMany(db *mongodb.Database, ids []primitive.ObjectID) error {
	filter := bson.M{"_id": bson.M{"$in": ids}}
	err := db.HardDeleteManyByFilter(&a, filter)
	if err != nil {
		return err
	}

	return nil
}
