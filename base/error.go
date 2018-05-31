package base

import "errors"

var (
	ErrParamType = errors.New("args error")

	ErrorPathNotExist = errors.New("no such file or directory")
)
