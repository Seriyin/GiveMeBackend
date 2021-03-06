// Copyright 2015 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package datastore

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

// Ensure memoryDB conforms to the GiveMeDatabase interface.
var _ GiveMeDatabase = &memoryDB{}

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
func (db *memoryDB) GetProfile(
	ctx context.Context,
	id string,
) (*Profile, error) {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	profile, ok := db.profiles[id]
	if !ok {
		return nil, fmt.Errorf("memorydb: profile not found with ID %d", id)
	}
	return profile, nil
}

func (db *memoryDB) GetProfileIdByPhoneNumber(
	ctx context.Context,
	phoneNumber string,
) (string, error) {
	panic("implement me")
}

func (db *memoryDB) GetProfileByPhoneNumber(
	ctx context.Context,
	phoneNumber string,
) (*Profile, error) {
	panic("implement me")
}

// AddProfile saves a given profile, assigning it a new ID.
func (db *memoryDB) AddProfile(
	ctx context.Context,
	p *Profile,
) (id string, err error) {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	db.profiles[p.Id] = p

	return p.Id, nil
}

// DeleteProfile removes a given profile by its ID.
func (db *memoryDB) DeleteProfile(
	ctx context.Context,
	id string,
) error {
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
func (db *memoryDB) UpdateProfile(
	ctx context.Context,
	p *Profile,
) error {
	if p.Id == "" {
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
		return nil, errors.New("empty userID for files shared")
	}

	//	db.mutex.Lock()
	//	defer db.mutex.Unlock()

	return nil, nil
}

func (db *memoryDB) RegenProfile(
	ctx context.Context,
	p *Profile,
) error {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	db.profiles[p.Id] = p

	return nil
}

func (db *memoryDB) IsBlocked(
	ctx context.Context,
	userId string,
	blocked string,
) (bool, error) {
	isBlocked := false
	blockedP := db.blocked[userId]
	for i := 0; isBlocked || i < len(blockedP); i++ {
		isBlocked = blockedP[i] == blocked
	}
	return isBlocked, nil
}

func (db *memoryDB) AddMonetaryRequest(
	ctx context.Context,
	userId string,
	transfer *MonetaryRequest,
	path string,
) (string, error) {
	panic("implement me")
}

func (db *memoryDB) AddMonetaryRequestByFullPath(
	ctx context.Context,
	transfer *MonetaryRequest,
	fullPath string,
) (string, error) {
	panic("implement me")
}

func (db *memoryDB) GetMonetaryRequestWithDate(
	ctx context.Context,
	userId string,
	date time.Time,
	snowflake string,
) (*MonetaryRequest, error) {
	panic("implement me")
}

func (db *memoryDB) GetMonetaryRequestWithDateString(
	ctx context.Context,
	userId string,
	date string,
	snowflake string,
) (*MonetaryRequest, error) {
	panic("implement me")
}

func (db *memoryDB) GetMonetaryRequestsDate(
	ctx context.Context,
	userId string,
	dateBefore time.Time,
) ([]*MonetaryRequest, error) {
	panic("implement me")
}

func (db *memoryDB) GetMonetaryRequestsInterval(
	ctx context.Context,
	userId string,
	dateAfter time.Time,
	dateBefore time.Time,
) ([]*MonetaryRequest, error) {
	panic("implement me")
}

func (db *memoryDB) GetMonetaryRequestsFromGroup(
	ctx context.Context,
	userId string,
	date time.Time,
	groupId int64,
) ([]*MonetaryRequest, error) {
	panic("implement me")
}

func (db *memoryDB) GetMonetaryRequestsRecurrent(
	ctx context.Context,
	userId string,
	recurrentId int64,
) ([]*MonetaryRequest, error) {
	panic("implement me")
}

func (db *memoryDB) SetMonetaryRequest(
	ctx context.Context,
	userId string,
	transfer *MonetaryRequest,
	path string,
) (string, error) {
	panic("implement me")
}

func (db *memoryDB) SetMonetaryRequestByFullPath(
	ctx context.Context,
	transfer *MonetaryRequest,
	fullPath string,
) (string, error) {
	panic("implement me")
}

func (db *memoryDB) SetMonetaryRequests(
	ctx context.Context,
	userId string,
	transfer []*MonetaryRequest,
	path string,
) error {
	panic("implement me")
}

func (db *memoryDB) SetMonetaryRequestsByFullPath(
	ctx context.Context,
	transfer []*MonetaryRequest,
	fullPath string,
) error {
	panic("implement me")
}

func (db *memoryDB) UpdateMonetaryRequestConfirmed(
	ctx context.Context,
	userId string,
	confirmedFrom bool, //If false ignore
	confirmedTo bool, //If false ignore
	path string,
	snowflake string,
) error {
	panic("implement me")
}

func (db *memoryDB) UpdateMonetaryRequestConfirmedByFullPath(
	ctx context.Context,
	confirmedFrom bool, //If false ignore
	confirmedTo bool, //If false ignore
	fullPath string,
	snowflake string,
) error {
	panic("implement me")
}
