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
	GroupId       int64     `firebase:"groupId" json:"groupId"`
	RecurrentId   int64     `firebase:"recurrentId" json:"recurrentId"`
}

type GroupTransfer struct {
	GroupId     int64     `firebase:"groupId" json:"groupId"`
	Included    bool      `firebase:"included" json:"included"`
	Date        time.Time `firebase:"date" json:"date"`
	AmountUnit  int64     `firebase:"amountUnit" json:"amountUnit"`
	AmountCents int64     `firebase:"amountCents" json:"amountCents"`
	Currency    string    `firebase:"currency" json:"currency"`
	Tos         []string  `firebase:"tos" json:"tos"`
}

type RecurrentTransfer struct {
	RecurrentId int64 `firebase:"recurrent" json:"recurrent"`
	Stamp       int64 `firebase:"stamp" json:"stamp"`
	Concluded   bool  `firebase:"concluded" json:"concluded"`
}
