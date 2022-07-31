package mongo

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type SubscriptionCollection *mongo.Collection

type AccountCollection *mongo.Collection

type CardCollection *mongo.Collection

func ProvideSubscriptionCollection(db *mongo.Database) (SubscriptionCollection, error) {
	return db.Collection("subscriptions"), nil
}

func ProvideAccountCollection(db *mongo.Database) (AccountCollection, error) {
	return db.Collection("accounts"), nil
}

func ProvideCardCollection(db *mongo.Database) (CardCollection, error) {
	return db.Collection("cards"), nil
}
