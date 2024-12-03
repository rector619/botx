package external_models

import "go.mongodb.org/mongo-driver/bson/primitive"

type AccessToken struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	AccountID  primitive.ObjectID `bson:"account_id" json:"account_id"`
	PublicKey  string             `bson:"public_key" json:"public_key"`
	PrivateKey string             `bson:"private_key" json:"private_key"`
	IsLive     bool               `bson:"is_live" json:"is_live"`
	CreatedAt  string             `bson:"created_at" json:"created_at"`
	UpdatedAt  string             `bson:"updated_at" json:"updated_at"`
}

type GetAccessTokenModel struct {
	Status  string      `json:"status"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    AccessToken `json:"data"`
}
