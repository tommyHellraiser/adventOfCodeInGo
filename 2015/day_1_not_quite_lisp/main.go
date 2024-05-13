package main

import "fmt"

const FLOOR_UP = '('
const FLOOR_DOWN = ')'

func main() {
	//input := getTestInput(5)
	input := getInput()
	floor := 0

	//	--- Part 1 ---
	for _, char := range input {
		switch char {
		case FLOOR_UP:
			floor++
			break
		case FLOOR_DOWN:
			floor--
			break
		default:
			fmt.Println("Unsupported char")
		}
	}

	fmt.Println("\nPART 1")
	fmt.Printf("Destination floor is %d\n", floor)

	//	--- Part 2 ---
	//	Define index and reset floor
	defIndex := 0
	floor = 0
	for index, char := range input {
		switch char {
		case FLOOR_UP:
			floor++
			break
		case FLOOR_DOWN:
			floor--
			break
		default:
			fmt.Println("Unsupported char")
		}

		if floor == -1 {
			defIndex = index + 1
			break
		}
	}

	fmt.Println("\nPART 2")
	fmt.Printf("Index of instruction to reach basement -1 is %d\n", defIndex)
}
