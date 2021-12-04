package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)

func main() {
	puzzleInput := readPuzzleInput("input.txt")
	answerPart1 := part1(puzzleInput)
	answerPart2 := part2(puzzleInput)

	fmt.Printf("Answer of part1 is %d\n", answerPart1)
	fmt.Printf("Answer of part2 is %d\n", answerPart2)


}

func readPuzzleInput(filePath string) []string {
	file, err := os.Open(filePath)

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	arr := []string{}

	for scanner.Scan() {
		lineString := scanner.Text()
		arr = append(arr, lineString)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	return arr
}

func part1(puzzleInput []string) int {
	x := 0
	depth := 0

	for _, command := range puzzleInput {
		direction, val := splitCommand(command)

		switch direction {
		case "forward":
			x += val
		case "down":
			depth += val
		case "up":
			depth -= val
		}
	}
	return x * depth
}

func splitCommand(command string) (string, int) {
	splitBySpace := strings.Fields(command)
	val, _ := strconv.Atoi(splitBySpace[1])
	return splitBySpace[0], val
}

func part2(puzzleInput []string) int {
	x := 0
	depth := 0
	aim := 0

	for _, command := range puzzleInput {
		direction, val := splitCommand(command)

		switch direction {
		case "forward":
			x += val
			depth += aim * val
		case "down":
			aim += val
		case "up":
			aim -= val
		}
	}
	return x * depth
}
