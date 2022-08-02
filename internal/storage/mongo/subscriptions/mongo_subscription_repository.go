package subscriptions

import (
	"context"
	"fmt"
	"github.com/modern-apis-architecture/banklo-cards-issuer/internal/domain/subscriptions"
	mongo2 "github.com/modern-apis-architecture/banklo-cards-issuer/internal/storage/mongo"
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
