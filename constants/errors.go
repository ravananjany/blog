package constants

import "errors"

var (
	ErrAlreadyExits    = errors.New("posts already Exists")
	ErrNotFound        = errors.New("post not found")
	ErrPubNotPermitted = errors.New("invalid in input data publication updation not permitted")
)
