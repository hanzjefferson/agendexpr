package agendexpr

// Field boundary pairs [min, max] for each cron field.
// Used internally by the parser to validate values against their allowed ranges.
var (
	// SecondBound is the valid range for seconds: [0, 59].
	SecondBound = [2]uint8{MinSecond, MaxSecond}
	// MinuteBound is the valid range for minutes: [0, 59].
	MinuteBound = [2]uint8{MinMinute, MaxMinute}
	// HourBound is the valid range for hours: [0, 23].
	HourBound = [2]uint8{MinHour, MaxHour}
	// DayOfMonthBound is the valid range for day-of-month: [1, 31].
	DayOfMonthBound = [2]uint8{MinDayOfMonth, MaxDayOfMonth}
	// MonthBound is the valid range for months: [1, 12].
	MonthBound = [2]uint8{MinMonth, MaxMonth}
	// DayOfWeekBound is the valid range for day-of-week: [0, 6].
	DayOfWeekBound = [2]uint8{MinDayOfWeek, MaxDayOfWeek}
)
