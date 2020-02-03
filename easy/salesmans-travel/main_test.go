package main

import (
	"strings"
	"testing"
)

var ad = `123 Main Street St. Louisville OH 43071, 432 Main Long Road St. Louisville OH 43071,786 High Street Pollocksville NY 56432,
  54 Holy Grail Street Niagara Town ZP 32908, 3200 Main Rd. Bern AE 56210,1 Gordon St. Atlanta RE 13000,
  10 Pussy Cat Rd. Chicago EX 34342, 10 Gordon St. Atlanta RE 13000, 58 Gordon Road Atlanta RE 13000,
  22 Tokyo Av. Tedmondville SW 43098, 674 Paris bd. Abbeville AA 45521, 10 Surta Alley Goodtown GG 30654,
  45 Holy Grail Al. Niagara Town ZP 32908, 320 Main Al. Bern AE 56210, 14 Gordon Park Atlanta RE 13000,
  100 Pussy Cat Rd. Chicago EX 34342, 2 Gordon St. Atlanta RE 13000, 5 Gordon Road Atlanta RE 13000,
  2200 Tokyo Av. Tedmondville SW 43098, 67 Paris St. Abbeville AA 45521, 11 Surta Avenue Goodtown GG 30654,
  45 Holy Grail Al. Niagara Town ZP 32918, 320 Main Al. Bern AE 56215, 14 Gordon Park Atlanta RE 13200,
  100 Pussy Cat Rd. Chicago EX 34345, 2 Gordon St. Atlanta RE 13222, 5 Gordon Road Atlanta RE 13001,
  2200 Tokyo Av. Tedmondville SW 43198, 67 Paris St. Abbeville AA 45522, 11 Surta Avenue Goodville GG 30655,
  2222 Tokyo Av. Tedmondville SW 43198, 670 Paris St. Abbeville AA 45522, 114 Surta Avenue Goodville GG 30655,
  2 Holy Grail Street Niagara Town ZP 32908, 3 Main Rd. Bern AE 56210, 77 Gordon St. Atlanta RE 13000,
  100 Pussy Cat Rd. Chicago OH 13201`

func TestTravel(t *testing.T) {
	cases := []struct {
		zip  string
		want string
	}{
		{"AA 45522", "AA 45522:Paris St. Abbeville,Paris St. Abbeville/67,670"},
	}

	for _, c := range cases {
		got := Travel(ad, c.zip)

		if strings.Compare(got, c.want) != 0 {
			t.Errorf("[%v]: got =%v= but want %v", c.zip, got, c.want)
		}
	}
}

func TestGetAddressesWithZip(t *testing.T) {
	cases := []struct {
		zip  string
		want map[string][]string
	}{
		{"AA 45522", map[string][]string{
			"houseNum": []string{"67", "670"},
			"street":   []string{"Paris St. Abbeville", "Paris St. Abbeville"},
		}},
	}

	for _, c := range cases {
		got := getAddressesWithZip(ad, c.zip)

		if len(got["houseNum"]) != len(c.want["houseNum"]) {
			t.Errorf("[%v]: got %v but want %v", c.zip, got, c.want)
		}

		for key, gotVal := range got {
			if len(gotVal) != len(c.want[key]) {
				t.Errorf("wrong number of %v", key)
			}

			for i, addr := range gotVal {
				if strings.Compare(addr, c.want[key][i]) != 0 {
					t.Errorf("[%v]: got =%v= but want %v", c.zip, addr, c.want[key][i])
				}
			}
		}
	}
}
