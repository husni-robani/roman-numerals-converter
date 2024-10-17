package converter

import (
	"log"
	"strings"
)

func RomanToNumber(roman_number string) int {
	if roman_number == ""{
		log.Fatal("roman number is empty")
	}

	splitted_roman := strings.Split(roman_number, "")
	result := 0
	for i:= 0; i < len(splitted_roman); i++ {
		if i == len(splitted_roman) - 1 {
			result += get_roman_value(splitted_roman[i])
		}else {
			if get_roman_value(splitted_roman[i + 1]) > get_roman_value(splitted_roman[i]){
				// do substraction
				result += get_roman_value(splitted_roman[i + 1]) - get_roman_value(splitted_roman[i])
				i++
			}else {
				if get_roman_value(splitted_roman[i + 2]) > get_roman_value(splitted_roman[i + 1]){
					// substract: splittedRoman[i + 2] - splittedRoman[i + 1] = temp
					temp := get_roman_value(splitted_roman[i + 2]) - get_roman_value(splitted_roman[i + 1])
					// addtion: splittedRoman[i] + temp
					result += get_roman_value(splitted_roman[i]) + temp
					i += 2
				}else{
					result += get_roman_value(splitted_roman[i]) + get_roman_value(splitted_roman[i + 1])
					i++
				}
			}
		}
	}
	return result
}

func get_roman_value(roman string) int {
	roman_list := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	value, exist := roman_list[roman]

	if !exist {
		log.Fatalf("%v is not a roman numerals", roman)
	}

	return value
}