// Copyright 2015 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package datastore

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
