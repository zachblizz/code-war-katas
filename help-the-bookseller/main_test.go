package main

import "testing"

func TestStockList(t *testing.T) {
	type args struct {
		listArt []string
		listCat []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"should give the correct Categories and number of books in that category",
			args{
				[]string{"BBAR 150", "CDXE 515", "BKWR 250", "BTSQ 890", "DRTY 600"},
				[]string{"A", "B", "C", "D"},
			},
			"(A : 0) - (B : 1290) - (C : 515) - (D : 600)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StockList(tt.args.listArt, tt.args.listCat); got != tt.want {
				t.Errorf("StockList() = %v, want %v", got, tt.want)
			}
		})
	}
}
