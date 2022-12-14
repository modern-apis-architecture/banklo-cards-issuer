package accounts

import (
	"context"
	"fmt"
	"github.com/modern-apis-architecture/banklo-cards-issuer/internal/domain/accounts"
	"github.com/modern-apis-architecture/banklo-cards-issuer/internal/storage/mongo"
	"go.mongodb.org/mongo-driver/bson"
	mongo2 "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoAccountRepository struct {
	collection *mongo.AccountCollection
}

func NewMongoAccountRepository(col *mongo.AccountCollection) *MongoAccountRepository {
	return &MongoAccountRepository{collection: col}
}

func (mar *MongoAccountRepository) Store(a *accounts.Account) (*accounts.AccountId, error) {
	ctx := context.Background()
	opts := options.InsertOne()
	doc, err := mar.collection.Account.InsertOne(ctx, a, opts)
	if err != nil {
		return nil, fmt.Errorf("could not save document to mongo: %w", err)
	}
	id := fmt.Sprintf("%v", doc.InsertedID)
	return accounts.NewAccountId(id), nil

}

func (mar *MongoAccountRepository) Get(id string) (*accounts.Account, error) {
	var doc accounts.Account
	ctx := context.Background()
	if err := mar.collection.Account.FindOne(ctx, bson.M{"_id": id}).Decode(&doc); err == mongo2.ErrNoDocuments {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("error finding cart in mongo collection: %w", err)
	} else {
		return &doc, nil
	}
}
