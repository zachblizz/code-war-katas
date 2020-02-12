package main

// lc - https://leetcode.com/problems/zigzag-conversion/

import (
	"fmt"
	"bytes"
	"strings"

)

// not very efficient...
func convert(s string, numRows int) string {
	if len(s) < 3 || numRows < 2 {
		return s
	}

	var ret bytes.Buffer
	order := [][]string{}
	row := 0
	col := 0
	goingUp := false

	for _, c := range s {
		char := string(c)

		if row < numRows {
			if len(order) == row {
				order = append(order, []string{char})
			} else {
				for len(order[row]) < col {
					order[row] = append(order[row], " ")
				}

				order[row] = append(order[row], char)
			}

			if goingUp && row > 0 {
				row--
				col++
			} else {
				row++
				goingUp = false
			}
		} else {
			goingUp = true
			row -= 2
			order[row] = append(order[row], char)

			if row > 0 {
				row--
			} else {
				row++
			}

			col += 2
		}
	}

	for _, row := range order {
		for _, c := range row {
			if c != " " {
				ret.WriteString(c)	
			}
		}
	}

	return ret.String()
}


// ME LIKEY - #stolen
func convertTwo(s string, numRows int) string {
    if len(s) == 0 || numRows == 0 {
        return ""
	}

    if numRows == 1 {
        return s
	}

    rowArr := make([]string, numRows)
    cnt := 0
	down := true

    for _, c := range s {
		rowArr[cnt] += string(c)

        if cnt == 0 {
            down=true
            cnt++
        } else if cnt == numRows-1 {
            down=false
            cnt--
        } else if down == true {
            cnt++
        } else if down == false {
            cnt--
        }
	}

    retStr := strings.Join(rowArr,"")
    
    return retStr
}

func main() {
	// fmt.Println(convert("HELLO", 3))
	// fmt.Println(convert("PAYPALISHIRING", 4))
	// fmt.Println(convert("ABC", 2))
	fmt.Println(convert("ABCD", 2))
	fmt.Println(convertTwo("PAYPALISHIRING", 4))
}
