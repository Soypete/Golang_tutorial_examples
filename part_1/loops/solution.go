package loops

import "time"

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
	year, _, _ := time.Now().Date()
	endYear := year

	// check if we need to add a year to the end day
	if isFullYear && startMonth >= 7 {
		endYear++
		//TODO: println
	}

	// define last day of school
	endTime := time.Date(endYear, time.Month(endMonth), endDay, 0, 0, 0, 0, time.UTC)

	// start counter
	dateTracker := time.Now()
	var daysLeftOfSchool int

	for dateTracker.Unix() <= endTime.Unix() {
		// add one day to dateTracker
		dateTracker.AddDate(0, 0, 1)

		// add one daty to days left
		daysLeftOfSchool++
		//TODO: println
	}

	// return days left
	return daysLeftOfSchool
}
