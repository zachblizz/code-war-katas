package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

// StockList - retuns a string that contains each
// category with the amount of books in each category
func StockList(listArt, listCat []string) string {
	if len(listArt) == 0 {
		return ""
	}

	var result bytes.Buffer
	catMap := make(map[string]int)

	for _, art := range listArt {
		fields := strings.Fields(art)
		bookAmnt, _ := strconv.Atoi(fields[1])
		bookCat := string(fields[0][0])

		if _, ok := catMap[bookCat]; ok {
			catMap[bookCat] += bookAmnt
		} else {
			catMap[bookCat] = bookAmnt
		}
	}

	for i, cat := range listCat {
		result.WriteString(fmt.Sprintf("(%v : %v)", cat, catMap[cat]))

		if i < len(listCat) - 1 {
			result.WriteString(" - ")
		}
	}

	return result.String()
}
