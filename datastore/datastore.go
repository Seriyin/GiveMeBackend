// Copyright 2015 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package datastore

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	"cloud.google.com/go/storage"
)

type FirestoreDatabase interface {
	IsBlocked(userID string) (bool, error)

	// ListProfiles returns a list of profiles.
	ListProfiles() ([]*Profile, error)

	// ListFilesSharedBy returns a list of files, ordered by timestamp,
	// filtered by the user who created the files.
	ListFilesSharedBy(userID string) (*Files, error)

	// GetProfile retrieves a profile by its ID.
	GetProfile(userID string) (*Profile, error)

	// AddProfile saves a given profile, assigning it a new ID.
	AddProfile(p *Profile) (userID string, err error)

	// DeleteProfile removes a given profile by its ID.
	DeleteProfile(userID string) error

	// UpdateProfile updates the entry for a given profile.
	UpdateProfile(p *Profile) error

	// RegenProfile discards and reinitializes known profile
	RegenProfile(p *Profile) error

	// Close closes the database, freeing up any available resources.
	Close() error
}


type firestoreDB struct {
	client *firestore.Client
}

type ObjectData struct {
	attrs  *storage.ObjectAttrs
	handle *storage.ObjectHandle
}

// Ensure firestoreDB conforms to the FirestoreDatabase interface.
var _ FirestoreDatabase = &firestoreDB{}


// newFirestoreDB creates a new FirestoreDatabase backed by Cloud Firestore.
// See the firestore and google packages for details on creating a suitable Client:
// https://godoc.org/cloud.google.com/go/firestore
func newFirestoreDB(client *firestore.Client) (FirestoreDatabase, error) {
	ctx := context.Background()
	// Verify that we can communicate and authenticate with the datastore service.
	err := client.RunTransaction(ctx, func(i context.Context, transaction *firestore.Transaction) error {
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("datastoredb: could not connect: %v", err)
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
func (db *firestoreDB) GetProfile(userID string) (*Profile, error) {
	ctx := context.Background()
	doc := db.client.Collection("profiles").Doc(userID)
	profile := &Profile{}
	if docSnap, err := doc.Get(ctx); err != nil {
		return nil, fmt.Errorf("datastoredb: could not get Profile: %v", err)
	}
	dataMap := docSnap.
	profile.UserId = data
	return profile, nil
}

// AddProfile saves a given profile, assigning it a new ID.
func (db *firestoreDB) AddProfile(p *Profile) (id int64, err error) {
	ctx := context.Background()
	k := db.client.Doc("Profile")
	k.
	k, err = db.client.Put(ctx, k, p)
	if err != nil {
		return 0, fmt.Errorf("datastoredb: could not put Profile: %v", err)
	}
	return k.ID, nil
}

// DeleteProfile removes a given profile by its ID.
func (db *firestoreDB) DeleteProfile(id int64) error {
	ctx := context.Background()
	k := db.datastoreKey(id)
	if err := db.client.Delete(ctx, k); err != nil {
		return fmt.Errorf("datastoredb: could not delete Profile: %v", err)
	}
	return nil
}

// UpdateProfile updates the entry for a given profile.
func (db *firestoreDB) UpdateProfile(p *Profile) error {
	ctx := context.Background()
	k := db.datastoreKey(p.UserId.Id)
	if _, err := db.client.Put(ctx, k, p); err != nil {
		return fmt.Errorf("datastoredb: could not update Profile: %v", err)
	}
	return nil
}

// ListFilesSharedBy returns a list of files, ordered by timestamp,
//filtered by the profile who shared the files.
func (db *firestoreDB) ListFilesSharedBy(userID string) (*Files, error) {
	ctx := context.Background()
	if userID == "" {
		return db.ListBooks()
	}

	files := make([]*storage.ObjectHandle, 0)
	q := datastore.NewQuery("File").
		Filter("CreatedByID =", userID)

	keys, err := db.client.GetAll(ctx, q, &files)

	if err != nil {
		return nil, fmt.Errorf("datastoredb: could not list files: %v", err)
	}

	for i, k := range keys {
		files[i].ID = k.ID
	}

	return books, nil
}

func (db *firestoreDB) ListProfiles() ([]*Profile, error) {
	return nil, nil
}
