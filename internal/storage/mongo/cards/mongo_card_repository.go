package cards

import (
	"context"
	"fmt"
	"github.com/modern-apis-architecture/banklo-cards-issuer/internal/domain/cards"
	mongo2 "github.com/modern-apis-architecture/banklo-cards-issuer/internal/storage/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoCardRepository struct {
	collection *mongo2.CardCollection
}

func NewMongoCardRepository(collection *mongo2.CardCollection) *MongoCardRepository {
	return &MongoCardRepository{collection: collection}
}

func (mcr *MongoCardRepository) Add(c *cards.Card) (*cards.CardCreated, error) {
	ctx := context.Background()
	opts := options.InsertOne()
	doc, err := mcr.collection.Card.InsertOne(ctx, c, opts)
	if err != nil {
		return nil, fmt.Errorf("could not save document to mongo: %w", err)
	}
	id := fmt.Sprintf("%v", doc.InsertedID)
	return &cards.CardCreated{
		Id:             id,
		LastFourDigits: "1234",
		FirstSixDigits: "123421",
	}, nil
}

func (mcr *MongoCardRepository) Get(id string) (*cards.Card, error) {
	var doc cards.Card
	ctx := context.Background()
	if err := mcr.collection.Card.FindOne(ctx, bson.M{"_id": id}).Decode(&doc); err == mongo.ErrNoDocuments {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("error finding cart in mongo collection: %w", err)
	} else {
		return &doc, nil
	}
}

func (mcr *MongoCardRepository) Delete(id string) error {
	ctx := context.Background()
	_, err := mcr.collection.Card.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return fmt.Errorf("error finding cart in mongo collection: %w", err)
	}
	return nil
}
