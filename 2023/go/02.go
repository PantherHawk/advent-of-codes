/*
--- Part Two ---
Your calculation isn't quite right.
It looks like some of the digits are actually spelled out with letters:
one, two, three, four, five, six, seven, eight, and nine also count as valid "digits".

Equipped with this new information, you now need to find the real first and last digit on each line. For example:

two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen
In this example, the calibration values are 29, 83, 13, 24, 42, 14, and 76. Adding these together produces 281.

What is the sum of all of the calibration values?
*/
package main

import (
	"fmt"
	"strconv"

	"github.com/PantherHawk/adventofcode/2023/go/reader"
)

const biggestNumber = 1000

func main() {
	data := reader.GetFileInput("./2023/input/02.txt")
	var sum int

	for _, datum := range data {
		sum += handle(datum)
	}
	fmt.Printf("total: %d", sum)
}

func handle(line string) int {
	var first, last string
	// search line for a number-word from the begining
	firstNumWord, idxFirstWord := searchForNumberWord(line, false)
	// search line for a number-word starting from the end
	reversed := reverse(line)
	lastNumWord, idxLastWord := searchForNumberWord(reversed, true)

	// search for a number from beginning, save index
	firstInt, idxFirstInt := searchForFirstInteger(line)
	// search for a number from the end, save index
	lastInt, idxLastInt := searchForLastInteger(line)

	// if index of start number greater than start word
	if len(firstNumWord) > 0 {

		if idxFirstWord < idxFirstInt {
			first = firstNumWord
		} else {
			first = firstInt
		}
	} else {
		first = firstInt
	}
	if len(lastNumWord) > 0 {

		if idxLastWord > idxLastInt {
			last = lastNumWord
		} else {
			last = lastInt
		}
	} else {
		last = lastInt
	}
	calibrated, _ := strconv.Atoi(first + last)
	return calibrated
}

func numberWordSearch(line string) {

}

var nums = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func searchForNumberWord(line string, reversed bool) (string, int) {
	// iterate
	for i := 0; i < len(line); i++ {
		for j := i + 1; j < len(line); j++ {
			candidate := line[i:j]
			if reversed {
				candidate = reverse(candidate)
			}
			if v, ok := nums[candidate]; ok {
				if reversed {
					return v, len(line) - i
				}
				return v, i
			}
		}
	}

	// if character is number
	// break
	// add character to accumulated word
	// if word in accumulated is a number
	// break
	return "", biggestNumber
}

func reverse(input string) string {
	n := 0
	runes := make([]rune, len(input))
	for _, r := range input {
		runes[n] = r
		n++
	}
	runes = runes[0:n]
	// Reverse
	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}
	// Convert back to UTF-8.
	return string(runes)
}

func searchForFirstInteger(s string) (string, int) {
	// find first number
	var first int
	j := 1
	for i := 0; i < len(s); i++ {
		n, err := strconv.Atoi(s[i:j])
		j++
		if err != nil {
			continue
		} else {
			first = n
			break
		}
	}

	f := strconv.Itoa(first)
	return f, j - 1
}

func searchForLastInteger(s string) (string, int) {
	// find last number
	var last int
	j := len(s) - 1
	for i := len(s); i >= 1; i-- {

		n, err := strconv.Atoi(s[j:i])
		j--
		if err != nil {
			fmt.Println("Not a number!")
			continue
		} else {
			last = n
			break
		}
	}
	l := strconv.Itoa(last)
	return l, j + 1
}
