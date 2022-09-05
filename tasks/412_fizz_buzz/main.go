package main

import "strconv"

func fizzBuzz(n int) []string {
	var result []string
	for i := 1; i <= n; i++ {
		remainderThree := i % 3
		remainderFive := i % 5

		if remainderThree == 0 && remainderFive == 0 {
			result = append(result, "FizzBuzz")
		} else if remainderThree == 0 {
			result = append(result, "Fizz")
		} else if remainderFive == 0 {
			result = append(result, "Buzz")
		} else {
			result = append(result, strconv.Itoa(i))
		}
	}

	return result
}
