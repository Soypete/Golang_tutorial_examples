package loops

type Movie struct {
	Name          string
	Rating        string
	TimeInMinutes int
}

func CaclulateWatchTime(movies []Movie) int {
	var totalTime int
	for _, movie := range movies {
		totalTime += movie.TimeInMinutes
	}
	return totalTime
}
