package main

func convert(s string, numRows int) string {
	matrix := make([][]rune, numRows)

	chars := []rune(s)

	lineNumber := 0
	down := true
	for _, ch := range chars {
		matrix[lineNumber] = append(matrix[lineNumber], ch)

		if numRows > 1 {
			ternary(down, func() { lineNumber++ }, func() { lineNumber-- })

			if lineNumber == numRows {
				down = false
				lineNumber = lineNumber - 2
			} else if lineNumber == -1 {
				down = true
				lineNumber = lineNumber + 2
			}
		}
	}

	result := ""
	for _, row := range matrix {
		for _, char := range row {
			result += string(char)
		}
	}

	return result
}

func ternary(condition bool, a, b func()) {
	if condition {
		a()
	} else {
		b()
	}
}
