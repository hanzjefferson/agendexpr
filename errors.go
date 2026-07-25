package agendexpr

import "errors"

var (
	ErrParseDigits = errors.New("unable to parse digits")
	ErrOverflow = errors.New("value out of range for uint8 (overflow)")
	ErrNumberOutOfBound = errors.New("number was out of bound")
	ErrUnknownAlias = errors.New("unknown alias for an expression")
	ErrInvalidFormat = errors.New("invalid format")
	ErrEmpty = errors.New("empty is not allowed")
)