// Copyright 2015 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package datastore

import (
	"errors"
	"fmt"
	"sort"
	"sync"
)

// Ensure memoryDB conforms to the ProfileDatabase interface.
var _ ProfileDatabase = &memoryDB{}

// memoryDB is a simple in-memory persistence layer for profiles.
type memoryDB struct {
	mutex    sync.Mutex
	nextId   int64              // next ID to assign to a profile.
	profiles map[int64]*Profile // maps from profile ID to profile.
	files    map[int64]*Files
}

func newMemoryDB() *memoryDB {
	return &memoryDB{
		profiles: make(map[int64]*Profile),
		nextId:   1,
	}
}

// Close closes the database.
func (db *memoryDB) Close() error {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	db.profiles = nil

	return nil
}

// GetProfile retrieves a profile by its ID.
func (db *memoryDB) GetProfile(id int64) (*Profile, error) {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	profile, ok := db.profiles[id]
	if !ok {
		return nil, fmt.Errorf("memorydb: profile not found with ID %d", id)
	}
	return profile, nil
}

// AddProfile saves a given profile, assigning it a new ID.
func (db *memoryDB) AddProfile(p *Profile) (id int64, err error) {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	p.Id = db.nextId
	db.profiles[p.Id] = p

	db.nextId++

	return p.Id, nil
}

// DeleteProfile removes a given profile by its ID.
func (db *memoryDB) DeleteProfile(id int64) error {
	if id == 0 {
		return errors.New("memorydb: profile with unassigned ID passed into deleteProfile")
	}

	db.mutex.Lock()
	defer db.mutex.Unlock()

	if _, ok := db.profiles[id]; !ok {
		return fmt.Errorf("memorydb: could not delete profile with ID %d, does not exist", id)
	}
	delete(db.profiles, id)
	return nil
}

// UpdateProfile updates the entry for a given profile.
func (db *memoryDB) UpdateProfile(p *Profile) error {
	if p.Id == 0 {
		return errors.New("memorydb: profile with unassigned ID passed into updateProfile")
	}

	db.mutex.Lock()
	defer db.mutex.Unlock()

	db.profiles[p.Id] = p
	return nil
}

// ListFilesSharedBy returns a list of files, ordered by timestamp,
// filtered by the profile who created the files.
func (db *memoryDB) ListFilesSharedBy(userID string) (*Files, error) {
	if userID == "" {
		return db.ListBooks()
	}

	db.mutex.Lock()
	defer db.mutex.Unlock()

	var files Files
	for _, p := range db.profiles[userId] {
		if b.CreatedByID == userID {
			books = append(books, b)
		}
	}

	sort.Sort(booksByTitle(books))
	return books, nil
}
