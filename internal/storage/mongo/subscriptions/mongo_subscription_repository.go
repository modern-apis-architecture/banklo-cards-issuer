package subscriptions

import (
	"context"
	"fmt"
	"github.com/modern-apis-architecture/banklo-cards-issuer/internal/domain/subscriptions"
	mongo2 "github.com/modern-apis-architecture/banklo-cards-issuer/internal/storage/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoSubscriptionRepository struct {
	collection *mongo2.SubscriptionCollection
}

func NewMongoSubscriptionRepository(collection *mongo2.SubscriptionCollection) *MongoSubscriptionRepository {
	return &MongoSubscriptionRepository{collection: collection}
}

func (msr *MongoSubscriptionRepository) Store(s *subscriptions.Subscription) error {
	ctx := context.Background()
	opts := options.InsertOne()
	_, err := msr.collection.Subs.InsertOne(ctx, s, opts)
	if err != nil {
		return fmt.Errorf("could not save document to mongo: %w", err)
	}
	return nil
}

func (msr *MongoSubscriptionRepository) Find(cardId string) ([]*subscriptions.Subscription, error) {
	ctx := context.Background()
	filter := bson.D{{"card_id", cardId}}
	var subs []*subscriptions.Subscription
	cur, err := msr.collection.Subs.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("could not find subscriptions by card id: %w", err)
	}
	for cur.Next(ctx) {
		var s subscriptions.Subscription
		err := cur.Decode(&s)
		if err != nil {
			return subs, err
		}
		subs = append(subs, &s)
	}
	if err := cur.Err(); err != nil {
		return subs, err
	}
	cur.Close(ctx)
	return subs, nil
}
