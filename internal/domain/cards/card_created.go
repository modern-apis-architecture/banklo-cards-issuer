package cards

type CardCreated struct {
	Id             string `json:"id"`
	LastFourDigits string `json:"last_four_digits"`
	FirstSixDigits string `json:"first_six_digits"`
}
