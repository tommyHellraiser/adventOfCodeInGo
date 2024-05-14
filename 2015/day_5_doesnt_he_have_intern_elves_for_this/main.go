package main

import (
	"fmt"
	"strings"
)

func main() {
	//input := getTestInput()
	//input := getTestInput2()
	//input := getTestInput3()
	input := getInput()

	//	Split input by lines
	lines := strings.Split(input, "\n")

	//	Initialize vars
	niceStringsAmount := 0

	//	--- Part 1 ---
	for _, line := range lines {
		//	Step 1: check for forbidden strings
		if hasForbiddenStrings(line) {
			continue
		}

		//	Step 2: contains at least 3 vowels
		if !containsThreeVowels(line) {
			continue
		}

		//	Step 3: Contains at least one instance of two duplicate letters in a row
		if !hasDuplicatesInRow(line) {
			continue
		}

		//	Otherwise, it's a nice string
		niceStringsAmount++
	}

	fmt.Println("\nPART 1")
	fmt.Printf("There are %d nice strings in input\n", niceStringsAmount)

	//	---	Part 2 ---
	niceStringsAmount = 0

	for _, line := range lines {
		//	Step 1: check for a pair of letters wit no overlap
		if !hasTwoPairOfLetters(line) {
			continue
		}

		//	Step 2: check for something like exe or efe or asa
		if !hasRepeatedWithOneInBetween(line) {
			continue
		}

		//	Otherwise, it's a nice string
		niceStringsAmount++
	}

	fmt.Println("\nPART 2")
	fmt.Printf("There are %d nice strings in input\n", niceStringsAmount)

}

func hasForbiddenStrings(line string) (forbidden bool) {

	if strings.Contains(line, "ab") {
		return true
	}
	if strings.Contains(line, "cd") {
		return true
	}
	if strings.Contains(line, "pq") {
		return true
	}
	if strings.Contains(line, "xy") {
		return true
	}

	return false
}

func containsThreeVowels(line string) (hasThreeVowels bool) {

	vowelCounter := 0

	//	Counter the number of vowels and return early if threshold was reached early
	vowelCounter += strings.Count(line, "a")
	if vowelCounter >= 3 {
		return true
	}
	vowelCounter += strings.Count(line, "e")
	if vowelCounter >= 3 {
		return true
	}
	vowelCounter += strings.Count(line, "i")
	if vowelCounter >= 3 {
		return true
	}
	vowelCounter += strings.Count(line, "o")
	if vowelCounter >= 3 {
		return true
	}
	vowelCounter += strings.Count(line, "u")

	return vowelCounter >= 3
}

func hasDuplicatesInRow(line string) (hasDuplicates bool) {

	for index := 0; index < len(line)-1; index++ {
		if line[index] == line[index+1] {
			return true
		}
	}

	return false
}

func hasTwoPairOfLetters(line string) (thereAreTwoPairs bool) {

	for index := 0; index < len(line)-1; index++ {
		if strings.Contains(line[index+2:], line[index:index+2]) {
			return true
		}
	}

	return false
}

func hasRepeatedWithOneInBetween(line string) (hasRepeated bool) {

	for index := 0; index < len(line)-2; index++ {
		if line[index] == line[index+2] {
			return true
		}
	}

	return false
}
