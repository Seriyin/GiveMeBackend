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

	// If no profile can be gathered, the user may not exist.
	// Either by network error or profile not existing, must return.
	id, profile, err := db.GetProfileByPhoneNumber(ctx, monetaryT.To)
	if err != nil {
		return err
	}

	err = extractAndReplaceMonetaryTransfer(ctx, id, &monetaryT, e.Value.Name)
	if err != nil {
		return err
	}

	err = produceAndSendNotification(ctx, profile, &monetaryT)
	return err
}

func extractAndReplaceMonetaryTransfer(
	ctx context.Context,
	id string,
	mon *datastore.MonetaryTransfer,
	networkPath string,
) error {

	// Use the debtor id to split the path
	// so it gets the path next to the creditor ID
	// to be replaced by the debtor ID, so it inserts on the correct collection.

	//Split at gcpstuff in [0] and full db path in [1].
	path := strings.Split(networkPath, "/documents/")[1]
	//Split by collection and extract [1] which should be fromId.
	fromId := strings.Split(path, "/")[1]
	dbPath := strings.Replace(path, fromId, id, 1)

	_, err := db.SetMonetaryTransfer(ctx, id, mon, dbPath)
	return err
}

func produceAndSendNotification(
	ctx context.Context,
	profile *datastore.Profile,
	transfer *datastore.MonetaryTransfer,
) error {
	//generate notification message
	token := profile.UserId.Token
	message := messaging.GenerateRequestNotification(
		token,
		transfer.AmountUnit,
		transfer.AmountCents,
		transfer.Currency,
		transfer.From,
	)

	str, err := mesClient.Send(ctx, message)
	if err != nil {
		log.Print(str)
	}
	return err
}
