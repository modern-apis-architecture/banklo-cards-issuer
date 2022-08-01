package service

import "github.com/modern-apis-architecture/banklo-cards-issuer/internal/domain/subscriptions/repository"

type SubscriptionService struct {
	repo repository.SubscriptionRepository
}

func NewSubscriptionService(repo repository.SubscriptionRepository) *SubscriptionService {
	return &SubscriptionService{repo: repo}
}
