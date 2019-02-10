package firestore

import (
	"cloud.google.com/go/firestore"
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/api/iterator"
	"log"
	"strings"
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
	ctx context.Context,
	userId string,
) (*datastore.Profile, error) {
	doc := db.client.Collection("Profiles").Doc(userId)
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

// GetProfileByPhoneNumber retrieves a profile by its associated phone number.
func (db *firestoreDB) GetProfileByPhoneNumber(
	ctx context.Context,
	phoneNumber string,
) (string, *datastore.Profile, error) {
	docs := db.client.Collection("Profiles").Where("id", "==", phoneNumber).Documents(ctx)
	profile := &datastore.Profile{}
	defer docs.Stop()
	docSnap, err := docs.Next()
	if err != nil {
		return "", nil, fmt.Errorf(
			"datastoredb: could not get Profile: %v",
			err,
		)
	}
	if err := docSnap.DataTo(profile); err != nil {
		return "", nil, fmt.Errorf(
			"datastoredb: could not populate Profile: %v",
			err,
		)
	}
	return docSnap.Ref.ID, profile, nil
}

// GetProfileIdByPhoneNumber retrieves a profile's Id by its associated phone number.
func (db *firestoreDB) GetProfileIdByPhoneNumber(
	ctx context.Context,
	phoneNumber string,
) (string, error) {
	docs := db.client.Collection("Profiles").Where(
		"id",
		"==",
		phoneNumber,
	).Documents(ctx)
	defer docs.Stop()
	docSnap, err := docs.Next()
	if err != nil {
		return "", fmt.Errorf(
			"datastoredb: could not get Profile: %v",
			err,
		)
	}
	return docSnap.Ref.ID, nil
}

// AddProfile saves a given profile, assigning it a new ID.
func (db *firestoreDB) AddProfile(
	ctx context.Context,
	p *datastore.Profile,
) (string, error) {
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
func (db *firestoreDB) DeleteProfile(
	ctx context.Context,
	userId string,
) error {
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
func (db *firestoreDB) UpdateProfile(
	ctx context.Context,
	p *datastore.Profile,
) error {
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
	ctx context.Context,
	userId string,
) (*datastore.Files, error) {
	return nil, fmt.Errorf(
		"datastoredb: not implemented",
	)
}

func (db *firestoreDB) RegenProfile(
	ctx context.Context,
	p *datastore.Profile,
) error {
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
	ctx context.Context,
	userId string,
	blocked string,
) (bool, error) {
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

func (db *firestoreDB) GetMonetaryTransferWithDate(
	ctx context.Context,
	userId string,
	date time.Time,
	snowflake string,
) (*datastore.MonetaryTransfer, error) {
	dt := time.Date(date.Year(), date.Month(), 1, 0, 0, 0, 0, time.UTC)
	doc := db.client.Collection(
		buildCollectionPathWithDate(userId, dt),
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

func (db *firestoreDB) GetMonetaryTransferWithDateString(
	ctx context.Context,
	userId string,
	date string,
	snowflake string,
) (*datastore.MonetaryTransfer, error) {
	doc := db.client.Collection(
		buildCollectionPath(userId, date),
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
	ctx context.Context,
	userId string,
	dateBefore time.Time,
) ([]*datastore.MonetaryTransfer, error) {
	pathRoot := "MonetaryTransfer/" + userId + "/"

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
	ctx context.Context,
	pathRoot string,
	dt time.Time,
	db *firestoreDB,
) {
	path := pathRoot + dt.Format("2006-01")
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
	ctx context.Context,
	userId string,
	dateAfter time.Time,
	dateBefore time.Time,
) ([]*datastore.MonetaryTransfer, error) {
	panic("implement me")
}

func (db *firestoreDB) GetMonetaryTransfersFromGroup(
	ctx context.Context,
	userId string,
	date time.Time,
	groupId int64,
) ([]*datastore.MonetaryTransfer, error) {
	dt := time.Date(date.Year(), date.Month(), 1, 0, 0, 0, 0, time.UTC)
	docs, err := db.client.Collection(
		buildCollectionPathWithDate(userId, dt),
	).Where("groupId", "==", groupId).Documents(ctx).GetAll()
	if err != nil {
		return nil, fmt.Errorf(
			"datastoredb: could not get MonetaryTransfers: %v",
			err,
		)
	}
	mts := make([]*datastore.MonetaryTransfer, len(docs))
	for _, r := range docs {
		var mon datastore.MonetaryTransfer
		err = r.DataTo(&mon)
		if err != nil {
			return nil, fmt.Errorf(
				"datastoredb: could not convert to monetary_transfer: %v",
				err,
			)
		}
		mts = append(mts, &mon)
	}
	return mts, nil
}

func (db *firestoreDB) GetMonetaryTransfersRecurrent(
	ctx context.Context,
	userId string,
	recurrentId int64,
) ([]*datastore.MonetaryTransfer, error) {
	panic("implement me")
}

func (db *firestoreDB) SetMonetaryTransfer(
	ctx context.Context,
	userId string,
	transfer *datastore.MonetaryTransfer,
	path string,
) (string, error) {
	pathT := buildCollectionPath(userId, path)
	return db.SetMonetaryTransferByFullPath(
		ctx,
		transfer,
		pathT,
	)
}

func (db *firestoreDB) SetMonetaryTransferByFullPath(
	ctx context.Context,
	transfer *datastore.MonetaryTransfer,
	fullPath string,
) (string, error) {
	doc, wr, err := db.client.Collection(
		fullPath,
	).Add(ctx, transfer)
	if err != nil {
		return "", fmt.Errorf(
			"datastoredb: failed to add monetary transfer in %v: %v %v",
			fullPath,
			wr,
			err,
		)
	}
	return doc.ID, err
}

func (db *firestoreDB) SetMonetaryTransfers(
	ctx context.Context,
	userId string,
	transfers []*datastore.MonetaryTransfer,
	path string,
) error {
	pathT := buildCollectionPath(userId, path)
	return db.SetMonetaryTransfersByFullPath(
		ctx,
		transfers,
		pathT,
	)
}

func (db *firestoreDB) SetMonetaryTransfersByFullPath(
	ctx context.Context,
	transfers []*datastore.MonetaryTransfer,
	fullPath string,
) error {
	batch := db.client.Batch()
	cl := db.client.Collection(
		fullPath,
	)
	for _, transfer := range transfers {
		doc := cl.NewDoc()
		batch.Set(doc, transfer)
	}
	_, err := batch.Commit(ctx)
	return err
}

func buildCollectionPathWithDate(
	userId string,
	date time.Time,
) string {
	dateString := date.Format("2006-01")
	return buildCollectionPath(userId, dateString)
}

func buildCollectionPath(
	userId string,
	path string,
) string {
	str := strings.Builder{}
	str.Grow(len(userId) + 20)
	str.WriteString("MoneyTransfer/")
	str.WriteString(userId)
	str.WriteByte('/')
	str.WriteString(path)
	return str.String()
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
	// database.
	Fields     json.RawMessage `json:"fields"`
	Name       string          `json:"name"`
	UpdateTime time.Time       `json:"updateTime"`
}
