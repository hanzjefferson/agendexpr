package agendexpr

import "errors"

var (
	ErrParseDigits = errors.New("unable to parse digits")
)

type digit interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

func parseDigits[T digit](str string) (T, error) {
	if len(str) == 0 {
		return 0, nil
	}

	var result T
	for i := 0; i < len(str); i++ {
		if str[i] < '0' || str[i] > '9' {
			return 0, ErrParseDigits
		}

		result = (result * 10) + T(str[i]-'0')
	}
	return result, nil
}

func fillMask[T expressionMask](min, max uint8) T {
	return 1<<(max+1) - 1<<min
}

func stepMask[T expressionMask](max, step uint8) T {
	var result T
	for stepIdx := step; stepIdx < max; stepIdx += step {
		result |= 1 << stepIdx
	}
	return result
}
