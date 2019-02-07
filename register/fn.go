package register

import (
	"context"
	"github.com/Seriyin/GibMe-backend/config/firebase"
	"github.com/Seriyin/GibMe-backend/config/firebase/firestore"
)

var db = firebase.GetDB()

func Register(
	ctx context.Context,
	e firestore.Event,
) {
}
