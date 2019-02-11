package datastore

import (
	"time"
)

type MonetaryTransfer struct {
	From          string    `firebase:"from" json:"from"`
	To            string    `firebase:"to" json:"to"`
	Desc          string    `firebase:"desc" json:"desc"`
	Date          time.Time `firebase:"date" json:"date"`
	AmountUnit    int64     `firebase:"amountUnit" json:"amountUnit"`
	AmountCents   int64     `firebase:"amountCents" json:"amountCents"`
	Currency      string    `firebase:"currency" json:"currency"`
	ConfirmedFrom bool      `firebase:"confirmedFrom" json:"confirmedFrom"`
	ConfirmedTo   bool      `firebase:"confirmedTo" json:"confirmedTo"`
	Snowflake     string    `firebase:"snowflake" json:"snowflake"`
	GroupId       int64     `firebase:"groupId" json:"groupId"`
	RecurrentId   int64     `firebase:"recurrentId" json:"recurrentId"`
}

type GroupTransfer struct {
	From        string    `firebase:"from" json:"from"`
	Tos         []string  `firebase:"tos" json:"tos"`
	Desc        string    `firebase:"desc" json:"desc"`
	Date        time.Time `firebase:"date" json:"date"`
	Included    bool      `firebase:"included" json:"included"`
	AmountUnit  int64     `firebase:"amountUnit" json:"amountUnit"`
	AmountCents int64     `firebase:"amountCents" json:"amountCents"`
	Currency    string    `firebase:"currency" json:"currency"`
	GroupId     int64     `firebase:"groupId" json:"groupId"`
}

type RecurrentTransfer struct {
	RecurrentId int64 `firebase:"recurrent" json:"recurrent"`
	Stamp       int64 `firebase:"stamp" json:"stamp"`
	Concluded   bool  `firebase:"concluded" json:"concluded"`
}
