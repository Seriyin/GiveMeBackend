package firestore

import (
	"encoding/json"
	"github.com/Seriyin/GiveMeBackend/config/datastore"
	"time"
)

type StringValue struct {
	StringValue string `json:"stringValue"`
}

type IntegerValue struct {
	IntegerValue int64 `json:"integerValue, string"`
}

type BooleanValue struct {
	BooleanValue bool `json:"booleanValue"`
}

type TimestampValue struct {
	TimestampValue time.Time `json:"timestampValue"`
}

type monetaryRequest struct {
	From          StringValue    `json:"from"`
	To            StringValue    `json:"to"`
	Desc          StringValue    `json:"desc"`
	Date          TimestampValue `json:"date"`
	AmountUnit    IntegerValue   `json:"amountUnit"`
	AmountCents   IntegerValue   `json:"amountCents"`
	Currency      StringValue    `json:"currency"`
	ConfirmedFrom BooleanValue   `json:"confirmedFrom"`
	ConfirmedTo   BooleanValue   `json:"confirmedTo"`
	Snowflake     StringValue    `json:"snowflake"`
	GroupId       IntegerValue   `json:"groupId"`
	RecurrentId   IntegerValue   `json:"recurrentId"`
}

type groupRequest struct {
	From        StringValue    `json:"from"`
	Tos         []StringValue  `json:"tos"`
	Desc        StringValue    `json:"desc"`
	Date        TimestampValue `json:"date"`
	Included    BooleanValue   `json:"included"`
	AmountUnit  IntegerValue   `json:"amountUnit"`
	AmountCents IntegerValue   `json:"amountCents"`
	Currency    StringValue    `json:"currency"`
	GroupId     IntegerValue   `json:"groupId"`
}

func UnmarshallAndConvertMonetary(
	message json.RawMessage,
) (*datastore.MonetaryRequest, error) {
	var mon monetaryRequest
	err := json.Unmarshal(message, &mon)
	if err != nil {
		return nil, err
	}
	return &datastore.MonetaryRequest{
		From:          mon.From.StringValue,
		To:            mon.To.StringValue,
		Desc:          mon.Desc.StringValue,
		Date:          mon.Date.TimestampValue,
		AmountUnit:    mon.AmountUnit.IntegerValue,
		AmountCents:   mon.AmountCents.IntegerValue,
		Currency:      mon.Currency.StringValue,
		ConfirmedFrom: mon.ConfirmedFrom.BooleanValue,
		ConfirmedTo:   mon.ConfirmedTo.BooleanValue,
		Snowflake:     mon.Snowflake.StringValue,
		GroupId:       mon.GroupId.IntegerValue,
		RecurrentId:   mon.RecurrentId.IntegerValue,
	}, nil
}

func UnmarshallAndConvertGroup(
	message json.RawMessage,
) (*datastore.GroupRequest, error) {
	var grp groupRequest
	err := json.Unmarshal(message, &grp)
	if err != nil {
		return nil, err
	}
	tos := make([]string, len(grp.Tos))
	for _, to := range grp.Tos {
		tos = append(tos, to.StringValue)
	}
	return &datastore.GroupRequest{
		From:        grp.From.StringValue,
		Tos:         tos,
		Desc:        grp.Desc.StringValue,
		Date:        grp.Date.TimestampValue,
		Included:    grp.Included.BooleanValue,
		AmountUnit:  grp.AmountUnit.IntegerValue,
		AmountCents: grp.AmountCents.IntegerValue,
		Currency:    grp.Currency.StringValue,
		GroupId:     grp.GroupId.IntegerValue,
	}, nil
}
