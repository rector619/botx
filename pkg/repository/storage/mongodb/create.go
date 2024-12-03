package mongodb

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (db *Database) CreateOneRecord(model interface{}) error {
	coll, err := db.GetCollectionForModel(model)
	if err != nil {
		return err
	}

	result, err := coll.InsertOne(context.Background(), model)
	if err != nil {
		return err
	}

	if result.InsertedID == nil {
		return fmt.Errorf("record creation for %v failed", coll.Name())
	}

	insertedID := result.InsertedID.(primitive.ObjectID)

	filter := bson.M{"_id": insertedID}
	update := bson.M{"$set": bson.M{"updated_at": time.Now(), "created_at": time.Now(), "deleted": false}}
	re, err := coll.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	if re.ModifiedCount > 0 {
		err := db.SelectOneFromDb(model, bson.M{"_id": insertedID})
		if err != nil {
			return err
		}
	}
	return nil
}

// delete one record from the database
func (db *Database) DeleteOneRecord(model interface{}, filter bson.M) error {
	coll, err := db.GetCollectionForModel(model)
	if err != nil {
		return err
	}

	update := bson.M{"$set": bson.M{"updated_at": time.Now(), "deleted": true}}
	_, err = coll.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}
	return nil
}
