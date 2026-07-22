package agendexpr

import "errors"

var (
	ErrParseDigits = errors.New("unable to parse digits")
)


func parseDigits(str string) (uint8, error) {
	if len(str) == 0 {
		return 0, nil
	}

	var result uint8
	for i := 0; i < len(str); i++ {
		if str[i] < '0' || str[i] > '9' {
			return 0, ErrParseDigits
		}

		result = (result * 10) + uint8(str[i]-'0')
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
