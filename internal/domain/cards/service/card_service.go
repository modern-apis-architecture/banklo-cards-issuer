package service

import (
	"github.com/modern-apis-architecture/banklo-cards-issuer/internal/domain/cards/repository"
)

type CardService struct {
	repo repository.CardRepository
}

func NewCardService(repo repository.CardRepository) *CardService {
	return &CardService{repo: repo}
}
