package converter

import "testing"

func TestNumberToRoman(t *testing.T){
	tests := []struct {
		number   int
		expected string
	}{
		{1, "I"},
		{2, "II"},
		{3, "III"},
		{4, "IV"},
		{5, "V"},
		{6, "VI"},
		{7, "VII"},
		{8, "VIII"},
		{9, "IX"},
		{10, "X"},
		{20, "XX"},
		{30, "XXX"},
		{40, "XL"},
		{50, "L"},
		{60, "LX"},
		{70, "LXX"},
		{80, "LXXX"},
		{90, "XC"},
		{100, "C"},
		{200, "CC"},
		{300, "CCC"},
		{400, "CD"},
		{500, "D"},
		{600, "DC"},
		{700, "DCC"},
		{800, "DCCC"},
		{900, "CM"},
		{1000, "M"},
		{2000, "MM"},
		{3000, "MMM"},
		{3999, "MMMCMXCIX"},
	}

	for _,test := range tests{
		result, _ := NumberToRoman(test.number)
		if result != test.expected {
			t.Errorf("NumberToRoman(%v) = %v; expected %v", test.number, result, test.expected)
		}
	}
}

func TestNegativeNumber(t *testing.T){
	number := -1
	expect_error := true
	_, err := NumberToRoman(number)

	if (err != nil) != expect_error {
		t.Errorf("RomanToNumber(%q) error = %v; expectError = %v", number, err, true)
	}
}