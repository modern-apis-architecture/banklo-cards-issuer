package subscriptions

type Subscription struct {
	Id     string `json:"id"`
	CardId string `json:"card_id"`
	Url    string `json:"url"`
	Token  string `json:"token"`
}
