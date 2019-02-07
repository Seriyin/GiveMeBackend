package firestore

import (
	"cloud.google.com/go/firestore"
	"context"
	"fmt"
	"google.golang.org/api/iterator"
	"log"
	"time"

	"github.com/Seriyin/GibMe-backend/config/datastore"
)

type firestoreDB struct {
	client *firestore.Client
}

// Ensure firestoreDB conforms to the GiveMeDatabase interface.
var (
	_ datastore.GiveMeDatabase = &firestoreDB{}
)

// NewFirestoreDB creates a new GiveMeDatabase backed by Cloud Firestore.
// See the firestore and google packages for details on creating a suitable Client:
// https://godoc.org/cloud.google.com/go/firestore
func NewFirestoreDB(
	client *firestore.Client,
) (datastore.GiveMeDatabase, error) {
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

func (db *firestoreDB) IsBlocked(
	userId string,
	blocked string,
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
	var blockedP []string
	err = docSnap.DataTo(&blockedP)
	if err != nil {
		return false, fmt.Errorf(
			"datastoredb: could not populate blocked array: %v",
			err,
		)
	}
	isBlocked := false
	for i := 0; isBlocked || i < len(blockedP); i++ {
		isBlocked = blockedP[i] == blocked
	}
	return isBlocked, nil
}

func (db *firestoreDB) GetMonetaryTransfer(
	userId string,
	snowflake string,
) (*datastore.MonetaryTransfer, error) {
	ctx := context.Background()
	doc := db.client.Collection(
		"monetary_transfer/" + userId,
	).Doc(snowflake)
	docSnap, err := doc.Get(ctx)
	if err != nil {
		return nil, fmt.Errorf(
			"datastoredb: could not get monetary_transfer: %v",
			err,
		)
	}
	var mon datastore.MonetaryTransfer
	err = docSnap.DataTo(&mon)
	if err != nil {
		return nil, fmt.Errorf(
			"datastoredb: could not convert to monetary_transfer: %v",
			err,
		)
	}
	return &mon, nil
}

func (db *firestoreDB) GetMonetaryTransfersDate(
	userId string,
	dateBefore time.Time,
) ([]*datastore.MonetaryTransfer, error) {
	pathRoot := "monetary_transfer/" + userId

	now := time.Now()
	dt := time.Date(
		now.Year(),
		now.Month(),
		now.Day(),
		0,
		0,
		0,
		0,
		time.UTC,
	)
	dtBefore := time.Date(
		dateBefore.Year(),
		dateBefore.Month(),
		dateBefore.Day(),
		0,
		0,
		0,
		0,
		time.UTC,
	)
	return nil, fmt.Errorf(
		"datastoredb: not implemented yet %v %v %v",
		pathRoot,
		dt,
		dtBefore,
	)
	/*
		for ; dt.After(dtBefore); dt.AddDate(0, 0, -1) {
			go collectFromDate(pathRoot, dt, db)
		}
		return &mon, nil
	*/
}

func collectFromDate(
	pathRoot string,
	dt time.Time,
	db *firestoreDB,
) {
	ctx := context.Background()
	path := pathRoot + "/" + dt.Format("Mon, 02 Jan 2006")
	docIter := db.client.Collection(path).Documents(ctx)
	doc, err := docIter.Next()
	var mon datastore.MonetaryTransfer
	for ; err != iterator.Done; doc, err = docIter.Next() {
		err = doc.DataTo(&mon)
		if err != nil {
			log.Print(
				fmt.Errorf(
					"datastoredb: could not convert to monetary_transfer: %v",
					err,
				),
			)
		}
	}
}

func (db *firestoreDB) GetMonetaryTransfersInterval(
	userId string,
	dateAfter time.Time,
	dateBefore time.Time,
) ([]*datastore.MonetaryTransfer, error) {
	panic("implement me")
}

func (db *firestoreDB) GetMonetaryTransfersFromGroup(
	userId string,
	groupId int64,
) ([]*datastore.MonetaryTransfer, error) {
	panic("implement me")
}

func (db *firestoreDB) GetMonetaryTransfersRecurrent(
	userId string,
	recurrentId int64,
) ([]*datastore.MonetaryTransfer, error) {
	panic("implement me")
}

func (db *firestoreDB) SetMonetaryTransfer(
	userId string,
	transfer *datastore.MonetaryTransfer,
	path string,
) error {
	panic("implement me")
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
