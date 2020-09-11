package wareki

import (
	"testing"
	"time"
)

func TestEra(t *testing.T) {
	t.Parallel()
	type value struct {
		t Time
		s string
		y int
	}
	values := []value{
		{t: New(time.Date(2019, 10, 2, 0, 0, 0, 0, locJST)), s: "R", y: 1},
		{t: New(time.Date(1000, 1, 1, 0, 0, 0, 0, locJST)), s: "", y: -1}, // may not be parsed
		{t: New(time.Date(1873, 1, 1, 0, 0, 0, 0, locJST)), s: "M", y: 6},
		{t: New(time.Date(1912, 7, 29, 0, 0, 0, 0, locJST)), s: "M", y: 45},
		{t: New(time.Date(1912, 7, 30, 0, 0, 0, 0, locJST)), s: "T", y: 1},
		{t: New(time.Date(1926, 12, 24, 0, 0, 0, 0, locJST)), s: "T", y: 15},
		{t: New(time.Date(1926, 12, 25, 0, 0, 0, 0, locJST)), s: "S", y: 1},
		{t: New(time.Date(1989, 1, 7, 0, 0, 0, 0, locJST)), s: "S", y: 64},
		{t: New(time.Date(1989, 1, 8, 0, 0, 0, 0, locJST)), s: "H", y: 1},
		{t: New(time.Date(2019, 4, 30, 0, 0, 0, 0, locJST)), s: "H", y: 31},
		{t: New(time.Date(2019, 5, 1, 0, 0, 0, 0, locJST)), s: "R", y: 1},
		{t: New(time.Date(2200, 1, 1, 0, 0, 0, 0, locJST)), s: "R", y: 182},
		{t: New(time.Date(2118, 11, 31, 0, 0, 0, 0, locJST)), s: "R", y: 100},
	}
	for _, x := range values {
		e := x.t.Era()
		y := x.t.YearEra()
		if e.ShortName != x.s || y != x.y {
			t.Error(x.t, "must be parsed to", x.s, x.y, ", result:", e.ShortName, y)
		}
	}
}

func TestDate(t *testing.T) {
	t.Parallel()
	type value struct {
		t Time
		y int
	}
	values := []value{
		{t: Date("R", 1, 10, 2, 0, 0, 0, 0, locJST), y: 2019},
		{t: Date("X", 1000, 1, 1, 0, 0, 0, 0, locJST), y: 1000}, // may not be parsed
		{t: Date("M", 6, 1, 1, 0, 0, 0, 0, locJST), y: 1873},
		{t: Date("M", 45, 7, 29, 0, 0, 0, 0, locJST), y: 1912},
		{t: Date("T", 1, 7, 30, 0, 0, 0, 0, locJST), y: 1912},
		{t: Date("T", 15, 12, 24, 0, 0, 0, 0, locJST), y: 1926},
		{t: Date("S", 1, 12, 25, 0, 0, 0, 0, locJST), y: 1926},
		{t: Date("S", 64, 1, 7, 0, 0, 0, 0, locJST), y: 1989},
		{t: Date("H", 1, 1, 8, 0, 0, 0, 0, locJST), y: 1989},
		{t: Date("H", 31, 4, 30, 0, 0, 0, 0, locJST), y: 2019},
		{t: Date("R", 1, 5, 1, 0, 0, 0, 0, locJST), y: 2019},
		{t: Date("R", 182, 1, 1, 0, 0, 0, 0, locJST), y: 2200},
		{t: Date("R", 100, 11, 31, 0, 0, 0, 0, locJST), y: 2118},
		{t: Date("reiwa", 1, 10, 2, 0, 0, 0, 0, locJST), y: 2019},
		{t: Date("令和", 1, 10, 2, 0, 0, 0, 0, locJST), y: 2019},
	}
	for _, x := range values {
		if x.t.Year() != x.y {
			t.Error(x.t, "must be parsed to", x.y, ", result:", x.t.Year())
		}
	}
}

func BenchmarkEra(b *testing.B) {
	t := New(time.Date(1912, 7, 29, 0, 0, 0, 0, locJST))
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			t.Era()
			t.YearEra()
		}
	})
}

func BenchmarkDate1(b *testing.B) {
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Date("showa", 1, 12, 25, 0, 0, 0, 0, locJST)
		}
	})
}
func BenchmarkDate2(b *testing.B) {
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Date("S", 1, 12, 25, 0, 0, 0, 0, locJST)
		}
	})
}
func BenchmarkDate3(b *testing.B) {
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Date("大正", 1, 12, 25, 0, 0, 0, 0, locJST)
		}
	})
}
