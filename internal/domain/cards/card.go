package cards

import "time"

type Card struct {
	Id             string    `json:"id" bson:"_id"`
	ValidUntil     time.Time `json:"valid_until" bson:"valid_until"`
	AccountId      string    `json:"account_id" bson:"account_id"`
	LastFourDigits string    `json:"last_four_digits" bson:"last_four_digits"`
	FirstSixDigits string    `json:"first_six_digits" bson:"first_six_digits"`
	CreatedAt      time.Time `json:"created_at" bson:"created_at"`
	Name           string    `json:"name" bson:"name"`
}
