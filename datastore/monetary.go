package datastore

import (
	"cloud.google.com/go/firestore"
	"context"
	"fmt"
)

type MonetaryTransfer struct {
	From          string `firebase:"from" json:"from"`
	To            string `firebase:"to" json:"to"`
	Desc          string `firebase:"desc" json:"desc"`
	Date          string `firebase:"date" json:"date"`
	AmountUnit    int64  `firebase:"amount_unit" json:"amount_unit"`
	AmountCents   byte   `firebase:"amount_cents" json:"amount_cents"`
	Currency      string `firebase:"currency" json:"currency"`
	ConfirmedFrom bool   `firebase:"confirmed_from" json:"confirmed_from"`
	ConfirmedTo   bool   `firebase:"confirmed_to" json:"confirmed_to"`
	GroupId       int64  `firebase:"group_id" json:"group_id"`
	RecurrentId   int64  `firebase:"recurrent_id" json:"recurrent_id"`
}

func getMonetaryTransferFromDB(doc firestore.DocumentRef) (*MonetaryTransfer, error) {
	var mon MonetaryTransfer
	docsnap, err := doc.Get(context.Background())
	if err != nil {
		return nil, fmt.Errorf("no monetary transfer present: %v", err)
	}
	if err = docsnap.DataTo(&mon); err != nil {
		return nil, fmt.Errorf("monetary transfer is not conformant: %v", err)
	}
	return &mon, nil
}

type GroupTransfer struct {
	GroupId     int64                    `firebase:"group_id" json:"group_id"`
	Included    bool                     `firebase:"included" json:"included"`
	AmountUnit  int64                    `firebase:"amount_unit" json:"amount_unit"`
	AmountCents byte                     `firebase:"amount_cents" json:"amount_cents"`
	Tos         []string                 `firebase:"tos" json:"tos"`
	Generated   []*firestore.DocumentRef `firebase:"generated" json:"generated"`
}

func getGroupTransferFromDB(doc firestore.DocumentRef) (*GroupTransfer, error) {
	var grp GroupTransfer
	docsnap, err := doc.Get(context.Background())
	if err != nil {
		return nil, fmt.Errorf("no monetary transfer present: %v", err)
	}
	if err = docsnap.DataTo(&grp); err != nil {
		return nil, fmt.Errorf("monetary transfer is not conformant: %v", err)
	}
	return &grp, nil
}

type RecurrentTransfer struct {
	RecurrentId int64 `firebase:"recurrent" json:"recurrent"`
	Stamp       int64 `firebase:"stamp" json:"stamp"`
	Concluded   bool  `firebase:"concluded" json:"concluded"`
}

func getRecurrentTransferFromDB(doc firestore.DocumentRef) (*RecurrentTransfer, error) {
	var rec RecurrentTransfer
	docsnap, err := doc.Get(context.Background())
	if err != nil {
		return nil, fmt.Errorf("no monetary transfer present: %v", err)
	}
	if err = docsnap.DataTo(&rec); err != nil {
		return nil, fmt.Errorf("monetary transfer is not conformant: %v", err)
	}
	return &rec, nil
}
