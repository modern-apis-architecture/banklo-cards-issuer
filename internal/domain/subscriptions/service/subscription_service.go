package service

import (
	"github.com/modern-apis-architecture/banklo-cards-issuer/internal/domain/subscriptions"
	"github.com/modern-apis-architecture/banklo-cards-issuer/internal/domain/subscriptions/repository"
)

type SubscriptionService struct {
	repo repository.SubscriptionRepository
}

func NewSubscriptionService(repo repository.SubscriptionRepository) *SubscriptionService {
	return &SubscriptionService{repo: repo}
}

func (ss *SubscriptionService) Store(s *subscriptions.Subscription) error {
	return ss.repo.Store(s)
}
func (ss *SubscriptionService) FindSubsByCardId(cardId string) ([]*subscriptions.Subscription, error) {
	return ss.repo.Find(cardId)
}
