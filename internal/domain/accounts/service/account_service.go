package service

import (
	"github.com/modern-apis-architecture/banklo-cards-issuer/internal/api"
	"github.com/modern-apis-architecture/banklo-cards-issuer/internal/domain/accounts"
	"github.com/modern-apis-architecture/banklo-cards-issuer/internal/domain/accounts/repository"
)

type AccountService struct {
	repo repository.AccountRepository
}

func NewAccountService(repo repository.AccountRepository) *AccountService {
	return &AccountService{repo: repo}
}

func (ac *AccountService) CreateAccount(ar *api.CreateAccountRequest) (*accounts.AccountId, error) {
	return nil, nil
}

func (ac *AccountService) Get(id string) (*accounts.Account, error) {

	return nil, nil
}
