package confirm

import (
	"context"
	"encoding/json"
	"log"

	"github.com/Seriyin/GiveMeBackend/config/datastore"
	"github.com/Seriyin/GiveMeBackend/config/firebase"
	"github.com/Seriyin/GiveMeBackend/config/firebase/firestore"
	"github.com/Seriyin/GiveMeBackend/config/firebase/messaging"
	"github.com/Seriyin/GiveMeBackend/config/firebase/paths"
)

var db = firebase.GetDB()
var mesClient = firebase.GetMessaging()

func Confirm(
	ctx context.Context,
	e firestore.Event,
) error {
	var monetaryT datastore.MonetaryTransfer          // Monetary Structure object
	err := json.Unmarshal(e.Value.Fields, &monetaryT) // Json object to Monetary Structure
	if err != nil {
		return err
	}

	profile, err := db.GetProfileByPhoneNumber(
		ctx,
		monetaryT.To,
	)
	if err != nil {
		return err
	}

	snowflake, dbPath :=
		paths.ExtractAndReplaceMethodIdAndDatePathWithSnowflake(
			profile.Id,
			e.Value.Name,
		)

	for _, a := range e.UpdateMask.FieldPaths {
		if a == "confirmedFrom" {

			err = db.UpdateMonetaryRequestConfirmedByFullPath(
				ctx,
				true, //ConfirmedFrom
				true, //ConfirmedTo
				dbPath,
				snowflake,
			)
			if err != nil {
				return err
			}

			err = produceAndSendFromNotification(
				ctx,
				profile,
				&monetaryT,
			)
			return err
		} else if a == "confirmedTo" {
			err = db.UpdateMonetaryRequestConfirmedByFullPath(
				ctx,
				false, //ConfirmedFrom
				true,  //ConfirmedTo
				dbPath,
				snowflake,
			)
			if err != nil {
				return err
			}

			err = produceAndSendToNotification(
				ctx,
				profile,
				&monetaryT,
			)
			return err
		}
	}
	return err
}

func produceAndSendFromNotification(
	ctx context.Context,
	profile *datastore.Profile,
	transfer *datastore.MonetaryTransfer,
) error {
	//generate notification message
	token := profile.Token
	message := messaging.GenerateConfirmedFromNotification(
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

func produceAndSendToNotification(
	ctx context.Context,
	profile *datastore.Profile,
	transfer *datastore.MonetaryTransfer,
) error {
	//generate notification message
	token := profile.Token
	message := messaging.GenerateConfirmedToNotification(
		token,
		transfer.AmountUnit,
		transfer.AmountCents,
		transfer.Currency,
		transfer.To,
	)

	str, err := mesClient.Send(ctx, message)
	if err != nil {
		log.Print(str)
	}
	return err
}
