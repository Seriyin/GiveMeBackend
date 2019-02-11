package confirm

import (
	"context"

	"github.com/Seriyin/GibMe-backend/config/datastore"
	"github.com/Seriyin/GibMe-backend/config/firebase"
	"github.com/Seriyin/GibMe-backend/config/firebase/firestore"
	"github.com/Seriyin/GibMe-backend/config/firebase/messaging"
	"github.com/Seriyin/GibMe-backend/config/firebase/paths"
)

var db = firebase.GetDB()

func Confirm(
	ctx context.Context,
	e firestore.Event,
) {
	var monetaryT datastore.MonetaryTransfer          // Monetary Structure object
	err := json.Unmarshal(e.Value.Fields, &monetaryT) // Json object to Monetary Structure
	if err != nil {
		return err
	}
	for _, a := range e.UpdateMask {
		if a == "ConfirmedFrom" {
			profile, err := db.GetProfileByPhoneNumber(
				ctx,
				monetaryT.To,
			)
			if err != nil {
				return err
			}
		
			dbPath := paths.ExtractAndReplaceMethodIdAndDatePath(
				profile.Id,
				e.Value.Name,
			)
		
			_, err = db.UpdateMonetaryTransferConfirmed(
				ctx,
				profile.Id,
				true, //ConfirmedFrom
				true, //ConfirmedTo
				dbPath,
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
		}
		else if a == "ConfirmedTo" {
			profile, err := db.GetProfileByPhoneNumber(
				ctx,
				monetaryT.To,
			)
			if err != nil {
				return err
			}
		
			dbPath := paths.ExtractAndReplaceMethodIdAndDatePath(
				profile.Id,
				e.Value.Name,
			)
		
			_, err = db.UpdateMonetaryTransferConfirmed(
				ctx,
				profile.Id,
				false, //ConfirmedFrom
				true, //ConfirmedTo
				dbPath,
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
