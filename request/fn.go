package request

import (
	"context"
	"github.com/Seriyin/GibMe-backend/config/firebase"
	"github.com/Seriyin/GibMe-backend/config/firebase/firestore"
)

var db = firebase.GetDB()

func Request(
	ctx context.Context,
	e firestore.Event,
) {
}
