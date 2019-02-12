package acceptRefuse

import (
	"context"
	"github.com/Seriyin/GiveMeBackend/config/firebase"
	"github.com/Seriyin/GiveMeBackend/config/firebase/firestore"
	"log"
)

var db = firebase.GetDB()

func AcceptanceOrRefusal(
	ctx context.Context,
	e firestore.Event,
) error {
	log.Printf(e.Value.Name)
	return nil
}
