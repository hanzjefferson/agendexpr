package agendexpr

import "strings"

func ParseSecond(tok string) (Mask[uint64], error) {
	raw, err := parseField[uint64](tok, SecondBound)
	if err != nil {
		return Mask[uint64]{}, err
	}
	return Mask[uint64]{
		raw: raw,
	}, nil
}

func MustParseSecond(tok string) Mask[uint64] {
	raw, err := parseField[uint64](tok, SecondBound)
	if err != nil {
		panic(err)
	}
	return Mask[uint64]{
		raw: raw,
	}
}

func ParseMinute(tok string) (Mask[uint64], error) {
	raw, err := parseField[uint64](tok, MinuteBound)
	if err != nil {
		return Mask[uint64]{}, err
	}
	return Mask[uint64]{
		raw: raw,
	}, nil
}

func MustParseMinute(tok string) Mask[uint64] {
	raw, err := parseField[uint64](tok, MinuteBound)
	if err != nil {
		panic(err)
	}
	return Mask[uint64]{
		raw: raw,
	}
}

func ParseHour(tok string) (Mask[uint32], error) {
	raw, err := parseField[uint32](tok, HourBound)
	if err != nil {
		return Mask[uint32]{}, err
	}
	return Mask[uint32]{
		raw: raw,
	}, nil
}

func MustParseHour(tok string) Mask[uint32] {
	raw, err := parseField[uint32](tok, HourBound)
	if err != nil {
		panic(err)
	}
	return Mask[uint32]{
		raw: raw,
	}
}

func ParseDayOfMonth(tok string) (Mask[uint32], error) {
	raw, err := parseField[uint32](tok, DayOfMonthBound)
	if err != nil {
		return Mask[uint32]{}, err
	}
	return Mask[uint32]{
		raw: raw,
	}, nil
}

func MustParseDayOfMonth(tok string) Mask[uint32] {
	raw, err := parseField[uint32](tok, DayOfMonthBound)
	if err != nil {
		panic(err)
	}
	return Mask[uint32]{
		raw: raw,
	}
}

func ParseMonth(tok string) (Mask[uint16], error) {
	raw, err := parseField[uint16](tok, MonthBound)
	if err != nil {
		return Mask[uint16]{}, err
	}
	return Mask[uint16]{
		raw: raw,
	}, nil
}

func MustParseMonth(tok string) Mask[uint16] {
	raw, err := parseField[uint16](tok, MonthBound)
	if err != nil {
		panic(err)
	}
	return Mask[uint16]{
		raw: raw,
	}
}

func ParseDayOfWeek(tok string) (Mask[uint8], error) {
	raw, err := parseField[uint8](tok, DayOfWeekBound)
	if err != nil {
		return Mask[uint8]{}, err
	}
	return Mask[uint8]{
		raw: raw,
	}, nil
}

func MustParseDayOfWeek(tok string) Mask[uint8] {
	raw, err := parseField[uint8](tok, DayOfWeekBound)
	if err != nil {
		panic(err)
	}
	return Mask[uint8]{
		raw: raw,
	}
}

func Parse(expr string) (Expression, error) {
	toks, err := normalize(expr)
	if err != nil {
		return Expression{}, err
	}

	seconds, err := ParseSecond(toks[0])
	if err != nil {
		return Expression{}, err
	}

	minutes, err := ParseMinute(toks[1])
	if err != nil {
		return Expression{}, err
	}
	hours, err := ParseHour(toks[2])
	if err != nil {
		return Expression{}, err
	}
	daysOfMonth, err := ParseDayOfMonth(toks[3])
	if err != nil {
		return Expression{}, err
	}
	months, err := ParseMonth(toks[4])
	if err != nil {
		return Expression{}, err
	}
	daysOfWeek, err := ParseDayOfWeek(toks[5])
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

func MustParse(expr string) Expression {
	toks, err := normalize(expr)
	if err != nil {
		panic(err)
	}

	seconds := MustParseSecond(toks[0])
	minutes := MustParseMinute(toks[1])
	hours := MustParseHour(toks[2])
	daysOfMonth := MustParseDayOfMonth(toks[3])
	months := MustParseMonth(toks[4])
	daysOfWeek := MustParseDayOfWeek(toks[5])

	return Expression{
		seconds:    seconds,
		minutes:    minutes,
		hours:      hours,
		dayOfMonth: daysOfMonth,
		months:     months,
		dayOfWeek:  daysOfWeek,
	}
}

func Alias(input string) (string, bool) {
	switch {
	case strings.EqualFold(input, "monthly"):
		return MonthlyAlias, true
	case strings.EqualFold(input, "weekly"):
		return WeeklyAlias, true
	case strings.EqualFold(input, "hourly"):
		return HourlyAlias, true
	case strings.EqualFold(input, "minutely"):
		return MinutelyAlias, true
	case strings.EqualFold(input, "5minutes"):
		return FiveMinutesAlias, true
	case strings.EqualFold(input, "10minutes"):
		return TenMinutesAlias, true
	case strings.EqualFold(input, "30minutes"):
		return HalfMinutesAlias, true
	case strings.EqualFold(input, "ticks"):
		return TicksAlias, true
	default:
		return "", false
	}
}