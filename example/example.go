package main

import (
	"fmt"

	"github.com/husni-robani/roman-numerals-converter/converter"
)

func main(){
	roman := "XIV"
    number_result, err := converter.RomanToNumber(roman)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Roman numeral", roman, "is", number_result)
    }

    number := 14
    roman_result, err := converter.NumberToRoman(number)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Numeric value", number, "is", roman_result)
    }
}