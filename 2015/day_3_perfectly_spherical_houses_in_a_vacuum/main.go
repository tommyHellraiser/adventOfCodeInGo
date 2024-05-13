package main

import (
	"fmt"
)

type Key struct {
	xLoc, yLoc int
}

type MaxValues struct {
	negX, posX, negY, posY int
}

func main() {
	//input := getTestInput()
	input := getInput()

	//	Create map with starting location containing one gift
	housesMap := map[Key]int{}
	location := Key{0, 0}
	housesMap[location] = 1

	//	Initialize the maximum values to identify all the visited houses
	maxValues := MaxValues{0, 0, 0, 0}

	//	--- Part 1 ---
	//	Iterate through all chars to get populate the map
	for _, char := range input {
		//	Move through houses and deploy presents
		housesMap, location = advanceHouses(housesMap, location, char)

		//	Update positional maximum values to iterate at the end
		maxValues = determineMaximums(location, maxValues)
	}

	//	Check how many houses have at least one gift
	amountOfHouses := housesWithAnyGifts(housesMap, maxValues)
	//printHousesMap(housesMap, maxValues)

	fmt.Println("\nPART 1")
	fmt.Printf("%d houses got at least one gift\n", amountOfHouses)

	//	--- Part 2 ---
	//	Initialize new variables
	newHousesMap := map[Key]int{}
	newHousesMap[Key{0, 0}] = 1
	santaLocation := Key{0, 0}
	roboSantaLocation := Key{0, 0}
	newMaxValues := MaxValues{0, 0, 0, 0}

	//	Instead of iterating one by one the chars of location input, we'll iterate 2 by 2 to allow both santas to work
	for index := 0; index < len(input); index += 2 {
		//	Tracking of santa's location and update maximums
		newHousesMap, santaLocation = advanceHouses(newHousesMap, santaLocation, int32(input[index]))
		newMaxValues = determineMaximums(santaLocation, newMaxValues)

		//	Tracking of robo-santa's location and update maximums
		newHousesMap, roboSantaLocation = advanceHouses(newHousesMap, roboSantaLocation, int32(input[index+1]))
		newMaxValues = determineMaximums(roboSantaLocation, newMaxValues)
	}

	roboAmountOfHouses := housesWithAnyGifts(newHousesMap, newMaxValues)

	fmt.Println("\nPART 2")
	fmt.Printf("%d houses got at least one gift\n", roboAmountOfHouses)

	//printHousesMap(newHousesMap, newMaxValues)
}

func advanceHouses(housesMap map[Key]int, location Key, direction int32) (updatedMap map[Key]int, updatedLocation Key) {
	//	Initialize return variables
	updatedMap = housesMap
	updatedLocation = location

	//	Execute movement
	switch direction {
	case '>':
		updatedLocation.xLoc++
		updatedMap[updatedLocation]++
	case '<':
		updatedLocation.xLoc--
		updatedMap[updatedLocation]++
	case '^':
		updatedLocation.yLoc++
		updatedMap[updatedLocation]++
	case 'v':
		updatedLocation.yLoc--
		updatedMap[updatedLocation]++
	default:
		fmt.Println("Unexpected travel direction")
	}

	return updatedMap, updatedLocation
}

func determineMaximums(location Key, currentMax MaxValues) (newMax MaxValues) {

	newMax = currentMax

	//	Updating x-axis locations if they exceed positive or negative extremes
	if location.xLoc > currentMax.posX {
		newMax.posX = location.xLoc
	} else if location.xLoc < currentMax.negX {
		newMax.negX = location.xLoc
	}

	//	Updating y-axis locations if they exceed positive or negative extremes
	if location.yLoc > currentMax.posY {
		newMax.posY = location.yLoc
	} else if location.yLoc < currentMax.negY {
		newMax.negY = location.yLoc
	}

	return newMax
}

func printHousesMap(housesMap map[Key]int, maxValues MaxValues) {

	for yIndex := maxValues.posY; yIndex >= maxValues.negY; yIndex-- {
		for xIndex := maxValues.negX; xIndex <= maxValues.posX; xIndex++ {
			if xIndex == 0 && yIndex == 0 {
				fmt.Print("|")
				fmt.Printf("%d", housesMap[Key{xIndex, yIndex}])
				fmt.Print("|")
			} else {
				fmt.Printf(" %d ", housesMap[Key{xIndex, yIndex}])
			}
		}
		fmt.Println(" ")
		//fmt.Println(" ")
	}
}

func housesWithAnyGifts(housesMap map[Key]int, maxValues MaxValues) (amount int) {

	for xIndex := maxValues.negX; xIndex <= maxValues.posX; xIndex++ {
		for yIndex := maxValues.negY; yIndex <= maxValues.posY; yIndex++ {
			if housesMap[Key{xIndex, yIndex}] > 0 {
				amount++
			}
		}
	}

	return amount
}
