package wareki

import "testing"

func TestEraNotFound(t *testing.T) {
	t.Parallel()
	if (Era{}) != EraNotFound {
		t.Error("era error type was not match")
	}
	e := new(Era)
	if *e != EraNotFound {
		t.Error("era error type was not match")
	}
	if EraNotFound.String() != "an empty era" {
		t.Error("era not found message was invalid")
	}
	if (Era{}).String() != "an empty era" {
		t.Error("era not found message was invalid")
	}
	if (Era{ShortName: "X", KanjiName: "XX"}).String() == "an empty era" {
		t.Error("era message was invalid")
	}
}
