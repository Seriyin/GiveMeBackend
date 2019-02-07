package request

import (
	"strings"
	"json"
	"context"
	"github.com/Seriyin/GibMe-backend/config/firebase"
	"github.com/Seriyin/GibMe-backend/config/firebase/firestore"
)

var db = firebase.GetDB()

func Request(
	ctx context.Context,
	e firestore.Event,
) {
	var monetaryT MonetaryTransfer //Monetary Structure object
	err := json.Unmarshal(e.Value.Fields, &monetaryT) //Json object to Monetary Structure
	debtor := monetaryT.To //Debtor present in the To value on the Monetary Struct

	/*
	*Use the debtor id to split the path
	*so it gets the path next to the creditor ID 
	*to be replaced by the debtor ID, so it inserts on the correct collection
	*/
	splitvalue := "/"+monetaryT.From+"/"
	path := strings.Split(e.Value.Name, splitvalue)[1]

	db.setMonetaryTransfer(debtor, newmonetaryT, path)
}
