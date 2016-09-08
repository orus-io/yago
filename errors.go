package yago

import (
	"errors"
)

var (
	// ErrRecordNotFound is returned by Update, Delete, or One() if
	// not matching record were found
	ErrRecordNotFound = errors.New("yago.RecordNotFound")

	// ErrMultipleRecords is returned by Update or One() if too many
	// records matched the statement
	ErrMultipleRecords = errors.New("yago.MultipleRecords")
)
