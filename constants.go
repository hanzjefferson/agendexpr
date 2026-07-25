package agendexpr

const (
	MinSecond uint8 = 0
	MaxSecond uint8 = 59

	MinMinute uint8 = 0
	MaxMinute uint8 = 59

	MinHour uint8 = 0
	MaxHour uint8 = 23

	MinDayOfMonth uint8 = 1
	MaxDayOfMonth uint8 = 31

	MinMonth uint8 = 1
	MaxMonth uint8 = 12

	MinDayOfWeek uint8 = 0
	MaxDayOfWeek uint8 = 6

	MonthlyAlias     = "0 0 0 1 * *"
	WeeklyAlias      = "0 0 0 * * 0"
	HourlyAlias      = "0 0 * * * *"
	MinutelyAlias    = "0 * * * * *"
	FiveMinutesAlias = "0 */5 * * * *"
	TenMinutesAlias  = "0 */10 * * * *"
	HalfMinutesAlias = "0 */30 * * * *"
	TicksAlias       = "* * * * * *"
)
