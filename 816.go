package main

import "strconv"

func ambiguousCoordinates(S string) []string {
	length := len(S)
	result := make([]string, 0)

	for i := 2; i < length-1; i++ {
		x, y := S[1:i], S[i:length-1]
		for j := 0; j < len(x); j++ {
			var xf string
			if j == 0 {
				xf = x
			} else {
				xf = x[:j] + "." + x[j:]
			}
			for k := 0; k < len(y); k++ {
				var yf string
				if k == 0 {
					yf = y
				} else {
					yf = y[:k] + "." + y[k:]
				}
				if isValidFloat(xf) && isValidFloat(yf) {
					result = append(result, "("+xf+", "+yf+")")
				}
			}
		}
	}

	return result
}

func isValidFloat(s string) bool {
	// Able to check both integer and float number
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return false
	}
	sf := strconv.FormatFloat(f, 'f', -1, 64)
	return sf == s
}
