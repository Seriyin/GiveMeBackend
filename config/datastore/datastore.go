package datastore

import "time"

// Interface for defining operations over underlying DB
// needed for the entire service.
type GiveMeDatabase interface {
	//Profile methods

	//Check if a user is blocked by a given user with userId.
	IsBlocked(
		userId string,
		blocked string,
	) (bool, error)

	// ListFilesSharedBy returns a list of files, ordered by timestamp,
	// filtered by the user who created the files.
	// ListFilesSharedBy(userId string) (*Files, error)

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

	// Monetary Transfer methods

	// GetMonetaryTransfer methods fetch appropriate instances of
	// MonetaryTransfer from db.

	// GetMonetaryTransfer by unique id (snowflake).
	GetMonetaryTransfer(
		userId string,
		snowflake string,
	) (*MonetaryTransfer, error)

	// GetMonetaryTransfers by current date until dateBefore.
	GetMonetaryTransfersDate(
		userId string,
		dateBefore time.Time,
	) ([]*MonetaryTransfer, error)

	// GetMonetaryTransfers between two dates.
	GetMonetaryTransfersInterval(
		userId string,
		dateAfter time.Time,
		dateBefore time.Time,
	) ([]*MonetaryTransfer, error)

	// GetMonetaryTransfers for a group transfer.
	GetMonetaryTransfersFromGroup(
		userId string,
		groupId int64,
	) ([]*MonetaryTransfer, error)

	// GetMonetaryTransfers for a recurrent transfer.
	GetMonetaryTransfersRecurrent(
		userId string,
		recurrentId int64,
	) ([]*MonetaryTransfer, error)

	SetMonetaryTransfer(
		userId string,
		transfer *MonetaryTransfer,
		path string,
	) error
}
