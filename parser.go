package agendexpr

import (
	"errors"
	"math/bits"
	"strings"
)

// bounds defines [min, max] inclusive for each cron field:
// seconds, minutes, hours, day-of-month, month, day-of-week.
var bounds = [][2]uint8{
	{0, 59}, // seconds
	{0, 59}, // minutes
	{0, 23}, // hours
	{1, 31}, // day of month
	{1, 12}, // month
	{0, 6},  // day of week
}

var (
	ErrParseDigits       = errors.New("unable to parse digits")
	ErrOverflow          = errors.New("value out of range for uint8 (overflow)")
	ErrNumberOutOfBounds = errors.New("number was out of bounds")
)

// parseField parses a single cron field field (e.g. "*/5", "1-5", "0,30")
// into a bitmask with the corresponding positions set, and returns the
// effective [min, max] of the set bits. Zero heap allocations.
func parseField[T fieldMask](tok string, fieldBounds [2]uint8) (T, [2]uint8, error) {
	if tok == "*" {
		return fillMask[T](fieldBounds[0], fieldBounds[1]), fieldBounds, nil
	}

	var result T

	for part := range strings.SplitSeq(tok, ",") {
		var (
			err  error
			step uint8
			min  = fieldBounds[0]
			max  = fieldBounds[1]
		)

		// Handle optional step notation: "expr/step"
		if before, after, ok := strings.Cut(part, "/"); ok {
			if step, err = parseDigits(after); err != nil {
				return 0, [2]uint8{}, err
			}
			if step < 1 {
				return 0, [2]uint8{}, ErrParseDigits
			}
			part = before
		}

		// Handle range "min-max", single value, or wildcard "*"
		if before, after, ok := strings.Cut(part, "-"); ok && before != "" {
			if min, err = parseDigits(before); err != nil {
				return 0, [2]uint8{}, err
			}
			if max, err = parseDigits(after); err != nil {
				return 0, [2]uint8{}, err
			}
		} else if part != "*" {
			// Single numeric value
			num, err := parseDigits(part)
			if err != nil {
				return 0, [2]uint8{}, err
			}
			if num < fieldBounds[0] || num > fieldBounds[1] {
				return 0, [2]uint8{}, ErrNumberOutOfBounds
			}
			if step >= 1 {
				min = num // use as base for step range
			} else {
				result |= T(1) << num
				continue
			}
		}

		if step >= 1 {
			result |= stepMask[T](min, max, step)
		} else {
			result |= fillMask[T](min, max)
		}
	}

	if result == 0 {
		return 0, [2]uint8{}, nil
	}

	// Derive effective [min, max] from the set bits using hardware intrinsics.
	m := uint64(result)
	return result, [2]uint8{
		uint8(bits.TrailingZeros64(m)), // index of lowest set bit  → effective min
		uint8(bits.Len64(m) - 1),      // index of highest set bit → effective max
	}, nil
}

// parseDigits converts a decimal string to uint8 without any heap allocations.
// An empty string returns (0, nil).
func parseDigits(str string) (uint8, error) {
	if len(str) == 0 {
		return 0, nil
	}

	var result uint8
	for i := 0; i < len(str); i++ {
		b := str[i]
		if b < '0' || b > '9' {
			return 0, ErrParseDigits
		}
		digit := b - '0'
		// Overflow guard: uint8 max is 255 = 25*10+5; anything beyond is an error.
		if result > 25 || (result == 25 && digit > 5) {
			return 0, ErrOverflow
		}
		result = result*10 + digit
	}
	return result, nil
}

// fillMask returns a bitmask with all bits set from min to max (inclusive).
// Using T(1) instead of an untyped 1 ensures shifts stay within the type's
// domain: e.g. uint64(1)<<64 wraps to 0, making the subtraction produce the
// correct all-bits-set mask for a full-width T.
func fillMask[T fieldMask](min, max uint8) T {
	return T(1)<<(max+1) - T(1)<<min
}

// stepMask returns a bitmask with bits set at every `step` positions
// starting from min, up to (but not including) max.
// NOTE: the upper bound is exclusive; this is intentional and matches
// the existing package semantics (a half-open [min, max) range with step).
func stepMask[T fieldMask](min, max, step uint8) T {
	var result T
	for i := min; i < max; i += step {
		result |= T(1) << i
	}
	return result
}