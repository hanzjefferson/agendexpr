package agendexpr

type Expression struct {
	seconds    Token[uint64]
	minutes    Token[uint64]
	hours      Token[uint32]
	dayOfMonth Token[uint32]
	months     Token[uint16]
	dayOfWeek  Token[uint8]
}

func NewExpression(expr string) (Expression, error) {
	toks, err := normalize(expr)
	if err != nil {
		return Expression{}, err
	}

	seconds, err := ParseToken[uint64](toks[0], bounds[0])
	if err != nil {
		return Expression{}, err
	}

	minutes, err := ParseToken[uint64](toks[1], bounds[1])
	if err != nil {
		return Expression{}, err
	}
	hours, err := ParseToken[uint32](toks[2], bounds[2])
	if err != nil {
		return Expression{}, err
	}
	daysOfMonth, err := ParseToken[uint32](toks[3], bounds[3])
	if err != nil {
		return Expression{}, err
	}
	months, err := ParseToken[uint16](toks[4], bounds[4])
	if err != nil {
		return Expression{}, err
	}
	daysOfWeek, err := ParseToken[uint8](toks[5], bounds[5])
	if err != nil {
		return Expression{}, err
	}

	return Expression{
		seconds:    seconds,
		minutes:    minutes,
		hours:      hours,
		dayOfMonth: daysOfMonth,
		months:     months,
		dayOfWeek:  daysOfWeek,
	}, nil
}
