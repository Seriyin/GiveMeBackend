package datastore

import (
	"io"
	"sort"
	"time"
)

type Profile struct {
	Id           int64
	UserId		 string
	DeviceIds	 []string
	PubKey       io.ReadWriter
	SignedPreKey io.ReadWriter
	PreKeyBundle []io.ReadWriter
}

// Files implements sort.Interface, ordering files by timestamp.
// https://golang.org/pkg/sort/#example__sortWrapper
type Files interface {
	sort.Interface
	getTimeCreated(int64) time.Time
	getMetadata(int64) *io.Reader
	getBytes(int64) *io.Reader
}

type memoryFiles struct {
	metadata io.ReadWriter
	bytes    io.ReadWriter
}

type storedFiles struct {
	files map[int64]ObjectData
}

type ProfileDatabase interface {
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

	// Close closes the database, freeing up any available resources.
	Close() error
}
