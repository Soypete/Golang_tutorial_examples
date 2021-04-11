package loops

import "testing"

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
