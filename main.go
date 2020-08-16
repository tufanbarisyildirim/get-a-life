package main

import (
	"fmt"
	"time"
)

func main() {

	crs := string2Matrices("   get a life   ")
	commits := crs.Flatten()

	//go to 1 year+ 1 week ago
	yAgo := time.Now().AddDate(-1, 0, -7)

	//seek to first sunday
	for yAgo.Weekday() != time.Sunday {
		yAgo = yAgo.AddDate(0, 0, 1)
	}

	predefined := map[string]int{
		"2020-1-13":  16,
		"2020-4-12":  15,
		"2020-4-11":  18,
		"2020-4-26":  13,
		"2020-4-27":  13,
		"2020-10-1":  12,
		"2020-10-13": 15,
		"2020-10-14": 9,
		"2020-10-15": 9,
		"2020-10-16": 9,
	}

	for i := 0; i < len(commits); i++ {
		commitCount := 1
		if commits[i] == 2 {
			commitCount = 50
		}
		daykey := fmt.Sprintf("%d-%d-%d", yAgo.Year(), yAgo.Month(), yAgo.Day())

		if ad, ok := predefined[daykey]; ok {
			commitCount = commitCount - ad
		}

		fmt.Println( commitCount, " commits for ",dateFormat(yAgo))

		for c := 0; c < commitCount; c++ {
			commit(yAgo)
		}

		yAgo = yAgo.AddDate(0, 0, 1)
	}

}

func newRange(min, max int) []int {
	r := make([]int, max-min+1)
	for i := 0; i < len(r); i++ {
		r[i] = min + i
	}
	return r
}

func dateFormat(tm time.Time) string {
	//Fri Dec 25 19:36:34 +03 2020
	return tm.Format(time.RubyDate)
}
