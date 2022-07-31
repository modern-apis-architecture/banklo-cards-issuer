package repository

import "github.com/modern-apis-architecture/banklo-cards-issuer/internal/domain/cards"

type CardRepository interface {
	Add(c *cards.Card) (*cards.CardCreated, error)
	Get(id string) (*cards.Card, error)
	Delete(id string) error
}
