package agendexpr

// Valid range boundaries for each cron field.
const (
	// MinSecond is the minimum valid second value (0).
	MinSecond uint8 = 0
	// MaxSecond is the maximum valid second value (59).
	MaxSecond uint8 = 59

	// MinMinute is the minimum valid minute value (0).
	MinMinute uint8 = 0
	// MaxMinute is the maximum valid minute value (59).
	MaxMinute uint8 = 59

	// MinHour is the minimum valid hour value (0).
	MinHour uint8 = 0
	// MaxHour is the maximum valid hour value (23).
	MaxHour uint8 = 23

	// MinDayOfMonth is the minimum valid day-of-month value (1).
	MinDayOfMonth uint8 = 1
	// MaxDayOfMonth is the maximum valid day-of-month value (31).
	MaxDayOfMonth uint8 = 31

	// MinMonth is the minimum valid month value (1).
	MinMonth uint8 = 1
	// MaxMonth is the maximum valid month value (12).
	MaxMonth uint8 = 12

	// MinDayOfWeek is the minimum valid day-of-week value (0 = Sunday).
	MinDayOfWeek uint8 = 0
	// MaxDayOfWeek is the maximum valid day-of-week value (6 = Saturday).
	MaxDayOfWeek uint8 = 6
)

// Named alias expressions. Each maps a human-readable name to its
// full 6-field cron-like expression string.
const (
	// MonthlyAlias is the expression for "@monthly": every first day of the month at midnight.
	MonthlyAlias = "0 0 0 1 * *"
	// WeeklyAlias is the expression for "@weekly": every Sunday at midnight.
	WeeklyAlias = "0 0 0 * * 0"
	// HourlyAlias is the expression for "@hourly": every hour at minute 0, second 0.
	HourlyAlias = "0 0 * * * *"
	// MinutelyAlias is the expression for "@minutely": every minute at second 0.
	MinutelyAlias = "0 * * * * *"
	// FiveMinutesAlias is the expression for "@5minutes": every 5 minutes.
	FiveMinutesAlias = "0 */5 * * * *"
	// TenMinutesAlias is the expression for "@10minutes": every 10 minutes.
	TenMinutesAlias = "0 */10 * * * *"
	// HalfMinutesAlias is the expression for "@30minutes": every 30 minutes.
	HalfMinutesAlias = "0 */30 * * * *"
	// TicksAlias is the expression for "@ticks": every second.
	TicksAlias = "* * * * * *"
)
