package firestore

import (
	"encoding/json"
	"github.com/Seriyin/GiveMeBackend/config/datastore"
	"strings"
	"time"
)

type monetaryRequest struct {
	From          string `json:"from"`
	To            string `json:"to"`
	Desc          string `json:"desc"`
	Date          string `json:"date"`
	AmountUnit    int64  `json:"amountUnit"`
	AmountCents   int64  `json:"amountCents"`
	Currency      string `json:"currency"`
	ConfirmedFrom bool   `json:"confirmedFrom"`
	ConfirmedTo   bool   `json:"confirmedTo"`
	Snowflake     string `json:"snowflake"`
	GroupId       int64  `json:"groupId"`
	RecurrentId   int64  `json:"recurrentId"`
}

func UnmarshallAndConvert(message json.RawMessage) (*datastore.MonetaryRequest, error) {
	var mon monetaryRequest
	err := json.Unmarshal(message, &mon)
	if err != nil {
		return nil, err
	}
	date, err := time.Parse("2006-01-02T15:04:05", strings.SplitN(mon.Date, ".", 1)[0])
	return &datastore.MonetaryRequest{
		From:          mon.From,
		To:            mon.To,
		Desc:          mon.Desc,
		Date:          date,
		AmountUnit:    mon.AmountUnit,
		AmountCents:   mon.AmountCents,
		Currency:      mon.Currency,
		ConfirmedFrom: mon.ConfirmedFrom,
		ConfirmedTo:   mon.ConfirmedTo,
		Snowflake:     mon.Snowflake,
		GroupId:       mon.GroupId,
		RecurrentId:   mon.RecurrentId,
	}, nil
}
