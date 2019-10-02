package wareki

import (
	"fmt"
	"time"
)

type (
	eraID string
	// Era struct
	Era struct {
		ShortName string
		KanjiName string
		StartedAt time.Time
	}
)

var (
	locJST, _ = time.LoadLocation("Asia/Tokyo")
	// Eras means list of era, this will be used by searching
	// refer from: https://ja.wikipedia.org/wiki/%E5%85%83%E5%8F%B7%E4%B8%80%E8%A6%A7_(%E6%97%A5%E6%9C%AC)
	// ! we can override this for other GENGOs
	Eras = map[eraID]Era{
		"meiji":  Era{ShortName: "M", KanjiName: "明治", StartedAt: time.Date(1868, 10, 23, 0, 0, 0, 0, locJST)}, //明治: 1868-10-23
		"taisho": Era{ShortName: "T", KanjiName: "大正", StartedAt: time.Date(1912, 7, 30, 0, 0, 0, 0, locJST)},  //大正: 1912-07-30
		"showa":  Era{ShortName: "S", KanjiName: "昭和", StartedAt: time.Date(1926, 12, 25, 0, 0, 0, 0, locJST)}, //昭和: 1926-12-25
		"heisei": Era{ShortName: "H", KanjiName: "平成", StartedAt: time.Date(1989, 1, 8, 0, 0, 0, 0, locJST)},   //平成: 1989-01-08
		"reiwa":  Era{ShortName: "R", KanjiName: "令和", StartedAt: time.Date(2019, 5, 1, 0, 0, 0, 0, locJST)},   //令和: 2019-05-01
	}
	// EraOrder must be correct, this will be used as order when searching
	// ! we can override this for other GENGOs
	EraOrder = []eraID{"reiwa", "heisei", "showa", "taisho", "meiji"}
	// EraNotFound using to era error type
	EraNotFound Era
)

func (e Era) String() string {
	if e == EraNotFound {
		return fmt.Sprintf("an empty era")
	}
	return fmt.Sprintf("%s: %s", e.KanjiName, e.StartedAt.Format("2006-01-02"))
}
