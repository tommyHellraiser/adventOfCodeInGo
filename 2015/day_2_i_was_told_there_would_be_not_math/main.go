package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//input := getTestInput()
	input := getInput()

	//	Split input into each line
	lines := strings.Split(input, "\n")

	//	--- Part 1 ---
	//	Initialize array of partial measures
	var totalWrappingPaper []int
	//	Parse lines by measures. Order is LxWxH
	for _, line := range lines {
		//	Parse measures from lines
		length, width, height := getMeasuresFromLine(line)

		//	Calculate measure for current present and append into totals array
		paperMeasure := calculateWrappingForPresent(length, width, height)
		totalWrappingPaper = append(totalWrappingPaper, paperMeasure)
	}

	//	Sum all elements in totals array and display
	var totalDimension int
	for _, dimension := range totalWrappingPaper {
		totalDimension += dimension
	}

	fmt.Println("\nPART 1")
	fmt.Printf("The total amount of wrapping paper needed is: %d\n", totalDimension)

	//	--- Part 2 ---
	//	Initialize array of partial measures
	var totalRibbonLength []int
	//	Parse lines by measures. Order is LxWxH
	for _, line := range lines {
		//	Parse measures from lines
		length, width, height := getMeasuresFromLine(line)

		//	Calculate measure for current present and append into totals array
		ribbonLength := calculateRibbonForPresent(length, width, height)
		totalRibbonLength = append(totalRibbonLength, ribbonLength)
	}

	//	Sum all elements in totals array and display
	var totalLength int
	for _, length := range totalRibbonLength {
		totalLength += length
	}

	fmt.Println("\nPART 2")
	fmt.Printf("The total length of ribbon needed is: %d\n", totalLength)
}

func getMeasuresFromLine(line string) (length, width, height int) {
	measures := strings.Split(line, "x")
	length, err1 := strconv.Atoi(measures[0])
	width, err2 := strconv.Atoi(measures[1])
	height, err3 := strconv.Atoi(measures[2])

	if err1 != nil || err2 != nil || err3 != nil {
		err := fmt.Sprintf("Failed to parse measures:\n%s, %s, %s", err1, err2, err3)
		panic(err)
	}

	return
}

func calculateWrappingForPresent(length, width, height int) (measure int) {

	//	Determine the smallest side of the present
	smallest := length * width
	if length*height < smallest {
		smallest = length * height
	}
	if width*height < smallest {
		smallest = width * height
	}

	//	Calculate total measure for this present and return
	measure = (2 * length * width) + (2 * width * height) + (2 * height * length) + smallest

	return measure
}

func calculateRibbonForPresent(length, width, height int) (ribbon int) {

	//	Determine the dimensions of the smallest side
	smallest1 := length
	smallest2 := width
	if length*height < smallest1*smallest2 {
		smallest1 = length
		smallest2 = height
	}
	if width*height < smallest1*smallest2 {
		smallest1 = width
		smallest2 = height
	}

	//	Calculate perimeter to get wrap around ribbon length
	ribbon = (2 * smallest1) + (2 * smallest2)

	//	Calculate present volume for the bow
	ribbon += length * width * height

	return ribbon
}
