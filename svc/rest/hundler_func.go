package rest

import "strconv"

func FizzBuzz(firstString string, secondString string, firstNumber int, secondNumber int, limit int) (result []string) {
	for i := 1; i <= limit; i++ {

		if i%firstNumber == 0 && i%secondNumber == 0 {
			result = append(result, firstString+secondString)
		} else if i%firstNumber == 0 {
			// Multiple of int1
			result = append(result, firstString)
		} else if i%secondNumber == 0 {
			// Multiple of int2
			result = append(result, secondString)
		} else {
			// Neither, so print the number itself
			result = append(result, strconv.Itoa(i))
		}
	}
	return result
}
