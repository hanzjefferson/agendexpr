package agendexpr

import "errors"

var (
	// ErrParseDigits is returned when a token contains non-digit characters
	// where a number was expected.
	ErrParseDigits = errors.New("unable to parse digits")

	// ErrOverflow is returned when a parsed number exceeds the uint8 range (0–255).
	ErrOverflow = errors.New("value out of range for uint8 (overflow)")

	// ErrNumberOutOfBound is returned when a parsed number falls outside
	// the valid range for its field (e.g. second > 59, month > 12).
	ErrNumberOutOfBound = errors.New("number was out of bound")

	// ErrUnknownAlias is returned when an "@" prefixed alias name is not recognized.
	ErrUnknownAlias = errors.New("unknown alias for an expression")

	// ErrInvalidFormat is returned when the expression contains more than 6 fields.
	ErrInvalidFormat = errors.New("invalid format")

	// ErrEmpty is returned when an empty expression or field token is provided.
	ErrEmpty = errors.New("empty is not allowed")
)
