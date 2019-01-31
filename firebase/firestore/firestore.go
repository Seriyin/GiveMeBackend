package firestore

import (
	"cloud.google.com/go/firestore"
	"context"
	"fmt"
	"time"

	"github.com/Seriyin/GibMe-backend/datastore"
)

type firestoreDB struct {
	client *firestore.Client
}

// Ensure firestoreDB conforms to the FirestoreDatabase interface.
var (
	_ datastore.FirestoreDatabase = &firestoreDB{}
)

// NewFirestoreDB creates a new FirestoreDatabase backed by Cloud Firestore.
// See the firestore and google packages for details on creating a suitable Client:
// https://godoc.org/cloud.google.com/go/firestore
func NewFirestoreDB(
	client *firestore.Client,
) (datastore.FirestoreDatabase, error) {
	ctx := context.Background()
	// Verify that we can communicate and authenticate with the datastore service.
	err := client.RunTransaction(
		ctx,
		func(
			i context.Context,
			transaction *firestore.Transaction,
		) error {
			return nil
		},
	)
	if err != nil {
		return nil, fmt.Errorf(
			"datastoredb: could not connect: %v",
			err,
		)
	}
	return &firestoreDB{
		client: client,
	}, nil
}

// Close closes the database.
func (db *firestoreDB) Close() error {
	return db.client.Close()
}

// GetProfile retrieves a profile by its ID.
func (db *firestoreDB) GetProfile(
	userId string,
) (*datastore.Profile, error) {
	ctx := context.Background()
	doc := db.client.Collection("profiles").Doc(userId)
	profile := &datastore.Profile{}
	docSnap, err := doc.Get(ctx)
	if err != nil {
		return nil, fmt.Errorf(
			"datastoredb: could not get Profile: %v",
			err,
		)
	}
	if err := docSnap.DataTo(profile); err != nil {
		return nil, fmt.Errorf(
			"datastoredb: could not populate Profile: %v",
			err,
		)
	}
	return profile, nil
}

// AddProfile saves a given profile, assigning it a new ID.
func (db *firestoreDB) AddProfile(p *datastore.Profile) (string, error) {
	ctx := context.Background()
	doc := db.client.Collection(
		"profiles",
	).Doc(p.UserId.Id)
	_, err := doc.Create(ctx, p)
	if err != nil {
		return "", fmt.Errorf(
			"datastoredb: could not put Profile: %v",
			err,
		)
	}
	return p.UserId.Id, nil
}

// DeleteProfile removes a given profile by its ID.
func (db *firestoreDB) DeleteProfile(userId string) error {
	ctx := context.Background()
	doc := db.client.Collection(
		"profiles",
	).Doc(userId)
	_, err := doc.Delete(ctx)
	if err != nil {
		return fmt.Errorf(
			"datastoredb: could not delete Profile: %v",
			err,
		)
	}
	return nil
}

// UpdateProfile updates the entry for a given profile.
func (db *firestoreDB) UpdateProfile(p *datastore.Profile) error {
	ctx := context.Background()
	doc := db.client.Collection(
		"profiles",
	).Doc(p.UserId.Id)
	_, err := doc.Set(ctx, p, firestore.MergeAll)
	if err != nil {
		return fmt.Errorf(
			"datastoredb: could not update Profile: %v",
			err,
		)
	}
	return nil
}

// ListFilesSharedBy returns a list of files, ordered by timestamp,
//filtered by the profile who shared the files.
func (db *firestoreDB) ListFilesSharedBy(
	userId string,
) (*datastore.Files, error) {
	return nil, fmt.Errorf(
		"datastoredb: not implemented",
	)
}

func (db *firestoreDB) ListProfiles() ([]*datastore.Profile, error) {
	return nil, nil
}

func (db *firestoreDB) IsBlocked(
	userId string,
) (bool, error) {
	ctx := context.Background()
	doc := db.client.Collection(
		"blocked",
	).Doc(userId)
	docSnap, err := doc.Get(ctx)
	if err != nil {
		return false, fmt.Errorf(
			"datastoredb: could not find Blocked: %v",
			err,
		)
	}
	var blocked []string
	err = docSnap.DataTo(&blocked)
	if err != nil {
		return false, fmt.Errorf(
			"datastoredb: could not populate blocked array: %v",
			err,
		)
	}
	isBlocked := false
	for i := 0; isBlocked || i < len(blocked); i++ {
		isBlocked = blocked[i] == userId
	}
	return isBlocked, nil
}

func (db *firestoreDB) RegenProfile(p *datastore.Profile) error {
	ctx := context.Background()
	doc := db.client.Collection(
		"profiles",
	).Doc(p.UserId.Id)
	_, err := doc.Set(ctx, p)
	if err != nil {
		return fmt.Errorf(
			"datastoredb: could not put Profile: %v",
			err,
		)
	}
	return nil
}

// Event is the payload of a Firestore event.
type Event struct {
	OldValue   Value `json:"oldValue"`
	Value      Value `json:"value"`
	UpdateMask struct {
		FieldPaths []string `json:"fieldPaths"`
	} `json:"updateMask"`
}

// Value holds Firestore fields.
type Value struct {
	CreateTime time.Time `json:"createTime"`
	// Fields is the data for this value. The type depends on the format of your
	// database. Log the interface{} value and inspect the result to see a JSON
	// representation of your database fields.
	Fields     interface{} `json:"fields"`
	Name       string      `json:"name"`
	UpdateTime time.Time   `json:"updateTime"`
}
