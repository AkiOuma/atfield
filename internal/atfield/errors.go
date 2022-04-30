package atfield

import "errors"

var (
	errInvalidPath = errors.New("error: invalid path")
	errReadPackage = errors.New("error: failed to read package")
)
