package helpgranny

import "testing"

func TestTour(t *testing.T) {
	type args struct {
		arrFriends []string
		ftwns      map[string]string
		h          map[string]float64
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"should get 889",
			args{
				[]string{"A1", "A2", "A3", "A4", "A5"},
				map[string]string{"A1": "X1", "A2": "X2", "A3": "X3", "A4": "X4"},
				map[string]float64{"X1": 100.0, "X2": 200.0, "X3": 250.0, "X4": 300.0},
			},
			889,
		},
		{
			"should get 168",
			args{
				[]string{ "B1", "B2" },
				map[string]string{"B1": "Y1", "B2": "Y2", "B3": "Y3", "B4": "Y4", "B5": "Y5"},
				map[string]float64{"Y1": 50.0, "Y2": 70.0, "Y3": 90.0, "Y4": 110.0, "Y5": 150.0},
			},
			168,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Tour(tt.args.arrFriends, tt.args.ftwns, tt.args.h); got != tt.want {
				t.Errorf("Tour() = %v, want %v", got, tt.want)
			}
		})
	}
}
