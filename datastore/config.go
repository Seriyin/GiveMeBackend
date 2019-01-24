package datastore

import (
	"context"
	"errors"
	"log"
	"os"

	"cloud.google.com/go/firestore"
	"cloud.google.com/go/pubsub"
	"cloud.google.com/go/storage"

	"github.com/gorilla/sessions"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	DB          FirestoreDatabase
	OAuthConfig *oauth2.Config

	StorageBucket     *storage.BucketHandle
	StorageBucketName string

	SessionStore sessions.Store

	PubsubClient *pubsub.Client
)

func init() {
	var err error

	//client_id := os.Getenv("CLOUD_AUTH_ID")
	//client_secret := os.Getenv("CLOUD_AUTH_SECRET")
	cookie_secret := os.Getenv("COOKIE_SECRET")
	project_id := os.Getenv("PROJECT_ID")

	//Read credentials from local file

	// To use the in-memory test database, uncomment the next line.
	//DB = newMemoryDB()

	// [START datastore]
	// To use Cloud Datastore, uncomment the following lines and update the
	// project ID.
	// More options can be set, see the google package docs for details:
	// http://godoc.org/golang.org/x/oauth2/google
	//
	DB, err = configureFirestoreDB(project_id)
	// [END datastore]

	if err != nil {
		log.Fatal(err)
	}

	// [START storage]
	// To configure Cloud Storage, uncomment the following lines and update the
	// bucket name.
	//
	// StorageBucketName = "<your-storage-bucket>"
	// StorageBucket, err = configureStorage(StorageBucketName)
	// [END storage]

	//if err != nil {
	//	log.Fatal(err)
	//}

	// [START auth]
	// To enable user sign-in, uncomment the following lines and update the
	// Client ID and Client Secret.
	// You will also need to update OAUTH2_CALLBACK in app.yaml when pushing to
	// production.
	//
	// OAuthConfig = configureOAuthClient(client_id, client_secret)
	// [END auth]

	// [START sessions]
	// Configure storage method for session-wide information.
	// Update "something-very-secret" with a hard to guess string or byte sequence.
	cookieStore := sessions.NewCookieStore([]byte(cookie_secret))
	cookieStore.Options = &sessions.Options{
		HttpOnly: true,
	}
	SessionStore = cookieStore
	// [END sessions]

	// [START pubsub]
	// To configure Pub/Sub, uncomment the following lines and update the project ID.
	//
	// PubsubClient, err = configurePubsub("<your-project-id>")
	// [END pubsub]

	if err != nil {
		log.Fatal(err)
	}
}

func configureFirestoreDB(projectID string) (FirestoreDatabase, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return nil, err
	}
	return newFirestoreDB(client)
}

func configureStorage(bucketID string) (*storage.BucketHandle, error) {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.Bucket(bucketID), nil
}

func configurePubsub(projectID string) (*pubsub.Client, error) {
	if _, ok := DB.(*memoryDB); ok {
		return nil, errors.New("Pub/Sub worker doesn't work with the in-memory DB ")
	}

	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return nil, err
	}

	// Create the topic if it doesn't exist.
	/*
		if exists, err := client.Topic(PubsubTopicID).Exists(ctx); err != nil {
			return nil, err
		} else if !exists {
			if _, err := client.CreateTopic(ctx, PubsubTopicID); err != nil {
				return nil, err
			}
		}
	*/
	return client, nil
}

func configureOAuthClient(clientID, clientSecret string) *oauth2.Config {
	redirectURL := os.Getenv("OAUTH2_CALLBACK")
	if redirectURL == "" {
		redirectURL = "http://localhost:8080/oauth2callback"
	}
	return &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL,
		Scopes:       []string{"email", "profile"},
		Endpoint:     google.Endpoint,
	}
}
