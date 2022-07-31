package repository

import "github.com/modern-apis-architecture/banklo-cards-issuer/internal/domain/accounts"

type AccountRepository interface {
	Store(a *accounts.Account) (*accounts.AccountId, error)
	Get(id string) (*accounts.Account, error)
}
