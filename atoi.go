package piscine

func Atoi(s string) int {
	if s == "" {
		return 0
	}
	var result int = 0
	Array := []rune(s)
	Negative := 1
	for i := range Array {
		if CheckForValid(Array[i]) {
			if i == 0 {
				if Array[i] == '+' {
					Negative = 1
				} else if Array[i] == '-' {
					Negative = -1
				} else {
					result = result*10 + int(Array[i]) - '0'
				}
			} else {
				if Array[i] == '+' || Array[i] == '-' {
					return 0
				} else {
					result = result*10 + int(Array[i]) - '0'
				}
			}
		} else {
			return 0
		}
	}
	return result * Negative
}

func CheckForValid(c rune) bool {
	numbers := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '+', '-'}
	for i := range numbers {
		if numbers[i] == c {
			return true
		}
	}
	return false
}
