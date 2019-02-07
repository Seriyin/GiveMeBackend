package acceptRefuse

import (
	"context"
	"github.com/Seriyin/GibMe-backend/config/firebase"
	"github.com/Seriyin/GibMe-backend/config/firebase/firestore"
)

var db = firebase.GetDB()

func AcceptanceOrRefusal(
	ctx context.Context,
	e firestore.Event,
) {
}
