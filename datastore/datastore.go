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
	IsBlocked(userId string) (bool, error)

	// ListProfiles returns a list of profiles.
	ListProfiles() ([]*Profile, error)

	// ListFilesSharedBy returns a list of files, ordered by timestamp,
	// filtered by the user who created the files.
	ListFilesSharedBy(userId string) (*Files, error)

	// GetProfile retrieves a profile by its ID.
	GetProfile(userId string) (*Profile, error)

	// AddProfile saves a given profile, assigning it a new ID.
	AddProfile(p *Profile) (userId string, err error)

	// DeleteProfile removes a given profile by its ID.
	DeleteProfile(userId string) error

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
func (db *firestoreDB) GetProfile(userId string) (*Profile, error) {
	ctx := context.Background()
	doc := db.client.Collection("profiles").Doc(userId)
	profile := &Profile{}
	docSnap, err := doc.Get(ctx)
	if err != nil {
		return nil, fmt.Errorf("datastoredb: could not get Profile: %v", err)
	}
	if err := docSnap.DataTo(profile); err != nil {
		return nil, fmt.Errorf("datastoredb: could not populate Profile: %v", err)
	}
	return profile, nil
}

// AddProfile saves a given profile, assigning it a new ID.
func (db *firestoreDB) AddProfile(p *Profile) (string, error) {
	ctx := context.Background()
	doc := db.client.Collection("profiles").Doc(p.UserId.Id)
	_, err := doc.Create(ctx, p)
	if err != nil {
		return "", fmt.Errorf("datastoredb: could not put Profile: %v", err)
	}
	return p.UserId.Id, nil
}

// DeleteProfile removes a given profile by its ID.
func (db *firestoreDB) DeleteProfile(userId string) error {
	ctx := context.Background()
	doc := db.client.Collection("profiles").Doc(userId)
	_, err := doc.Delete(ctx)
	if err != nil {
		return fmt.Errorf("datastoredb: could not delete Profile: %v", err)
	}
	return nil
}

// UpdateProfile updates the entry for a given profile.
func (db *firestoreDB) UpdateProfile(p *Profile) error {
	ctx := context.Background()
	doc := db.client.Collection("profiles").Doc(p.UserId.Id)
	_, err := doc.Set(ctx, p, firestore.MergeAll)
	if err != nil {
		return fmt.Errorf("datastoredb: could not update Profile: %v", err)
	}
	return nil
}

// ListFilesSharedBy returns a list of files, ordered by timestamp,
//filtered by the profile who shared the files.
func (db *firestoreDB) ListFilesSharedBy(userId string) (*Files, error) {
	return nil, fmt.Errorf("datastoredb: not implemented")
}

func (db *firestoreDB) ListProfiles() ([]*Profile, error) {
	return nil, nil
}

func (db *firestoreDB) IsBlocked(userId string) (bool, error) {
	ctx := context.Background()
	doc := db.client.Collection("blocked").Doc(userId)
	docSnap, err := doc.Get(ctx)
	if err != nil {
		return false, fmt.Errorf("datastoredb: could not find Blocked: %v", err)
	}
	var blocked []string
	err = docSnap.DataTo(&blocked)
	if err != nil {
		return false, fmt.Errorf("datastoredb: could not populate blocked array: %v", err)
	}
	isBlocked := false
	for i := 0; isBlocked || i < len(blocked); i++ {
		isBlocked = blocked[i] == userId
	}
	return isBlocked, nil
}

func (db *firestoreDB) RegenProfile(p *Profile) error {
	ctx := context.Background()
	doc := db.client.Collection("profiles").Doc(p.UserId.Id)
	_, err := doc.Set(ctx, p)
	if err != nil {
		return fmt.Errorf("datastoredb: could not put Profile: %v", err)
	}
	return nil
}
