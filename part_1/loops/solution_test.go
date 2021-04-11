package loops

import "testing"

func TestCaclulateWatchTime(t *testing.T) {
	type args struct {
		movies []Movie
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CaclulateWatchTime(tt.args.movies); got != tt.want {
				t.Errorf("CaclulateWatchTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
