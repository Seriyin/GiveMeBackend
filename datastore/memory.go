// Copyright 2015 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package datastore

import (
	"errors"
	"fmt"
	"sync"
)

// Ensure memoryDB conforms to the ProfileDatabase interface.
var _ FirestoreDatabase = &memoryDB{}

// memoryDB is a simple in-memory persistence layer for profiles.
type memoryDB struct {
	mutex    sync.Mutex
	profiles map[string]*Profile // maps from profile ID to profile.
	files    map[string]*Files
	blocked  map[string][]string
}

func newMemoryDB() *memoryDB {
	return &memoryDB{
		profiles: make(map[string]*Profile),
		files:    make(map[string]*Files),
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
func (db *memoryDB) GetProfile(id string) (*Profile, error) {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	profile, ok := db.profiles[id]
	if !ok {
		return nil, fmt.Errorf("memorydb: profile not found with ID %d", id)
	}
	return profile, nil
}

// AddProfile saves a given profile, assigning it a new ID.
func (db *memoryDB) AddProfile(p *Profile) (id string, err error) {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	db.profiles[p.UserId.Id] = p

	return p.UserId.Id, nil
}

// DeleteProfile removes a given profile by its ID.
func (db *memoryDB) DeleteProfile(id string) error {
	if id == "" {
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
	if p.UserId.Id == "" {
		return errors.New("memorydb: profile with unassigned ID passed into updateProfile")
	}

	db.mutex.Lock()
	defer db.mutex.Unlock()

	db.profiles[p.UserId.Id] = p
	return nil
}

// ListFilesSharedBy returns a list of files, ordered by timestamp,
// filtered by the profile who created the files.
func (db *memoryDB) ListFilesSharedBy(userID string) (*Files, error) {
	if userID == "" {
		return nil, errors.New("empty userID for files shared")
	}

	//	db.mutex.Lock()
	//	defer db.mutex.Unlock()

	return nil, nil
}

func (db *memoryDB) ListProfiles() ([]*Profile, error) {
	profiles := make([]*Profile, len(db.profiles))
	i := 0
	for _, p := range db.profiles {
		profiles[i] = p
		i++
	}
	return profiles, nil
}

func (db *memoryDB) IsBlocked(userId string) (bool, error) {
	isBlocked := false
	blocked := db.blocked[userId]
	for i := 0; isBlocked || i < len(blocked); i++ {
		isBlocked = blocked[i] == userId
	}
	return isBlocked, nil
}

func (db *memoryDB) RegenProfile(p *Profile) error {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	db.profiles[p.UserId.Id] = p

	return nil
}
