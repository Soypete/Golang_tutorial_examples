package loops

import (
	"fmt"
	"time"
)

type Movie struct {
	Name          string
	Rating        string
	TimeInMinutes int
}

// slices!
// func CaclulateWatchTime(movies []Movie) int {
// var totalTime int
// for _, movie := range movies {
// totalTime += movie.TimeInMinutes
// }
// return totalTime
// }

// c++ solution
func CaclulateWatchTime(movies []Movie) int {
	var totalTime int
	for i := 0; i < len(movies); i++ {
		totalTime += movies[i].TimeInMinutes
	}
	return totalTime
}

func GetDayLeftOfSchool(startMonth, endMonth, endDay int, isFullYear bool) int {
	// the the current year
	year, month, _ := time.Now().UTC().Date()
	endYear := year

	// check if we need to add a year to the end day
	if isFullYear && startMonth >= 7 && month >= 7 {
		endYear++
		fmt.Println(endYear)
	}

	// start counter
	dateTracker := time.Now().UTC()
	var daysLeftOfSchool int

	// this continues to run until the endDay and the endMonth match the days we are adding .to
	for dateTracker.Day() != endDay || dateTracker.Month() != time.Month(endMonth) {

		// add one day to dateTracker
		dateTracker = dateTracker.AddDate(0, 0, 1)

		// add one daty to days left
		daysLeftOfSchool++
	}

	// return days left
	return daysLeftOfSchool
}
