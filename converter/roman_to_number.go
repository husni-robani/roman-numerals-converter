package converter

import (
	"log"
)

func RomanToNumber(roman_number string) int {
	if roman_number == "" {
		log.Fatal("roman number is empty")
	}
	

	result := 0
	length := len(roman_number)

	for i := 0; i < length; i++ {
		// Get the current and next character values
		currentValue := get_roman_value(string(roman_number[i]))
		nextValue := 0

		if i < length-1 {
			nextValue = get_roman_value(string(roman_number[i+1]))
		}

		// If the current value is less than the next value, we have a subtraction case
		if currentValue < nextValue {
			result -= currentValue
		} else {
			result += currentValue
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