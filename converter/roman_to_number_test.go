package converter

import (
	"testing"
)

func TestRomanToNumber(t *testing.T){
	// test cases
	tests := []struct{
		roman_number string; expected int
	}{
	{"I", 1},
    {"II", 2},
    {"III", 3},
    {"IV", 4},
    {"V", 5},
    {"VI", 6},
    {"VII", 7},
    {"VIII", 8},
    {"IX", 9},
    {"X", 10},
    {"XI", 11},
    {"XII", 12},
    {"XIII", 13},
    {"XIV", 14},
    {"XV", 15},
    {"XVI", 16},
    {"XVII", 17},
    {"XVIII", 18},
    {"XIX", 19},
    {"XX", 20},
    {"XXX", 30},
    {"XL", 40},
    {"L", 50},
    {"LX", 60},
    {"LXX", 70},
    {"LXXX", 80},
    {"XC", 90},
    {"C", 100},
    {"CC", 200},
    {"CCC", 300},
    {"CD", 400},
    {"D", 500},
    {"DC", 600},
    {"DCC", 700},
    {"DCCC", 800},
    {"CM", 900},
    {"M", 1000},
    {"MM", 2000},
    {"MMM", 3000},
    {"MMMCMXCIX", 3999},
	}

	for _,test := range tests{
		result, _ := RomanToNumber(test.roman_number)
		if result != test.expected {
			t.Errorf("RomanToNumber(%v) = %d; expected %d", test.roman_number, result, test.expected)
		}
	}
}

func TestRomanEmpty(t *testing.T){
	roman_number := ""
	expect_error := true
	_, err := RomanToNumber(roman_number)

	if (err != nil) != expect_error {
		t.Errorf("RomanToNumber(%q) error = %v; expectError = %v", roman_number, err, true)
	}
}

func TestWrongRomanNumber(t *testing.T){
	// Test cases
	tests := []struct {
		roman_number string
		expect_error  bool
	}{
		{"K", true},
		{"XXB", true},
		{"XXI", false},
	}

	for _, test := range tests{
		_, err := RomanToNumber(test.roman_number)

		if (err != nil) != test.expect_error{
			t.Errorf("RomanToNumber(%q) error = %v; expectError = %v", test.roman_number, err, test.expect_error)
		}
	}
}