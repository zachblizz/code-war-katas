package main

// lc - https://leetcode.com/problems/valid-parentheses

// not the best approach...
func isValid(s string) bool {
	if len(s) == 0 {
		return true
	} else if len(s) == 1 {
		return false
	}

	closedMap := map[string]string {
		")": "(",
		"]": "[",
		"}": "{",
	}

	prev := string(s[0])
	for i := 1; i < len(s); i++ {
		curr := string(s[i])
		if _, pOk := closedMap[prev]; !pOk {
			if cO, cOk := closedMap[curr]; cOk {	
				if cO != prev {
					return false
				}
			}
		}

		prev = curr
	}

	return true
}

// much better...
func isValidStack(s string) bool {
	if len(s) == 0 {
		return true
	} else if len(s) == 1 {
		return false
	}

	closedMap := map[string]string {
		")": "(",
		"]": "[",
		"}": "{",
	}

	stack := []string{}

	for _, c := range s {
		cStr := string(c)
		
		if o, ok := closedMap[cStr]; ok && len(stack) > 0 {
			p := stack[len(stack)-1]
			if p != o {
				break
			}

			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, cStr)
		}
	}

	return len(stack) == 0
}
