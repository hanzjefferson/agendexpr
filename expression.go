package agendexpr

import (
	"time"
)

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

func (expr Expression) Next(from time.Time) time.Time {
	loc := from.Location()
 
	t := from.Truncate(time.Second).Add(time.Second)
 
	limit := t.AddDate(4, 0, 0)
 
	for t.Before(limit) {
		yr  := t.Year()
		mon := int(t.Month())
		dom := t.Day()
		hr  := t.Hour()
		min := t.Minute()
		sec := t.Second()
 
		if !expr.months.Match(uint8(mon)) {
			next := expr.months.Next(uint8(mon))
			if next == -1 {
				return time.Time{}
			}
			if next < mon {
				yr++
			}
			t = time.Date(yr, time.Month(next), 1, 0, 0, 0, 0, loc)
			continue
		}
 
		if !expr.dayOfMonth.Match(uint8(dom)) || !expr.dayOfWeek.Match(uint8(t.Weekday())) {
			t = time.Date(yr, time.Month(mon), dom+1, 0, 0, 0, 0, loc)
			continue
		}
 
		if !expr.hours.Match(uint8(hr)) {
			next := expr.hours.Next(uint8(hr))
			if next == -1 {
				return time.Time{}
			}
			if next < hr {
				t = time.Date(yr, time.Month(mon), dom+1, 0, 0, 0, 0, loc)
			} else {
				t = time.Date(yr, time.Month(mon), dom, next, 0, 0, 0, loc)
			}
			continue
		}
 
		if !expr.minutes.Match(uint8(min)) {
			next := expr.minutes.Next(uint8(min))
			if next == -1 {
				return time.Time{}
			}
			if next < min {
				t = time.Date(yr, time.Month(mon), dom, hr+1, 0, 0, 0, loc)
			} else {
				t = time.Date(yr, time.Month(mon), dom, hr, next, 0, 0, loc)
			}
			continue
		}
 
		if !expr.seconds.Match(uint8(sec)) {
			next := expr.seconds.Next(uint8(sec))
			if next == -1 {
				return time.Time{}
			}
			if next < sec {
				t = time.Date(yr, time.Month(mon), dom, hr, min+1, 0, 0, loc)
			} else {
				t = time.Date(yr, time.Month(mon), dom, hr, min, next, 0, loc)
			}
			continue
		}
 
		return t
	}
 
	return time.Time{}
}
 
func (expr Expression) DueMonth(of time.Time) bool {
	return expr.months.Match(uint8(of.Month()))
}
 
func (expr Expression) DueDay(of time.Time) bool {
	return expr.DueMonth(of) &&
		expr.dayOfMonth.Match(uint8(of.Day())) &&
		expr.dayOfWeek.Match(uint8(of.Weekday()))
}
 
func (expr Expression) DueHour(of time.Time) bool {
	return expr.DueDay(of) && expr.hours.Match(uint8(of.Hour()))
}
 
func (expr Expression) DueMinute(of time.Time) bool {
	return expr.DueHour(of) && expr.minutes.Match(uint8(of.Minute()))
}
 
func (expr Expression) Due(of time.Time) bool {
	return expr.DueMinute(of) && expr.seconds.Match(uint8(of.Second()))
}
 
func (expr Expression) Countdown(from time.Time) int64 {
	next := expr.Next(from)
	if next.IsZero() {
		return -1
	}
	return next.Unix() - from.Unix()
}