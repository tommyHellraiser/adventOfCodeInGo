package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
	"time"
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

	hashResult := ""
	numericInput := 0

	//	Measuring execution time
	start := time.Now()

	//	--- Part 1 ---
	for {
		hashInput := input + strconv.FormatInt(int64(numericInput), 10)
		hash := md5.Sum([]byte(hashInput))
		hashResult = hex.EncodeToString(hash[:])

		if hashResult[0:5] == "00000" {
			break
		}

		numericInput++
	}

	fmt.Println("\nPART 1")
	fmt.Printf("Result hash is: %s\n", hashResult)
	fmt.Printf("Result hash input is: %d\n", numericInput)
	fmt.Printf("Execution time was: %s\n", time.Since(start))

	//	--- Part 2 ---
	hashResult = ""
	numericInput = 0

	//	Measuring time again
	start = time.Now()

	for {
		hashInput := input + strconv.FormatInt(int64(numericInput), 10)
		hash := md5.Sum([]byte(hashInput))
		hashResult = hex.EncodeToString(hash[:])

		if hashResult[0:6] == "000000" {
			break
		}

		numericInput++
	}

	fmt.Println("\nPART 2")
	fmt.Printf("Result hash is: %s\n", hashResult)
	fmt.Printf("Result hash input is: %d\n", numericInput)
	fmt.Printf("Execution time was: %s\n", time.Since(start))

}
