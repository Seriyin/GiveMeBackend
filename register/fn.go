package register

import (
	"context"
	"github.com/Seriyin/GiveMeBackend/config/firebase"
	"github.com/Seriyin/GiveMeBackend/config/firebase/firestore"
)

var db = firebase.GetDB()

func Register(
	ctx context.Context,
	e firestore.Event,
) {
}
