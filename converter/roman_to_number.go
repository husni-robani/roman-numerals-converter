package converter

import (
	"errors"
	"fmt"
	"strings"
)

func RomanToNumber(roman_number string) (int, error) {
	roman_number = strings.ToUpper(roman_number)
	if roman_number == "" {
		return 0, errors.New("roman number is empty")
	}
	

	result := 0
	length := len(roman_number)

	for i := 0; i < length; i++ {
		// Get the current and next character values
		currentValue, err := get_roman_value(string(roman_number[i]))
		if err != nil{
			return 0, err
		}
		nextValue := 0

		if i < length-1 {
			nextValue, err = get_roman_value(string(roman_number[i+1]))
			if err != nil{
				return 0, err
			}
		}

		// If the current value is less than the next value, we have a subtraction case
		if currentValue < nextValue {
			result -= currentValue
		} else {
			result += currentValue
		}
	}

	return result, nil
}

func get_roman_value(roman string) (int, error) {
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
		return 0, fmt.Errorf("%v is not a roman numerals", roman)
	}

	return value, nil
}