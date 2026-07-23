package agendexpr

import (
	"errors"
	"strings"
)

var (
	bounds = [][]uint8{
		{0, 59},
		{0, 59},
		{0, 23},
		{1, 31},
		{1, 12},
		{0, 6},
	}

	ErrParseDigits = errors.New("unable to parse digits")
	ErrOverflow = errors.New("value out of range for uint8 (overflow)")
	ErrNumberOutOfBounds = errors.New("number was out of bounds")
)

type expressionMask interface {
	~uint8 | ~uint16 | ~uint32 | ~uint64
}

func parseToken[T expressionMask](tok string, bounds []uint8) (T, error) {
	if tok == "*" {
		return fillMask[T](bounds[0], bounds[1]), nil
	}
	var result T

	for part := range strings.SplitSeq(tok, ",") {
		var err error
		var step uint8
		min := bounds[0]
		max := bounds[1]

		if before, after, ok := strings.Cut(part, "/"); ok {
			step, err = parseDigits(after)
			if err != nil {
				return 0, err
			}
			if step < 1 {
				return 0, ErrParseDigits
			}
			part = before
		}

		if before, after, ok := strings.Cut(part, "-"); ok && before != "" {
			min, err = parseDigits(before)
			if err != nil {
				return 0, err
			}
			max, err = parseDigits(after)
			if err != nil {
				return 0, err
			}
		} else if part != "*" {
			num, err := parseDigits(part)
			if err != nil {
				return 0, err
			}
			if num < bounds[0] || num > bounds[1] {
				return 0, ErrNumberOutOfBounds
			}
			if step >= 1 {
				min = num
			} else {
				result |= 1 << num
				continue
			}
		}

		if step >= 1 {
			result |= stepMask[T](min, max, step)
		} else {
			result |= fillMask[T](min, max)
		}
	}
	return result, nil
}

func parseDigits(str string) (uint8, error) {
	if len(str) == 0 {
		return 0, nil
	}

	var result uint8
	for i := 0; i < len(str); i++ {
		if str[i] < '0' || str[i] > '9' {
			return 0, ErrParseDigits
		}

		digit := uint8(str[i] - '0')
		if result > 25 || (result == 25 && digit > 5) {
			return 0, ErrOverflow
		}

		result = (result * 10) + digit
	}
	return result, nil
}

func fillMask[T expressionMask](min, max uint8) T {
	return 1<<(max+1) - 1<<min
}

func stepMask[T expressionMask](min, max, step uint8) T {
	var result T
	for stepIdx := min; stepIdx < max; stepIdx += step {
		result |= 1 << stepIdx
	}
	return result
}


func parseSeconds(tok string) (uint64, error) {
	return parseToken[uint64](tok, bounds[0])
}

func parseMinutes(tok string) (uint64, error) {
	return parseToken[uint64](tok, bounds[1])
}

func parseHours(tok string) (uint32, error) {
	return parseToken[uint32](tok, bounds[2])
}

func parseDaysOfMonth(tok string) (uint32, error) {
	return parseToken[uint32](tok, bounds[3])
}

func parseMonths(tok string) (uint16, error) {
	return parseToken[uint16](tok, bounds[4])
}

func parseDaysOfWeek(tok string) (uint8, error) {
	return parseToken[uint8](tok, bounds[5])
}

func parseYears(tok string) (uint64, error) {
	return parseToken[uint64](tok, bounds[0])
}
