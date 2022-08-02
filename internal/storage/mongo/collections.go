package mongo

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type SubscriptionCollection struct {
	Subs *mongo.Collection
}

type AccountCollection struct {
	Account *mongo.Collection
}

type CardCollection struct {
	Card *mongo.Collection
}

func ProvideSubscriptionCollection(db *mongo.Database) (*SubscriptionCollection, error) {
	return &SubscriptionCollection{Subs: db.Collection("subscriptions")}, nil
}

func ProvideAccountCollection(db *mongo.Database) (*AccountCollection, error) {
	return &AccountCollection{Account: db.Collection("accounts")}, nil
}

func ProvideCardCollection(db *mongo.Database) (*CardCollection, error) {
	return &CardCollection{Card: db.Collection("cards")}, nil
}
