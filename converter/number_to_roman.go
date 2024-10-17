package converter

import "errors"

func NumberToRoman(num int) (string, error) {
	if num < 0 {
		return "", errors.New("numbers must be positive (n >= 0)")
	}
	roman_number := ""
	rest_num := num

	for rest_num > 0{
		result := roman_symbol(rest_num)
		rest_num = result.rest_num
		roman_number = roman_number + result.roman_value
	}

	return roman_number, nil
}

func roman_symbol(num int) struct{
	roman_value string
	rest_num int
} {
	if num - 1000 >= -100{
		// M
		if num - 1000 >= 0{
			return struct{roman_value string; rest_num int}{
				roman_value: "M",
				rest_num: num - 1000,
			}
		}
		return struct{roman_value string; rest_num int}{
			roman_value: "CM",
			rest_num: num - 900,
		}
	} else if (num - 500 >= -100) {
		// D
		if num - 500 >= 0 {
		  return struct{roman_value string; rest_num int}{
			roman_value: "D",
			rest_num: num - 500,
		  }
		}
		return struct{roman_value string; rest_num int}{
			roman_value: "CD",
			rest_num: num - 400,
		}
	  }else if num - 100 >= -10 {
		// C
		if num - 100 >= 0{
			return struct{roman_value string; rest_num int}{
				roman_value: "C",
				rest_num: num - 100,
			}
		}

		return struct{roman_value string; rest_num int}{
			roman_value: "XC",
			rest_num: num - 90,
		}
	  }else if num - 50 >= -10 {
		 // L
		if num - 50 >= 0{
			return struct{roman_value string; rest_num int}{
				roman_value: "L",
				rest_num: num - 50,
			}
		}
		return struct{roman_value string; rest_num int}{
			roman_value: "XL",
			rest_num: num - 40,
		}
	  }else if num - 10 >= -1{
		// X
		if num - 10 >= 0{
			return struct{roman_value string; rest_num int}{
				roman_value: "X",
				rest_num: num - 10,
			}
		}
		return struct{roman_value string; rest_num int}{
			roman_value: "IX",
			rest_num: 0,
		}
	  }else if num - 5 >= -1 {
		// V
		if num - 5 >= 0{
			return struct{roman_value string; rest_num int}{
				roman_value: "V",
				rest_num: num - 5,
			}
		}
		return struct{roman_value string; rest_num int}{
			roman_value: "IV",
			rest_num: 0,
		}
	  }else if num - 1 >= 0{
		// I
		return struct{roman_value string; rest_num int}{
			roman_value: "I",
			rest_num: num - 1,
		}
	  } else {
		return struct{roman_value string; rest_num int}{
			roman_value: "",
			rest_num: 0,
		}
	  }
}