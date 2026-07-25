package agendexpr

import (
	"time"
)

type Expression struct {
	seconds    Mask[uint64]
	minutes    Mask[uint64]
	hours      Mask[uint32]
	dayOfMonth Mask[uint32]
	months     Mask[uint16]
	dayOfWeek  Mask[uint8]
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
			if next < uint8(mon) {
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
			if next < uint8(hr) {
				t = time.Date(yr, time.Month(mon), dom+1, 0, 0, 0, 0, loc)
			} else {
				t = time.Date(yr, time.Month(mon), dom, int(next), 0, 0, 0, loc)
			}
			continue
		}
 
		if !expr.minutes.Match(uint8(min)) {
			next := expr.minutes.Next(uint8(min))
			if next < uint8(min) {
				t = time.Date(yr, time.Month(mon), dom, hr+1, 0, 0, 0, loc)
			} else {
				t = time.Date(yr, time.Month(mon), dom, hr, int(next), 0, 0, loc)
			}
			continue
		}
 
		if !expr.seconds.Match(uint8(sec)) {
			next := expr.seconds.Next(uint8(sec))
			if next < uint8(sec) {
				t = time.Date(yr, time.Month(mon), dom, hr, min+1, 0, 0, loc)
			} else {
				t = time.Date(yr, time.Month(mon), dom, hr, min, int(next), 0, loc)
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

func (expr Expression) CountdownNow() int64 {
	return expr.Countdown(time.Now())
}