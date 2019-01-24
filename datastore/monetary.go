package datastore

import (
	"cloud.google.com/go/firestore"
	"context"
	"fmt"
)

type MonetaryTransfer struct {
	From   string  `firebase:"from" json:"from"`
	To     string  `firebase:"to" json:"to"`
	Desc   string  `firebase:"desc" json:"desc"`
	Date   string  `firebase:"date" json:"date"`
	Amount float32 `firebase:"amount" json:"amount"`
	Payed  bool    `firebase:"payed" json:"payed"`
	Id     int64   `firebase:"id" json:"id"`
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
