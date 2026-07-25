package agendexpr

import (
	"strings"
)

func parseField[T MaskType](tok string, bound [2]uint8) (T, error) {
	tok = strings.TrimSpace(tok)
	if tok == "" {
		return 0, ErrEmpty
	}
	if tok == "*" {
		return fillMask[T](bound[0], bound[1]), nil
	}

	var result T

	for part := range strings.SplitSeq(tok, ",") {
		var (
			err            error
			step, min, max uint8
		)

		step, raw, err := parseStep(part)
		if err != nil {
			return 0, err
		}
		if step > 0 {
			part = raw
		}

		min, max, err = parseRange(part)
		if err != nil {
			return 0, err
		}
		if min == 0 && max == 0 {
			if part != "*" {
				num, err := digits(part)
				if err != nil {
					return 0, err
				}
				if num < bound[0] || num > bound[1] {
					return 0, ErrNumberOutOfBound
				}
				if step > 0 {
					min = num
					max = bound[1]
				} else {
					result |= T(1) << num
					continue
				}
			} else {
				min = bound[0]
				max = bound[1]
			}
		} else {
			if min < bound[0] || max > bound[1] {
				return 0, ErrNumberOutOfBound
			}
		}

		if step >= 1 {
			result |= stepMask[T](min, max, step)
		} else {
			result |= fillMask[T](min, max)
		}
	}

	if result == 0 {
		return 0, nil
	}

	return result, nil
}

func parseStep(tok string) (uint8, string, error) {
	if before, after, ok := strings.Cut(tok, "/"); ok {
		step, err := digits(after)
		if err != nil {
			return 0, "", err
		}
		if step < 1 {
			return 0, "", ErrParseDigits
		}

		return step, before, nil
	}
	return 0, "", nil
}

func parseRange(tok string) (uint8, uint8, error) {
	if before, after, ok := strings.Cut(tok, "-"); ok && before != "" {
		min, err := digits(before)
		if err != nil {
			return 0, 0, err
		}
		max, err := digits(after)
		if err != nil {
			return 0, 0, err
		}

		return min, max, nil
	}

	return 0, 0, nil
}

func digits(str string) (uint8, error) {
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
		if result > 25 || (result == 25 && digit > 5) {
			return 0, ErrOverflow
		}
		result = result*10 + digit
	}
	return result, nil
}

func fillMask[T MaskType](min, max uint8) T {
	return T(1)<<(max+1) - T(1)<<min
}

func stepMask[T MaskType](min, max, step uint8) T {
	var result T
	for i := min; i < max; i += step {
		result |= T(1) << i
	}
	return result
}
