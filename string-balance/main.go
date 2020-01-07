package main

// BalancedStringSplit - count the amount of balanced strings we can create
func BalancedStringSplit(s string) int {
    if len(s) == 0 || len(s) == 1 {
        return 0
    }
    
    result := 0
    lCount := 0
    start := string(s[0])
    
    for i, c := range s {
		strC := string(c)

        if strC == start {
            lCount++
        } else if strC != start && lCount != 0 {
            lCount--
            
            if lCount == 0 {
				if i+1 < len(s) {
					start = string(s[i+1])
				}

                result++
            }
        }
    }
    
    return result
}
