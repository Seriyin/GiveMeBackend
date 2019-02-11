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
			
			return err
		}
		else if a == "ConfirmedTo" {

			return err
		}
	}
	return err
}
