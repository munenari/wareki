package wareki

import (
	"time"
)

// Time struct
type Time struct {
	time.Time
}

// New return wareki.Time in JST time zone
func New(t time.Time) Time {
	return Time{t.In(locJST)}
}

// Date returns wareki.Time by func that looks like time.Date
func Date(era string, year int, month time.Month, day, hour, min, sec, nsec int, loc *time.Location) Time {
	year = func(y int) int {
		if e, ok := Eras[eraID(era)]; ok {
			return y + (e.StartedAt.Year() - 1)
		}
		for _, n := range EraOrder {
			e, ok := Eras[n]
			if !ok {
				continue
			}
			if e.ShortName == era || e.KanjiName == era {
				return y + (e.StartedAt.Year() - 1)
			}
		}
		return y
	}(year)
	return New(time.Date(year, month, day, hour, min, sec, nsec, loc))
}

// Era returns a year era name what belongs to
func (t Time) Era() Era {
	for _, n := range EraOrder {
		e, ok := Eras[n]
		if !ok {
			continue
		}
		if !t.Before(e.StartedAt) { // !Before: t >= e.StartedAt, After: t > e.StartedAt
			return e
		}
	}
	return EraNotFound
}

// YearEra returns year number in era
// this will return -1 when era was not found
func (t Time) YearEra() int {
	e := t.Era()
	if e == (EraNotFound) {
		return -1
	}
	return t.Year() - (e.StartedAt.Year() - 1)
}
