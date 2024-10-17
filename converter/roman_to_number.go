package converter

import (
	"strings"
)

var roman_list = map[string]int{
	"I": 1,
    "V": 5,
    "X": 10,
    "L": 50,
    "C": 100,
    "D": 500,
    "M": 1000,
}

func RomanToNumber(roman_number string) int {
	splitted_roman := strings.Split(roman_number, "")
	result := 0
	for i:= 0; i < len(splitted_roman); i++ {
		if i == len(splitted_roman) - 1 {
			result += roman_list[splitted_roman[i]]
		}else {
			if roman_list[splitted_roman[i + 1]] > roman_list[splitted_roman[i]]{
				// do substraction
				result += roman_list[splitted_roman[i + 1]] - roman_list[splitted_roman[i]]
				i++
			}else {
				if roman_list[splitted_roman[i + 2]] > roman_list[splitted_roman[i + 1]]{
					// substract: splittedRoman[i + 2] - splittedRoman[i + 1] = temp
					temp := roman_list[splitted_roman[i + 2]] - roman_list[splitted_roman[i + 1]]
					// addtion: splittedRoman[i] + temp
					result += roman_list[splitted_roman[i]] + temp
					i += 2
				}else{
					result += roman_list[splitted_roman[i]] + roman_list[splitted_roman[i + 1]]
					i++
				}
			}
		}
	}
	return result
}