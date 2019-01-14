package common

import (
	"errors"
)

type AESMode int

const (
	_ = iota
	CBC
	ECB
)

var (
	BadDataErr = errors.New("BAD_DATA")
)
