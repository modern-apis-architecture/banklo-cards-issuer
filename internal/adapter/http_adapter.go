package adapter

import (
	"github.com/labstack/echo/v4"
	"github.com/modern-apis-architecture/banklo-cards-issuer/internal/api"
	accountservice "github.com/modern-apis-architecture/banklo-cards-issuer/internal/domain/accounts/service"
	cardservice "github.com/modern-apis-architecture/banklo-cards-issuer/internal/domain/cards/service"
	subscriptionservice "github.com/modern-apis-architecture/banklo-cards-issuer/internal/domain/subscriptions/service"
)

type HttpAdapter struct {
	accountSvc *accountservice.AccountService
	cardSvc    *cardservice.CardService
	subsSvc    subscriptionservice.SubscriptionService
}

func NewHttpAdapter(accountSvc *accountservice.AccountService, cardSvc *cardservice.CardService, subsSvc subscriptionservice.SubscriptionService) *HttpAdapter {
	return &HttpAdapter{
		accountSvc: accountSvc,
		cardSvc:    cardSvc,
		subsSvc:    subsSvc,
	}
}

func (ha *HttpAdapter) CreateAccount(ctx echo.Context) error {
	car := &api.CreateAccountRequest{}
	if err := ctx.Bind(car); err != nil {
		return err
	}
	_, err := ha.accountSvc.CreateAccount(car)
	if err != nil {
		return err
	}

	return nil
}

func (ha *HttpAdapter) GetAccount(ctx echo.Context, id api.CardId) error {

	return nil
}
func (ha *HttpAdapter) CreateCards(ctx echo.Context, id api.CardId) error {
	ccr := &api.CreateCardRequest{}
	if err := ctx.Bind(ccr); err != nil {
		return err
	}

	return nil
}

func (ha *HttpAdapter) DeleteCard(ctx echo.Context, id api.CardId) error {
	return nil
}
func (ha *HttpAdapter) GetCard(ctx echo.Context, id api.CardId) error {
	return nil
}
func (ha *HttpAdapter) PostCardsIdSubscribe(ctx echo.Context, id api.CardId) error { return nil }
