package main

import "fmt"

// link - https://www.codewars.com/kata/rectangle-into-squares/train/go

// Point - a point with start row and start column
type Point struct {
	r int
	c int
}

// SquaresInRect - calculates the number of rects in a
// true rectangle ("true" rectangle meaning that the two dimensions are different)
func SquaresInRect(lng, wdth int) []int {
	if lng == wdth {
		return nil
	}

	// rect := makeRect(lng, wdth)
	
	lSquare := lng

	if lSquare > wdth {
		lSquare = wdth
	}

	// TODO: make this dynamic to the given length and width
	ret := []int{}
	canMake, newP := canMakeSquare(lSquare, lng, wdth, Point{0,0})

	if canMake {
		ret = append(ret, lSquare)
		canMake, newP = canMakeSquare(lng-wdth, lng, wdth, newP)

		if canMake {
			ret = append(ret, lng-wdth)
			canMake, newP = canMakeSquare(lSquare-newP.r, lng, wdth, newP)

			if canMake {
				ret = append(ret, lSquare - newP.r)
				canMake, _ = canMakeSquare(lSquare-newP.r, lng, wdth, newP)

				if canMake {
					ret = append(ret, lSquare - newP.r)
				}
			}
		}
	}

	return ret
}

func canMakeSquare(square, rectL, rectW int, p Point) (bool, Point) {
	l := p.c + square
	w := p.r + square

	isGoodSquare := l <= rectL && w <= rectW
	newP := Point{0, 0}

	if isGoodSquare {
		if l < rectL {
			newP.r = p.r
		} else if p.r + w <= rectW {
			newP.r = p.r + w
		}

		if p.c == 0 && l < rectL {
			newP.c = w
		} else {
			newP.c = p.c
		}
	}

	return isGoodSquare, newP
}

func main() {
	// fmt.Println(canMakeSquare(3, 5, 3, Point{0,0}))
	// fmt.Println(canMakeSquare(2, 5, 3, Point{0,3}))
	fmt.Println(SquaresInRect(5, 3))
}
