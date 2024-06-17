package main

import (
	"fmt"
	"os"
	"strconv"
)

const (
	I = 1
	V = 5
	X = 10
	L = 50
	C = 100
	D = 500
	M = 1000
)

func main() {
	if len(os.Args) == 1 {
		repl()
	} else {
		if roman, err := ParseAndConvert(os.Args[1]); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		} else {
			fmt.Println(roman)
		}
	}
}

// repl is a read-eval-print-loop for interactively converting a bunch of numbers from
// Arabic to Roman numerals.
func repl() {
	fmt.Println("Arabic to Roman numeral converter v1.0.0")

	for {
		fmt.Print(">>> ")

		var arabic string
		if _, err := fmt.Scanln(&arabic); err != nil {
			if err.Error() != "unexpected newline" {
				fmt.Fprintln(os.Stderr, err)
			}
		} else {
			if roman, err := ParseAndConvert(arabic); err != nil {
				fmt.Fprintln(os.Stderr, err)
			} else {
				fmt.Println(roman)
			}
		}
	}
}

// ParseAndConvert parses a string of an integer and converts it from Arabic to Roman
// numerals.
func ParseAndConvert(arabic string) (roman string, err error) {
	if arabicInt64, err := strconv.ParseInt(arabic, 10, 64); err != nil {
		return "", fmt.Errorf("expected integer")
	} else {
		return Convert(int(arabicInt64))
	}
}

// Convert converts a number from Arabic numerals to Roman numerals. Currently, only
// inputs within the inclusive range 1-3999 are supported.
func Convert(arabic int) (roman string, err error) {
	if arabic >= 4000 {
		return "", fmt.Errorf("unable to convert numbers greater than 3999")
	} else if arabic <= 0 {
		return "", fmt.Errorf("unable to convert numbers less than 1")
	}

	arabic, newRoman := partialConvert(arabic, D, M, M)
	roman += newRoman

	arabic, newRoman = partialConvert(arabic, C, D, M)
	roman += newRoman

	arabic, newRoman = partialConvert(arabic, L, C, D)
	roman += newRoman

	arabic, newRoman = partialConvert(arabic, X, L, C)
	roman += newRoman

	arabic, newRoman = partialConvert(arabic, V, X, L)
	roman += newRoman

	arabic, newRoman = partialConvert(arabic, I, V, X)
	roman += newRoman

	_, newRoman = partialConvert(arabic, I, I, V)
	roman += newRoman

	return roman, nil
}

// partialConvert converts part of a number from Arabic numerals to Roman numerals. The
// input named arabic is the starting number in Arabic numerals. aLesser is the number
// in Arabic numerals with the value of the Roman numeral of lesser value than that of
// aCompare. Same for aGreater, but greater. aCompare is the number in Arabic numerals
// with the value of the Roman numeral that should be compared to the starting number.
// The returned values are the remaining number in Arabic numerals, and the new part of
// the resulting number in Roman numerals.
func partialConvert(arabic, aLesser, aCompare, aGreater int) (int, string) {
	var roman string
	rLesser := getNumeral(aLesser)
	rCompare := getNumeral(aCompare)
	rGreater := getNumeral(aGreater)

	for arabic >= aCompare {
		if arabic == aCompare {
			roman += rCompare
			return 0, roman
		} else if aCompare != M && arabic >= aCompare*4 {
			roman += rCompare + rGreater
			arabic -= aGreater - aCompare
		} else if aCompare != M && aCompare != I && arabic >= aGreater-aLesser {
			roman += rLesser + rGreater
			arabic -= aGreater - aLesser
		} else {
			roman += rCompare
			arabic -= aCompare
		}
	}

	return arabic, roman
}

// getNumeral converts a number from Arabic numerals to one Roman numeral. If the input
// cannot be converted to one Roman numeral, this function panics.
func getNumeral(arabic int) (romanNumeral string) {
	switch arabic {
	case I:
		return "I"
	case V:
		return "V"
	case X:
		return "X"
	case L:
		return "L"
	case C:
		return "C"
	case D:
		return "D"
	case M:
		return "M"
	default:
		panic(fmt.Sprintf("invalid input: %d", arabic))
	}
}
