package loops

import (
	"fmt"
	"testing"
	"time"
)

var HarryPotter = []Movie{
	{
		Name:          "Harry Potter and the Sorcerer's Stone",
		Rating:        "PG",
		TimeInMinutes: 152,
	},
	{
		Name:          "Harry Potter and the Chamber of Secrets",
		Rating:        "PG",
		TimeInMinutes: 161,
	},
	{
		Name:          "Harry Potter and the Prisoner of Azkaban",
		Rating:        "PG",
		TimeInMinutes: 142,
	},
	{
		Name:          "Harry Potter and the Goblet of Fire",
		Rating:        "PG - 13",
		TimeInMinutes: 157,
	},
	{
		Name:          "Harry Potter and the Order of the Phoenix",
		Rating:        "PG - 13",
		TimeInMinutes: 129,
	},
	{
		Name:          "Harry Potter and the Half Blood Prince",
		Rating:        "PG - 13",
		TimeInMinutes: 153,
	},
	{
		Name:          "Harry Potter and the Deathly Hallows: Part 1",
		Rating:        "PG - 13",
		TimeInMinutes: 146,
	},
	{
		Name:          "Harry Potter and the Deathly Hallows: Part 2",
		Rating:        "PG - 13",
		TimeInMinutes: 130,
	},
}

var LOTR = []Movie{
	{
		Name:          "The Fellowship Of the Ring",
		Rating:        "PG-13",
		TimeInMinutes: 178,
	},
	{
		Name:          "The Two Towers",
		Rating:        "PG-13",
		TimeInMinutes: 179,
	},
	{
		Name:          "The Return of The King",
		Rating:        "PG-13",
		TimeInMinutes: 200,
	},
}

func TestCaclulateWatchTime(t *testing.T) {
	type args struct {
		movies []Movie
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "HarryPotter",
			args: args{movies: HarryPotter},
			want: 1170,
		},
		{
			name: "LordOfTheRing",
			args: args{movies: LOTR},
			want: 557,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CaclulateWatchTime(tt.args.movies); got != tt.want {
				t.Errorf("CaclulateWatchTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetDaysLeftInSchool(t *testing.T) {
	type args struct {
		sMonth     int
		eMonth     int
		eDay       int
		isFullYear bool
	}

	endDate := time.Date(2021, 5, 25, 0, 0, 0, 0, time.UTC)
	dur := endDate.Sub(time.Now().UTC())
	fmt.Println(dur)
	days := int(dur.Hours() / 24)
	fmt.Println(days)
	tests := []struct {
		name     string
		args     args
		wantDays int
	}{
		{
			name: "success - 100 days/isFullYear",
			args: args{
				sMonth:     8,
				eMonth:     5,
				eDay:       25,
				isFullYear: true,
			},
			wantDays: days + 1,
		},
	}
	for _, tt := range tests {
		if gotDays := GetDayLeftOfSchool(tt.args.sMonth, tt.args.eMonth, tt.args.eDay, tt.args.isFullYear); gotDays != tt.wantDays {
			t.Errorf("%s: wanted days: %d, got days: %d\n", tt.name, tt.wantDays, gotDays)
		}
	}
}
