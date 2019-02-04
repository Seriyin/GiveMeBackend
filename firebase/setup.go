package firebase

import (
	"context"
	"firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"github.com/Seriyin/GibMe-backend/datastore"
	"github.com/Seriyin/GibMe-backend/firebase/firestore"
	"log"
)

var (
	app *firebase.App
)

func init() {
	var err error
	app, err = firebase.NewApp(
		context.Background(),
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}
}

func GetDB() datastore.GiveMeDatabase {
	db, err := app.Firestore(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fdb, err := firestore.NewFirestoreDB(db)
	if err != nil {
		log.Fatal(err)
	}
	return fdb
}

func GetMessaging() *messaging.Client {
	mes, err := app.Messaging(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	return mes
}
