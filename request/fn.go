package request

import (
	"context"
	"encoding/json"
	"github.com/Seriyin/GiveMeBackend/config/datastore"
	"github.com/Seriyin/GiveMeBackend/config/firebase"
	"github.com/Seriyin/GiveMeBackend/config/firebase/firestore"
	"github.com/Seriyin/GiveMeBackend/config/firebase/messaging"
	"github.com/Seriyin/GiveMeBackend/config/firebase/paths"
	"log"
)

var db = firebase.GetDB()
var mesClient = firebase.GetMessaging()

func Request(
	ctx context.Context,
	e firestore.Event,
) error {
	var monetaryT datastore.MonetaryRequest           // Monetary Structure object
	err := json.Unmarshal(e.Value.Fields, &monetaryT) // Json object to Monetary Structure

	log.Print("Attempted unmarshal")
	if err != nil {
		return err
	}

	// If no profile can be gathered, the user may not exist.
	// Either by network error or profile not existing, must return.
	profile, err := db.GetProfileByPhoneNumber(
		ctx,
		monetaryT.To,
	)

	log.Print("Attempted profile grab")
	if err != nil {
		return err
	}

	dbPath := paths.ExtractAndReplaceMethodIdAndDatePath(
		profile.Id,
		e.Value.Name,
	)

	log.Printf("Extracted db path: %v", dbPath)
	_, err = db.SetMonetaryRequest(
		ctx,
		profile.Id,
		&monetaryT,
		dbPath,
	)
	if err != nil {
		return err
	}

	err = produceAndSendNotification(
		ctx,
		profile,
		&monetaryT,
	)

	log.Print("Attempted produce send notification")
	return err
}

func produceAndSendNotification(
	ctx context.Context,
	profile *datastore.Profile,
	transfer *datastore.MonetaryRequest,
) error {
	//generate notification message
	token := profile.Token
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
