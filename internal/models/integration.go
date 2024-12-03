package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IntegrationType string

var (
	WhatsApp IntegrationType = "whatsapp"

	//slice of implemented integrations
	AvailableIntegrations = []IntegrationType{WhatsApp}
)

type Integration struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name        IntegrationType    `bson:"name" json:"name"`
	Description string             `bson:"description" json:"description"`
	Logo        string             `bson:"logo" json:"logo"`
	IsAvailable bool               `bson:"is_available" json:"is_available"`
	Votes       int                `bson:"votes" json:"votes"`
	DeletedAt   time.Time          `bson:"deleted_at" json:"deleted_at"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at" json:"updated_at"`
	Deleted     bool               `bson:"deleted" json:"-"`
}
