package projectError

import "errors"

var (
	ErrDataNotFound = errors.New("data not found")

	ErrNoUpdatedData = errors.New("no data to update")

	ErrConflictingData = errors.New("data conflicts with existing data in unique column")

	ErrTokenCreation = errors.New("error creating token")

	ErrInvalidCredentials = errors.New("invalid username or password")
)
