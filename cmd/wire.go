//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/modern-apis-architecture/banklo-cards-issuer/internal/adapter"
	repository2 "github.com/modern-apis-architecture/banklo-cards-issuer/internal/domain/accounts/repository"
	"github.com/modern-apis-architecture/banklo-cards-issuer/internal/domain/accounts/service"
	"github.com/modern-apis-architecture/banklo-cards-issuer/internal/domain/cards/repository"
	service2 "github.com/modern-apis-architecture/banklo-cards-issuer/internal/domain/cards/service"
	repository3 "github.com/modern-apis-architecture/banklo-cards-issuer/internal/domain/subscriptions/repository"
	service3 "github.com/modern-apis-architecture/banklo-cards-issuer/internal/domain/subscriptions/service"
	"github.com/modern-apis-architecture/banklo-cards-issuer/internal/storage/mongo"
	"github.com/modern-apis-architecture/banklo-cards-issuer/internal/storage/mongo/accounts"
	"github.com/modern-apis-architecture/banklo-cards-issuer/internal/storage/mongo/cards"
	"github.com/modern-apis-architecture/banklo-cards-issuer/internal/storage/mongo/subscriptions"
)

func buildAppContainer() (*adapter.HttpAdapter, error) {
	wire.Build(mongo.NewDatabase, mongo.ProvideAccountCollection, mongo.ProvideCardCollection, mongo.ProvideSubscriptionCollection,
		cards.NewMongoCardRepository, accounts.NewMongoAccountRepository, subscriptions.NewMongoSubscriptionRepository,
		wire.Bind(new(repository.CardRepository), new(*cards.MongoCardRepository)),
		wire.Bind(new(repository3.SubscriptionRepository), new(*subscriptions.MongoSubscriptionRepository)),
		wire.Bind(new(repository2.AccountRepository), new(*accounts.MongoAccountRepository)),
		service.NewAccountService, service2.NewCardService, service3.NewSubscriptionService,
		adapter.NewHttpAdapter,
	)
	return nil, nil
}
