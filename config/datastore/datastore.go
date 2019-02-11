package datastore

import (
	"context"
	"time"
)

// Interface for defining operations over underlying DB
// needed for the entire service.
type GiveMeDatabase interface {
	//Profile methods

	//Check if a user is blocked by a given user with userId.
	IsBlocked(
		ctx context.Context,
		userId string,
		blocked string,
	) (bool, error)

	// ListFilesSharedBy returns a list of files, ordered by timestamp,
	// filtered by the user who created the files.
	// ListFilesSharedBy(userId string) (*Files, error)

	// GetProfile retrieves a profile by its ID.
	GetProfile(
		ctx context.Context,
		userId string,
	) (*Profile, error)

	// GetProfileByPhoneNumber retrieves a profile and Id by its associated phone number.
	GetProfileByPhoneNumber(
		ctx context.Context,
		phoneNumber string,
	) (*Profile, error)

	// GetProfileIdByPhoneNumber retrieves a profile's Id by its associated phone number.
	GetProfileIdByPhoneNumber(
		ctx context.Context,
		phoneNumber string,
	) (string, error)

	// AddProfile saves a given profile, assigning it a new ID.
	AddProfile(
		ctx context.Context,
		p *Profile,
	) (userId string, err error)

	// DeleteProfile removes a given profile by its ID.
	DeleteProfile(
		ctx context.Context,
		userId string,
	) error

	// UpdateProfile updates the entry for a given profile.
	UpdateProfile(
		ctx context.Context,
		p *Profile,
	) error

	// RegenProfile discards and reinitializes known profile
	RegenProfile(
		ctx context.Context,
		p *Profile,
	) error

	// Close closes the database, freeing up any available resources.
	Close() error

	// Monetary Transfer methods

	AddMonetaryTransfer(
		ctx context.Context,
		userId string,
		transfer *MonetaryTransfer,
		path string,
	) (string, error)

	AddMonetaryTransferByFullPath(
		ctx context.Context,
		transfer *MonetaryTransfer,
		fullPath string,
	) (string, error)

	// GetMonetaryTransfer methods fetch appropriate instances of
	// MonetaryTransfer from db.

	// GetMonetaryTransfer by unique id (snowflake).
	GetMonetaryTransferWithDate(
		ctx context.Context,
		userId string,
		date time.Time,
		snowflake string,
	) (*MonetaryTransfer, error)

	// GetMonetaryTransfer by unique id (snowflake).
	GetMonetaryTransferWithDateString(
		ctx context.Context,
		userId string,
		date string,
		snowflake string,
	) (*MonetaryTransfer, error)

	// GetMonetaryTransfers by current date until dateBefore.
	GetMonetaryTransfersDate(
		ctx context.Context,
		userId string,
		dateBefore time.Time,
	) ([]*MonetaryTransfer, error)

	// GetMonetaryTransfers between two dates.
	GetMonetaryTransfersInterval(
		ctx context.Context,
		userId string,
		dateAfter time.Time,
		dateBefore time.Time,
	) ([]*MonetaryTransfer, error)

	// GetMonetaryTransfers for a group transfer.
	GetMonetaryTransfersFromGroup(
		ctx context.Context,
		userId string,
		date time.Time,
		groupId int64,
	) ([]*MonetaryTransfer, error)

	// GetMonetaryTransfers for a recurrent transfer.
	GetMonetaryTransfersRecurrent(
		ctx context.Context,
		userId string,
		recurrentId int64,
	) ([]*MonetaryTransfer, error)

	SetMonetaryTransfer(
		ctx context.Context,
		userId string,
		transfer *MonetaryTransfer,
		path string,
	) (string, error)

	SetMonetaryTransferByFullPath(
		ctx context.Context,
		transfer *MonetaryTransfer,
		fullPath string,
	) (string, error)

	SetMonetaryTransfers(
		ctx context.Context,
		userId string,
		transfer []*MonetaryTransfer,
		path string,
	) error

	SetMonetaryTransfersByFullPath(
		ctx context.Context,
		transfer []*MonetaryTransfer,
		fullPath string,
	) error

	UpdateMonetaryTransferConfirmed(
		ctx context.Context,
		userId string,
		confirmedFrom bool, //If false ignore
		confirmedTo bool, //If false ignore
		path string,
		snowflake string,
	) error

	UpdateMonetaryTransferConfirmedByFullPath(
		ctx context.Context,
		confirmedFrom bool, //If false ignore
		confirmedTo bool, //If false ignore
		fullPath string,
		snowflake string,
	) error
}
