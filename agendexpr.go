package agendexpr

import "strings"

// ParseSecond parses a seconds field token (e.g. "0", "*/5", "1,15,30")
// into a [Mask] with valid second values (0–59) set.
func ParseSecond(tok string) (Mask[uint64], error) {
	raw, err := parseField[uint64](tok, SecondBound)
	if err != nil {
		return Mask[uint64]{}, err
	}
	return Mask[uint64]{
		raw: raw,
	}, nil
}

// MustParseSecond is like [ParseSecond] but panics on error.
// Intended for use in tests and static expressions.
func MustParseSecond(tok string) Mask[uint64] {
	raw, err := parseField[uint64](tok, SecondBound)
	if err != nil {
		panic(err)
	}
	return Mask[uint64]{
		raw: raw,
	}
}

// ParseMinute parses a minutes field token (e.g. "0", "*/15", "0,30")
// into a [Mask] with valid minute values (0–59) set.
func ParseMinute(tok string) (Mask[uint64], error) {
	raw, err := parseField[uint64](tok, MinuteBound)
	if err != nil {
		return Mask[uint64]{}, err
	}
	return Mask[uint64]{
		raw: raw,
	}, nil
}

// MustParseMinute is like [ParseMinute] but panics on error.
// Intended for use in tests and static expressions.
func MustParseMinute(tok string) Mask[uint64] {
	raw, err := parseField[uint64](tok, MinuteBound)
	if err != nil {
		panic(err)
	}
	return Mask[uint64]{
		raw: raw,
	}
}

// ParseHour parses an hours field token (e.g. "0", "8", "*/2")
// into a [Mask] with valid hour values (0–23) set.
func ParseHour(tok string) (Mask[uint32], error) {
	raw, err := parseField[uint32](tok, HourBound)
	if err != nil {
		return Mask[uint32]{}, err
	}
	return Mask[uint32]{
		raw: raw,
	}, nil
}

// MustParseHour is like [ParseHour] but panics on error.
// Intended for use in tests and static expressions.
func MustParseHour(tok string) Mask[uint32] {
	raw, err := parseField[uint32](tok, HourBound)
	if err != nil {
		panic(err)
	}
	return Mask[uint32]{
		raw: raw,
	}
}

// ParseDayOfMonth parses a day-of-month field token (e.g. "1", "15", "1-10")
// into a [Mask] with valid day-of-month values (1–31) set.
func ParseDayOfMonth(tok string) (Mask[uint32], error) {
	raw, err := parseField[uint32](tok, DayOfMonthBound)
	if err != nil {
		return Mask[uint32]{}, err
	}
	return Mask[uint32]{
		raw: raw,
	}, nil
}

// MustParseDayOfMonth is like [ParseDayOfMonth] but panics on error.
// Intended for use in tests and static expressions.
func MustParseDayOfMonth(tok string) Mask[uint32] {
	raw, err := parseField[uint32](tok, DayOfMonthBound)
	if err != nil {
		panic(err)
	}
	return Mask[uint32]{
		raw: raw,
	}
}

// ParseMonth parses a month field token (e.g. "1", "JUL", "*/3")
// into a [Mask] with valid month values (1–12) set.
// Month names (JAN–DEC) are case-insensitive.
func ParseMonth(tok string) (Mask[uint16], error) {
	raw, err := parseField[uint16](tok, MonthBound)
	if err != nil {
		return Mask[uint16]{}, err
	}
	return Mask[uint16]{
		raw: raw,
	}, nil
}

// MustParseMonth is like [ParseMonth] but panics on error.
// Intended for use in tests and static expressions.
func MustParseMonth(tok string) Mask[uint16] {
	raw, err := parseField[uint16](tok, MonthBound)
	if err != nil {
		panic(err)
	}
	return Mask[uint16]{
		raw: raw,
	}
}

// ParseDayOfWeek parses a day-of-week field token (e.g. "0", "FRI", "1-5")
// into a [Mask] with valid day-of-week values (0–6) set.
// Day names (SUN–SAT) are case-insensitive.
func ParseDayOfWeek(tok string) (Mask[uint8], error) {
	raw, err := parseField[uint8](tok, DayOfWeekBound)
	if err != nil {
		return Mask[uint8]{}, err
	}
	return Mask[uint8]{
		raw: raw,
	}, nil
}

// MustParseDayOfWeek is like [ParseDayOfWeek] but panics on error.
// Intended for use in tests and static expressions.
func MustParseDayOfWeek(tok string) Mask[uint8] {
	raw, err := parseField[uint8](tok, DayOfWeekBound)
	if err != nil {
		panic(err)
	}
	return Mask[uint8]{
		raw: raw,
	}
}

// Parse parses a 6-field cron-like expression string into an [Expression].
//
// The format is: "second minute hour day-of-month month day-of-week"
//
// Each field supports: single values (5), wildcards (*), ranges (1-5),
// steps (*/15, 1-30/5), and comma-separated lists (1,15,30).
// Month and day-of-week names are case-insensitive.
//
// Named aliases (e.g. "@hourly", "@daily") are also supported.
// Returns an error if the expression is malformed or contains out-of-bound values.
//
// Example:
//
//	expr, err := agendexpr.Parse("0 30 8 * * *") // every day at 08:30:00
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

// MustParse is like [Parse] but panics on error.
// Intended for use in tests and static expressions where the input is known to be valid.
//
// Example:
//
//	expr := agendexpr.MustParse("0 30 8 * * *")
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

// Alias resolves a named alias (case-insensitive) to its full 6-field expression string.
// Returns the expression and true if the alias is recognized, or ("", false) otherwise.
//
// Supported aliases: "ticks", "minutely", "5minutes", "10minutes", "30minutes",
// "hourly", "daily", "weekly", "monthly".
//
// Example:
//
//	expr, ok := agendexpr.Alias("hourly") // expr = "0 0 * * * *", ok = true
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
