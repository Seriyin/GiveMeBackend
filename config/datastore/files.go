package datastore

import (
	"io"
	"sort"
	"time"
)

type FileMetadata struct {
}

// Files implements sort.Interface, ordering files by timestamp.
// https://golang.org/pkg/sort/#example__sortWrapper
type Files interface {
	sort.Interface
	getTimeCreated(int64) time.Time
	//getMetadata(int64) *io.Reader
	getMetadata(int64) FileMetadata
	getBytes(int64) *io.Reader
}

type memoryFiles struct {
	metadata io.ReadWriter
	bytes    io.ReadWriter
}
