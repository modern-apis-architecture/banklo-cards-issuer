package subscriptions

import (
	"context"
	"fmt"
	"github.com/modern-apis-architecture/banklo-cards-issuer/internal/domain/subscriptions"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoSubscriptionRepository struct {
	collection *mongo.Collection
}

func NewMongoSubscriptionRepository(collection *mongo.Collection) *MongoSubscriptionRepository {
	return &MongoSubscriptionRepository{collection: collection}
}

func (msr *MongoSubscriptionRepository) Store(s *subscriptions.Subscription) error {
	ctx := context.Background()
	opts := options.InsertOne()
	_, err := msr.collection.InsertOne(ctx, s, opts)
	if err != nil {
		return fmt.Errorf("could not save document to mongo: %w", err)
	}
	return nil
}
