package helpgranny

import "math"

// kata - https://www.codewars.com/kata/help-your-granny/train/go

// Tour - find the distance granny has to travel to visit all friends
// a^2 + b^2 = c^2
// h - represents the distance from town[0] - town[i]
func Tour(arrFriends []string, ftwns map[string]string, h map[string]float64) int {
	// given a, c - need to find b
	// b = math.Sqrt(c^2 - a^2)
	totDist := 0.0

	for i, friend := range arrFriends {
		twn := ftwns[friend]
		if twn != "" {
			a := h[twn]

			if totDist == 0 {
				totDist += a
			}

			if i+1 < len(arrFriends) {
				nextFriend := arrFriends[i+1]
				if nextTwn, tOk := ftwns[nextFriend]; tOk {
					if c, distOk := h[nextTwn]; distOk {
						totDist += math.Sqrt(math.Pow(c, 2) - math.Pow(a, 2))
					}
				} else {
					totDist += a
				}
			} else {
				totDist += a
			}
		}
	}

	return int(totDist)
}
