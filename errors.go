package simple_substitution_cipher

import "errors"

var (
	// ErrKeyUnavailable key不可用
	ErrKeyUnavailable = errors.New("the key must be 26 characters and cover exactly A-Z")
)
