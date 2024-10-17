# Roman Numerals Converter

`roman-numerals-converter` is a Go library for converting between Roman numerals and numeric values. It provides two main functions: `RomanToNumber` for converting Roman numerals to numeric values, and `NumberToRoman` for converting numeric values to Roman numerals.

## Installation

To install the library, use `go get`:

```bash
    go get github.com/husni-robani/roman-numerals-converter
```

## Usage/Examples

```go
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
```

## Error Handling

Both functions return errors that you can handle in your application. This allows you to manage how errors are logged or displayed to the user.
