package main

import (
	"fmt"
	"strconv"
)

func main() {
	string1 := "1010"
	string2 := "10101"

	someShit := addBinary(string1, string2)

	fmt.Println(someShit)
}

func addBinary(a string, b string) string {
	len1 := len(a)
	len2 := len(b)
	maxLen := 0

	if len1 > len2 {
		maxLen = len1
	} else {
		maxLen = len2
	}

	result := make([]byte, maxLen+1)
	counter := maxLen
	additionalDigit := 0

	for true {
		if counter < 0 {
			break
		}

		currentDigit := additionalDigit

		if len1 > 0 {
			val, _ := strconv.Atoi(string(a[len1-1]))
			currentDigit += val
			len1--
		}
		if len2 > 0 {
			val, _ := strconv.Atoi(string(b[len2-1]))
			currentDigit += val
			len2--
		}

		if currentDigit > 1 {
			additionalDigit = 1
			if currentDigit == 2 {
				result[counter] = '0'
			} else {
				result[counter] = '1'
			}
		} else {
			additionalDigit = 0
			if currentDigit == 1 {
				result[counter] = '1'
			} else {
				result[counter] = '0'
			}
		}

		counter--
	}

	if result[0] != '1' {
		return string(result[1:])
	}

	return string(result)
}
