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

	// Monetary Request methods

	AddMonetaryRequest(
		ctx context.Context,
		userId string,
		transfer *MonetaryRequest,
		path string,
	) (string, error)

	AddMonetaryRequestByFullPath(
		ctx context.Context,
		transfer *MonetaryRequest,
		fullPath string,
	) (string, error)

	// GetMonetaryRequest methods fetch appropriate instances of
	// MonetaryRequest from db.

	// GetMonetaryRequest by unique id (snowflake).
	GetMonetaryRequestWithDate(
		ctx context.Context,
		userId string,
		date time.Time,
		snowflake string,
	) (*MonetaryRequest, error)

	// GetMonetaryRequest by unique id (snowflake).
	GetMonetaryRequestWithDateString(
		ctx context.Context,
		userId string,
		date string,
		snowflake string,
	) (*MonetaryRequest, error)

	// GetMonetaryRequests by current date until dateBefore.
	GetMonetaryRequestsDate(
		ctx context.Context,
		userId string,
		dateBefore time.Time,
	) ([]*MonetaryRequest, error)

	// GetMonetaryRequests between two dates.
	GetMonetaryRequestsInterval(
		ctx context.Context,
		userId string,
		dateAfter time.Time,
		dateBefore time.Time,
	) ([]*MonetaryRequest, error)

	// GetMonetaryRequests for a group transfer.
	GetMonetaryRequestsFromGroup(
		ctx context.Context,
		userId string,
		date time.Time,
		groupId int64,
	) ([]*MonetaryRequest, error)

	// GetMonetaryRequests for a recurrent transfer.
	GetMonetaryRequestsRecurrent(
		ctx context.Context,
		userId string,
		recurrentId int64,
	) ([]*MonetaryRequest, error)

	SetMonetaryRequest(
		ctx context.Context,
		userId string,
		transfer *MonetaryRequest,
		path string,
	) (string, error)

	SetMonetaryRequestByFullPath(
		ctx context.Context,
		transfer *MonetaryRequest,
		fullPath string,
	) (string, error)

	SetMonetaryRequests(
		ctx context.Context,
		userId string,
		transfer []*MonetaryRequest,
		path string,
	) error

	SetMonetaryRequestsByFullPath(
		ctx context.Context,
		transfer []*MonetaryRequest,
		fullPath string,
	) error

	UpdateMonetaryRequestConfirmed(
		ctx context.Context,
		userId string,
		confirmedFrom bool, //If false ignore
		confirmedTo bool, //If false ignore
		path string,
		snowflake string,
	) error

	UpdateMonetaryRequestConfirmedByFullPath(
		ctx context.Context,
		confirmedFrom bool, //If false ignore
		confirmedTo bool, //If false ignore
		fullPath string,
		snowflake string,
	) error
}
