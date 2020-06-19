package message

import "errors"

var (
	Success  = errors.New("request success")
	NotFound = errors.New("not found")
)
