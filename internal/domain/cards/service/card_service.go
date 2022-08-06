package service

import (
	"errors"
	"github.com/google/uuid"
	"github.com/modern-apis-architecture/banklo-cards-issuer/internal/api"
	repository2 "github.com/modern-apis-architecture/banklo-cards-issuer/internal/domain/accounts/repository"
	"github.com/modern-apis-architecture/banklo-cards-issuer/internal/domain/cards"
	"github.com/modern-apis-architecture/banklo-cards-issuer/internal/domain/cards/repository"
	"time"
)

var ErrorAccountNotFound = errors.New("Account not found ")

type CardService struct {
	cardRepo    repository.CardRepository
	accountRepo repository2.AccountRepository
}

func NewCardService(cardRepo repository.CardRepository, acRepo repository2.AccountRepository) *CardService {
	return &CardService{cardRepo: cardRepo, accountRepo: acRepo}
}

func (cs *CardService) Delete(id string) error {
	return cs.cardRepo.Delete(id)
}

func (cs *CardService) Get(id string) (*cards.Card, error) {
	return cs.cardRepo.Get(id)
}

func (cs *CardService) Create(acId string, cr *api.CreateCardRequest) (*cards.CardCreated, error) {
	if _, err := cs.accountRepo.Get(acId); err != nil {
		return nil, ErrorAccountNotFound
	}
	id, _ := uuid.NewUUID()
	card := &cards.Card{
		Id:             id.String(),
		ValidUntil:     time.Now().Add(time.Hour * 60),
		AccountId:      acId,
		LastFourDigits: "4444",
		FirstSixDigits: "525342",
		Name:           cr.Name,
	}
	cc, err := cs.cardRepo.Add(card)
	if err != nil {
		return nil, err
	}
	return cc, nil
}
