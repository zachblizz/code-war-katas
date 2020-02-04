package main

// Arithmetic - performs the given operation (op) on the two given numbers
func Arithmetic(a, b int, op string) int {
	result := 0

	switch op {
	case "add":
		result = a + b
	case "subtract":
		result = a - b
	case "multiply":
		result = a * b
	case "divide":
		result = a / b
	}

	return result
}
