package firestore

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"

	"github.com/Seriyin/GiveMeBackend/config/datastore"
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
) (*datastore.Profile, error) {
	docs := db.client.Collection("Profiles").Where(
		"phone",
		"==",
		phoneNumber,
	).Documents(ctx)
	profile := &datastore.Profile{}
	defer docs.Stop()
	docSnap, err := docs.Next()
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

// GetProfileIdByPhoneNumber retrieves a profile's Id by its associated phone number.
func (db *firestoreDB) GetProfileIdByPhoneNumber(
	ctx context.Context,
	phoneNumber string,
) (string, error) {
	docs := db.client.Collection("Profiles").Where(
		"phone",
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
		"Profiles",
	).Doc(p.Id)
	_, err := doc.Create(ctx, p)
	if err != nil {
		return "", fmt.Errorf(
			"datastoredb: could not put Profile: %v",
			err,
		)
	}
	return p.Id, nil
}

// DeleteProfile removes a given profile by its ID.
func (db *firestoreDB) DeleteProfile(
	ctx context.Context,
	userId string,
) error {
	doc := db.client.Collection(
		"Profiles",
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
		"Profiles",
	).Doc(p.Id)
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
		"Profiles",
	).Doc(p.Id)
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
		"Blocked",
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

func (db *firestoreDB) AddMonetaryRequest(
	ctx context.Context,
	userId string,
	transfer *datastore.MonetaryRequest,
	path string,
) (string, error) {
	pathT := buildCollectionPath(userId, path)
	return db.AddMonetaryRequestByFullPath(
		ctx,
		transfer,
		pathT,
	)
}

func (db *firestoreDB) AddMonetaryRequestByFullPath(
	ctx context.Context,
	transfer *datastore.MonetaryRequest,
	fullPath string,
) (string, error) {
	doc := db.client.Collection(
		fullPath,
	).NewDoc()
	transfer.Snowflake = doc.ID
	wr, err := doc.Set(ctx, transfer)
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

func (db *firestoreDB) GetMonetaryRequestWithDate(
	ctx context.Context,
	userId string,
	date time.Time,
	snowflake string,
) (*datastore.MonetaryRequest, error) {
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
	var mon datastore.MonetaryRequest
	err = docSnap.DataTo(&mon)
	if err != nil {
		return nil, fmt.Errorf(
			"datastoredb: could not convert to monetary_transfer: %v",
			err,
		)
	}
	return &mon, nil
}

func (db *firestoreDB) GetMonetaryRequestWithDateString(
	ctx context.Context,
	userId string,
	date string,
	snowflake string,
) (*datastore.MonetaryRequest, error) {
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
	var mon datastore.MonetaryRequest
	err = docSnap.DataTo(&mon)
	if err != nil {
		return nil, fmt.Errorf(
			"datastoredb: could not convert to monetary_transfer: %v",
			err,
		)
	}
	return &mon, nil
}

func (db *firestoreDB) GetMonetaryRequestsDate(
	ctx context.Context,
	userId string,
	dateBefore time.Time,
) ([]*datastore.MonetaryRequest, error) {
	pathRoot := "MonetaryRequest/" + userId + "/"

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
	var mon datastore.MonetaryRequest
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

func (db *firestoreDB) GetMonetaryRequestsInterval(
	ctx context.Context,
	userId string,
	dateAfter time.Time,
	dateBefore time.Time,
) ([]*datastore.MonetaryRequest, error) {
	panic("implement me")
}

func (db *firestoreDB) GetMonetaryRequestsFromGroup(
	ctx context.Context,
	userId string,
	date time.Time,
	groupId int64,
) ([]*datastore.MonetaryRequest, error) {
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
	mts := make([]*datastore.MonetaryRequest, len(docs))
	for _, r := range docs {
		var mon datastore.MonetaryRequest
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

func (db *firestoreDB) GetMonetaryRequestsRecurrent(
	ctx context.Context,
	userId string,
	recurrentId int64,
) ([]*datastore.MonetaryRequest, error) {
	panic("implement me")
}

func (db *firestoreDB) SetMonetaryRequest(
	ctx context.Context,
	userId string,
	transfer *datastore.MonetaryRequest,
	path string,
) (string, error) {
	pathT := buildCollectionPath(userId, path)
	return db.SetMonetaryRequestByFullPath(
		ctx,
		transfer,
		pathT,
	)
}

func (db *firestoreDB) SetMonetaryRequestByFullPath(
	ctx context.Context,
	transfer *datastore.MonetaryRequest,
	fullPath string,
) (string, error) {
	doc := db.client.Collection(
		fullPath,
	).Doc(transfer.Snowflake)
	wr, err := doc.Set(ctx, transfer)
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

func (db *firestoreDB) SetMonetaryRequests(
	ctx context.Context,
	userId string,
	transfers []*datastore.MonetaryRequest,
	path string,
) error {
	pathT := buildCollectionPath(userId, path)
	return db.SetMonetaryRequestsByFullPath(
		ctx,
		transfers,
		pathT,
	)
}

func (db *firestoreDB) SetMonetaryRequestsByFullPath(
	ctx context.Context,
	transfers []*datastore.MonetaryRequest,
	fullPath string,
) error {
	batch := db.client.Batch()
	cl := db.client.Collection(
		fullPath,
	)
	for _, transfer := range transfers {
		doc := cl.Doc(transfer.Snowflake)
		batch.Set(doc, transfer)
	}
	_, err := batch.Commit(ctx)
	return err
}

func (db *firestoreDB) UpdateMonetaryRequestConfirmed(
	ctx context.Context,
	userId string,
	confirmedFrom bool, //If false ignore
	confirmedTo bool, //If false ignore
	path string,
	linkedId string,
) error {
	pathT := buildCollectionPath(userId, path)
	return db.UpdateMonetaryRequestConfirmedByFullPath(
		ctx,
		confirmedFrom,
		confirmedTo,
		pathT,
		linkedId,
	)
}

func (db *firestoreDB) UpdateMonetaryRequestConfirmedByFullPath(
	ctx context.Context,
	confirmedFrom bool, //If false ignore
	confirmedTo bool, //If false ignore
	fullPath string,
	linkedId string,
) error {
	doc := db.client.Collection(
		fullPath,
	).Doc(linkedId)
	wr, err := doc.Update(ctx, []firestore.Update{
		{Path: "confirmedFrom", Value: confirmedFrom},
		{Path: "confirmedTo", Value: confirmedTo},
	})
	if err != nil {
		return fmt.Errorf(
			"datastoredb: failed to update monetary transfer %v in %v: %v %v",
			linkedId,
			fullPath,
			wr,
			err,
		)
	}
	return nil
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
	str.WriteString("MonetaryRequest/")
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
