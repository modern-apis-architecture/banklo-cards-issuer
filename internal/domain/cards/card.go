package cards

import "time"

type Card struct {
	Id             string    `json:"id"`
	ValidUntil     time.Time `json:"valid_until"`
	AccountId      string    `json:"account_id"`
	LastFourDigits string    `json:"last_four_digits"`
	FirstSixDigits string    `json:"first_six_digits"`
}
