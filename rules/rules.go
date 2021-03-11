package rules

import (
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// str rules
var (
	// error thrown when string is not spaceless
	ErrIsSpaceless = validation.NewError("validation_is_spaceless", "must not contain spaces")
	// error thrown when string is not numeric
	ErrIsNumeric = validation.NewError("validation_is_numeric", "must be numeric")
	// error thrown when string is not int
	ErrIsInt = validation.NewError("validation_is_int", "must be int")
	// error thrown when date is not a valid iso8601 date
	ErrIsIso8601Date = validation.NewError("validation_is_iso8601", "date is not valid iso 8601 date")
	// error thrown when date is not a millisecond date
	ErrIsMillisecondDate = validation.NewError("validation_is_millisecond_date", "date is not millisecond format")
)

// path rules
var (
	// error thrown when path is invalid (or not exists)
	ErrIsPath = validation.NewError("validation_is_invalid_path", "path is invalid")
	// error thrown when path is not a dir (or not exists)
	ErrIsDir = validation.NewError("validation_is_invalid_dir", "path either does not exist or is not a directory")
	// error thrown when path is not a file (or not exists)
	ErrIsFile = validation.NewError("validation_is_invalid_file", "path either does not exist or is not a file")
)

// net rules
var (
	// error when port is invalid
	ErrIsInvalidPort = validation.NewError("validation_is_valid_port", fmt.Sprintf("port is invalid, expected port to be greater than %d and less than %d as per RFC-6056 (https://tools.ietf.org/html/rfc6056)", MinPort-1, MaxPort+1))
	// error when port is unavailable
	ErrPortUnavailable = validation.NewError("validation_is_port_unavailable", "port is unavailable")
)
