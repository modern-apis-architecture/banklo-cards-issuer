package repository

import "github.com/modern-apis-architecture/banklo-cards-issuer/internal/domain/subscriptions"

type SubscriptionRepository interface {
	Store(s *subscriptions.Subscription) error
	Find(cardId string) ([]*subscriptions.Subscription, error)
}
