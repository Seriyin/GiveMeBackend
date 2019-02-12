package remind

import (
	"context"
	"github.com/Seriyin/GiveMeBackend/config/firebase"
	"github.com/Seriyin/GiveMeBackend/config/firebase/firestore"
)

var db = firebase.GetDB()

func Remind(
	ctx context.Context,
	e firestore.Event,
) error {
	return nil
}
