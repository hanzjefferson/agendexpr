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

	ErrAbnormalTokens    = errors.New("abnormal tokens")
	ErrUnknownSyntax     = errors.New("unknown expression syntax")
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
		var partResult, filter T

		if before, after, ok := strings.Cut(part, "/"); ok {
			step, err := parseDigits(after)
			if err != nil {
				return 0, err
			}

			filter |= stepMask[T](bounds[1], step)
			part = before
		}

		if before, after, ok := strings.Cut(part, "-"); ok && before != "" {
			min, err := parseDigits(before)
			if err != nil {
				return 0, err
			}
			max, err := parseDigits(after)
			if err != nil {
				return 0, err
			}
			partResult |= fillMask[T](min, max)
		} else if part == "*" {
			partResult |= fillMask[T](bounds[0], bounds[1])
		} else {
			num, err := parseDigits(part)
			if err != nil {
				return 0, err
			}
			if filter != 0 {
				partResult |= 1<<(bounds[1]+1) - 1<<num
			} else {
				partResult |= 1 << num
			}
		}

		if filter != 0 {
			partResult &= filter
		}
		result |= partResult
	}
	return result, nil
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
