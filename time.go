package wareki

import (
	"time"
)

type Time struct {
	time.Time
}

// New return wareki.Time in JST time zone
func New(t time.Time) Time {
	return Time{t.In(JST)}
}

// Date returns wareki.Time by func that looks like time.Date
//
// Date("R", 1, 10, 2, 0, 0, 0, 0, JST)
// Date("令和", 1, 10, 2, 0, 0, 0, 0, JST)
func Date(era string, year int, month time.Month, day, hour, min, sec, nsec int, loc *time.Location) Time {
	year = getYearInEra(era, year)
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

func getYearInEra(era string, year int) (eraYear int) {
	if e, ok := Eras[eraID(era)]; ok {
		return year + (e.StartedAt.Year() - 1)
	}
	for _, n := range EraOrder {
		e, ok := Eras[n]
		if !ok {
			continue
		}
		if e.ShortName == era || e.KanjiName == era {
			return year + (e.StartedAt.Year() - 1)
		}
	}
	return year
}
