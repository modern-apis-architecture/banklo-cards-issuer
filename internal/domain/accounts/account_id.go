package accounts

type AccountId struct {
	Id string `json:"id"`
}

func NewAccountId(id string) *AccountId {
	return &AccountId{Id: id}
}
