package roman_to_integer

import "strings"

func romanToInt(s string) int {

	romanNumerals := generateTableRomanNumerals()
	letters := strings.Split(s, "")
	result := 0

	if len(letters) == 1 {
		return romanNumerals[letters[0]]
	}

	for i := 0; i < len(letters); i++ {
		letter1 := letters[i]
		result += romanNumerals[letter1]
		if i + 1 != len(letters){
			letter2 := letters[i+1]
			if letter1 == "I" && (letter2 == "V" || letter2 == "X" ) {
				value := romanNumerals[letter2] - romanNumerals[letter1]
				result += value - romanNumerals[letter1]
				i++
			}   else if letter1 == "X" && (letter2 == "L"|| letter2 == "C" ) {
				value :=  romanNumerals[letter2] - romanNumerals[letter1]
				result += value - romanNumerals[letter1]
				i++
			}   else if letter1 == "C" && (letter2 == "D" || letter2 == "M" ) {
				value := romanNumerals[letter2] - romanNumerals[letter1]
				result += value - romanNumerals[letter1]
				i++
			}
		}

	}

	return result



}

func generateTableRomanNumerals() map[string]int {
	romanNumerals := make(map[string]int)
	romanNumerals["I"] = 1
	romanNumerals["V"] = 5
	romanNumerals["X"] = 10
	romanNumerals["L"] = 50
	romanNumerals["C"] = 100
	romanNumerals["D"] = 500
	romanNumerals["M"] = 1000
	return romanNumerals
}
