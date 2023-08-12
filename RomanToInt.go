package main

var Romans = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func RomanToInt(s string) int {
	result := 0

	for i := 0; i < len(s); i++ {
		n := Romans[s[i]]

		// if we are not the last character
		// check if the next character is greater than current character
		// if it is - then subtract the number instead of adding it
		// e.g in IV the I subtracts 1 and V adds 5 - so the result is 4
		if i < len(s)-1 {
			if Romans[s[i+1]] > n {
				result -= n
				continue
			}
		}

		result += n
	}
	return result
}
