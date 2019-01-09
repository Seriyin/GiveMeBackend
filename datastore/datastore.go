// Copyright 2015 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package datastore

import (
	"context"
	"fmt"

	"cloud.google.com/go/datastore"
	"cloud.google.com/go/storage"
)

type profileDB struct {
	client *datastore.Client
}

type ObjectData struct {
	attrs  *storage.ObjectAttrs
	handle *storage.ObjectHandle
}

// Ensure fileDB conforms to the FileDatabase interface.
var _ ProfileDatabase = &profileDB{}

// newDatastoreDB creates a new ProfileDatabase backed by Cloud Datastore.
// See the datastore and google packages for details on creating a suitable Client:
// https://godoc.org/cloud.google.com/go/datastore
func newProfileDB(client *datastore.Client) (ProfileDatabase, error) {
	ctx := context.Background()
	// Verify that we can communicate and authenticate with the datastore service.
	t, err := client.NewTransaction(ctx)
	if err != nil {
		return nil, fmt.Errorf("datastoredb: could not connect: %v", err)
	}
	if err := t.Rollback(); err != nil {
		return nil, fmt.Errorf("datastoredb: could not connect: %v", err)
	}
	return &profileDB{
		client: client,
	}, nil
}

// Close closes the database.
func (db *profileDB) Close() error {
	return db.client.Close()
}

func (db *profileDB) datastoreKey(id int64) *datastore.Key {
	return datastore.IDKey("Profile", id, nil)
}

// GetProfile retrieves a file by its ID.
func (db *profileDB) GetProfile(id int64) (*Profile, error) {
	ctx := context.Background()
	k := db.datastoreKey(id)
	profile := &SimpleProfile{}
	if err := db.client.Get(ctx, k, profile); err != nil {
		return nil, fmt.Errorf("datastoredb: could not get Profile: %v", err)
	}
	profile.ID = id
	return profile, nil
}

// AddProfile saves a given profile, assigning it a new ID.
func (db *profileDB) AddProfile(p *Profile) (id int64, err error) {
	ctx := context.Background()
	k := datastore.IncompleteKey("Profile", nil)
	k, err = db.client.Put(ctx, k, p)
	if err != nil {
		return 0, fmt.Errorf("datastoredb: could not put Profile: %v", err)
	}
	return k.ID, nil
}

// DeleteProfile removes a given profile by its ID.
func (db *profileDB) DeleteProfile(id int64) error {
	ctx := context.Background()
	k := db.datastoreKey(id)
	if err := db.client.Delete(ctx, k); err != nil {
		return fmt.Errorf("datastoredb: could not delete Profile: %v", err)
	}
	return nil
}

// UpdateProfile updates the entry for a given profile.
func (db *profileDB) UpdateProfile(p *Profile) error {
	ctx := context.Background()
	k := db.datastoreKey(p.ID)
	if _, err := db.client.Put(ctx, k, p); err != nil {
		return fmt.Errorf("datastoredb: could not update Profile: %v", err)
	}
	return nil
}

// ListFilesSharedBy returns a list of files, ordered by timestamp,
//filtered by the profile who shared the files.
func (db *profileDB) ListFilesSharedBy(userID string) (*Files, error) {
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
