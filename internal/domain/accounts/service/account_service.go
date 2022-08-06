package service

import (
	"github.com/google/uuid"
	"github.com/modern-apis-architecture/banklo-cards-issuer/internal/api"
	"github.com/modern-apis-architecture/banklo-cards-issuer/internal/domain/accounts"
	"github.com/modern-apis-architecture/banklo-cards-issuer/internal/domain/accounts/repository"
	"time"
)

type AccountService struct {
	repo repository.AccountRepository
}

func NewAccountService(repo repository.AccountRepository) *AccountService {
	return &AccountService{repo: repo}
}

func (ac *AccountService) CreateAccount(ar *api.CreateAccountRequest) (*accounts.AccountId, error) {
	id, _ := uuid.NewUUID()
	acc := &accounts.Account{
		Id:         id.String(),
		Name:       ar.PersonalData.Name,
		LastName:   ar.PersonalData.LastName,
		MotherName: ar.PersonalData.MotherName,
		Document:   ar.PersonalData.Document,
		BirthDate:  ar.PersonalData.BirthDate.String(),
		Address: accounts.Address{
			ZipCode: ar.PersonalData.Address.ZipCode,
			Number:  ar.PersonalData.Address.Number,
		},
		Phone: accounts.Phone{
			Code:   ar.PersonalData.Phone.Code,
			Number: ar.PersonalData.Phone.Number,
		},
		CreatedAt: time.Now(),
	}
	return ac.repo.Store(acc)
}

func (ac *AccountService) Get(id string) (*accounts.Account, error) {
	return ac.repo.Get(id)
}
