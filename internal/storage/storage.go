package storage

import "errors"

var (
	ErrURLNotFound = errors.New("url not found")
    ErrURlExist = errors.New("url exists")
)
