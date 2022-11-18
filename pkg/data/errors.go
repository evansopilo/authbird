package data

import "errors"

var (
	ErrNoDocument     = errors.New("error no document")
	ErrCreateDocument = errors.New("error create document")
	ErrUpdateDocument = errors.New("error update document")
	ErrDeleteDocument = errors.New("error delete document")
)
