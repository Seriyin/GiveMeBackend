package datastore

import (
	"io"
)

type Profile struct {
	UserId   UID
	Metadata Metadata
}

type UID struct {
	Id     string
	Device io.ReadWriter
	PubKey io.ReadWriter
	Token  io.ReadWriter
	//	SignedPreKey io.ReadWriter
	//	PreKeyBundle []io.ReadWriter
}

type Metadata struct {
	PaymentProviders []PaymentProvider
	NumberPayments   uint64
}


type ProfileDatabase interface {
	IsBlocked()	(bool, error)

	// ListProfiles returns a list of profiles.
	ListProfiles() ([]*Profile, error)

	// ListFilesSharedBy returns a list of files, ordered by timestamp,
	// filtered by the user who created the files.
	ListFilesSharedBy(userID string) (*Files, error)

	// GetProfile retrieves a profile by its ID.
	GetProfile(id int64) (*Profile, error)

	// AddProfile saves a given profile, assigning it a new ID.
	AddProfile(p *Profile) (id int64, err error)

	// DeleteProfile removes a given profile by its ID.
	DeleteProfile(id int64) error

	// UpdateProfile updates the entry for a given profile.
	UpdateProfile(p *Profile) error

	// RegenProfile discards and reinitializes known profile
	RegenProfile(p *Profile) error

	// Close closes the database, freeing up any available resources.
	Close() error
}
