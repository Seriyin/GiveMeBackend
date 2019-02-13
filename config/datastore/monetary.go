package datastore

import (
	"time"
)

type MonetaryRequest struct {
	From          string    `firestore:"from" json:"from"`
	To            string    `firestore:"to" json:"to"`
	Desc          string    `firestore:"desc" json:"desc"`
	Date          time.Time `firestore:"date" json:"date"`
	AmountUnit    int64     `firestore:"amountUnit" json:"amountUnit"`
	AmountCents   int64     `firestore:"amountCents" json:"amountCents"`
	Currency      string    `firestore:"currency" json:"currency"`
	ConfirmedFrom bool      `firestore:"confirmedFrom" json:"confirmedFrom"`
	ConfirmedTo   bool      `firestore:"confirmedTo" json:"confirmedTo"`
	Snowflake     string    `firestore:"snowflake" json:"snowflake"`
	GroupId       int64     `firestore:"groupId" json:"groupId"`
	RecurrentId   int64     `firestore:"recurrentId" json:"recurrentId"`
}

type GroupRequest struct {
	From        string    `firestore:"from" json:"from"`
	Tos         []string  `firestore:"tos" json:"tos"`
	Desc        string    `firestore:"desc" json:"desc"`
	Date        time.Time `firestore:"date" json:"date"`
	Included    bool      `firestore:"included" json:"included"`
	AmountUnit  int64     `firestore:"amountUnit" json:"amountUnit"`
	AmountCents int64     `firestore:"amountCents" json:"amountCents"`
	Currency    string    `firestore:"currency" json:"currency"`
	GroupId     int64     `firestore:"groupId" json:"groupId"`
}

type RecurrentRequest struct {
	RecurrentId int64 `firestore:"recurrent" json:"recurrent"`
	Stamp       int64 `firestore:"stamp" json:"stamp"`
	Concluded   bool  `firestore:"concluded" json:"concluded"`
}
