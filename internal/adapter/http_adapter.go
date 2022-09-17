package adapter

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/modern-apis-architecture/banklo-cards-issuer/internal/api"
	accountservice "github.com/modern-apis-architecture/banklo-cards-issuer/internal/domain/accounts/service"
	cardservice "github.com/modern-apis-architecture/banklo-cards-issuer/internal/domain/cards/service"
	"github.com/modern-apis-architecture/banklo-cards-issuer/internal/domain/subscriptions"
	subscriptionservice "github.com/modern-apis-architecture/banklo-cards-issuer/internal/domain/subscriptions/service"
)

type HttpAdapter struct {
	accountSvc *accountservice.AccountService
	cardSvc    *cardservice.CardService
	subsSvc    *subscriptionservice.SubscriptionService
}

func NewHttpAdapter(accountSvc *accountservice.AccountService, cardSvc *cardservice.CardService, subsSvc *subscriptionservice.SubscriptionService) *HttpAdapter {
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
	acc, err := ha.accountSvc.CreateAccount(car)
	if err != nil {
		return ctx.JSON(500, api.Error5xxResponse{
			Code:        "010",
			Description: "Internal server error",
		})
	}
	return ctx.JSON(201, acc)
}

func (ha *HttpAdapter) GetAccount(ctx echo.Context, id api.CardId) error {
	acc, err := ha.accountSvc.Get(id)
	if err != nil {
		return ctx.JSON(500, api.Error5xxResponse{
			Code:        "003",
			Description: "internal error",
		})
	}

	if acc == nil {
		return ctx.JSON(404, api.Error4xxResponse{
			Code:        "002",
			Description: "Account not found",
		})
	}
	return ctx.JSON(200, acc)
}
func (ha *HttpAdapter) CreateCards(ctx echo.Context, accountId api.CardId) error {
	ccr := &api.CreateCardRequest{}
	if err := ctx.Bind(ccr); err != nil {
		return err
	}
	cc, err := ha.cardSvc.Create(accountId, ccr)
	if err != nil {
		return ctx.JSON(400, api.Error4xxResponse{
			Code:        "009",
			Description: "Invalid data",
		})
	}
	return ctx.JSON(201, cc)
}

func (ha *HttpAdapter) DeleteCard(ctx echo.Context, id api.CardId) error {
	err := ha.cardSvc.Delete(id)
	if err != nil {
		return ctx.JSON(400, api.Error4xxResponse{
			Code:        "009",
			Description: "Invalid data",
		})
	}
	return ctx.JSON(204, nil)
}
func (ha *HttpAdapter) GetCard(ctx echo.Context, id api.CardId) error {
	card, err := ha.cardSvc.Get(id)
	if err != nil {
		ctx.JSON(404, nil)
	}
	return ctx.JSON(200, card)
}

func (ha *HttpAdapter) PostCardsIdSubscribe(ctx echo.Context, id api.CardId) error {
	cardSub := &api.PostCardsIdSubscribeJSONRequestBody{}
	if err := ctx.Bind(cardSub); err != nil {
		return err
	}
	subId, _ := uuid.NewUUID()
	subs := &subscriptions.Subscription{
		Id:     subId.String(),
		CardId: id,
		Url:    cardSub.Url,
		Token:  cardSub.Token,
	}

	err := ha.subsSvc.Store(subs)
	if err != nil {
		return ctx.JSON(400, api.Error4xxResponse{
			Code:        "009",
			Description: "Invalid subs data",
		})
	}
	return ctx.JSON(201, subs)
}
