package request

import (
	"context"
	"encoding/json"
	"log"
	"strings"

	"github.com/Seriyin/GibMe-backend/config/datastore"
	"github.com/Seriyin/GibMe-backend/config/firebase"
	"github.com/Seriyin/GibMe-backend/config/firebase/firestore"
	"github.com/Seriyin/GibMe-backend/config/firebase/messaging"
)

var db = firebase.GetDB()
var mesClient = firebase.GetMessaging()

func Request(
	ctx context.Context,
	e firestore.Event,
) error {
	var monetaryT datastore.MonetaryTransfer          // Monetary Structure object
	err := json.Unmarshal(e.Value.Fields, &monetaryT) // Json object to Monetary Structure
	if err != nil {
		return err
	}
	debtor := monetaryT.To // Debtor present in the To value on the Monetary Struct

	// Use the debtor id to split the path
	// so it gets the path next to the creditor ID
	// to be replaced by the debtor ID, so it inserts on the correct collection.
	splitvalue := "/" + monetaryT.From + "/"
	path := strings.Split(e.Value.Name, splitvalue)[1]

	_, err = db.SetMonetaryTransfer(debtor, &monetaryT, path)
	if err != nil {
		return err
	}

	//generate notification message
	token := "" //IMPLEMENT ME
	message := messaging.GenerateRequestNotification(
		token,
		monetaryT.AmountUnit,
		monetaryT.AmountCents,
		monetaryT.Currency,
		monetaryT.From,
	)

	var str string
	str, err = mesClient.Send(ctx, message)
	if err != nil {
		log.Print(str)
	}
	return err
}
