package main

import (
	"fmt"
	"time"

	"github.com/munenari/wareki"
)

func main() {
	loc, _ := time.LoadLocation("Asia/Tokyo")
	t := wareki.New(time.Date(2019, 10, 2, 0, 0, 0, 0, loc))
	e := t.Era()
	if e == wareki.EraNotFound {
		return
	}
	fmt.Printf("%s%02d年%s", e.KanjiName, t.YearEra(), t.Format("01月02日"))
}
