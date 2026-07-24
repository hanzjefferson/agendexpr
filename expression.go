package agendexpr

type Expression struct {
	seconds    Field[uint64]
	minutes    Field[uint64]
	hours      Field[uint32]
	dayOfMonth Field[uint32]
	months     Field[uint16]
	dayOfWeek  Field[uint8]
}

func NewExpression(expr string) (Expression, error) {
	toks, err := normalize(expr)
	if err != nil {
		return Expression{}, err
	}

	seconds, err := ParseField[uint64](toks[0], bounds[0])
	if err != nil {
		return Expression{}, err
	}

	minutes, err := ParseField[uint64](toks[1], bounds[1])
	if err != nil {
		return Expression{}, err
	}
	hours, err := ParseField[uint32](toks[2], bounds[2])
	if err != nil {
		return Expression{}, err
	}
	daysOfMonth, err := ParseField[uint32](toks[3], bounds[3])
	if err != nil {
		return Expression{}, err
	}
	months, err := ParseField[uint16](toks[4], bounds[4])
	if err != nil {
		return Expression{}, err
	}
	daysOfWeek, err := ParseField[uint8](toks[5], bounds[5])
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
