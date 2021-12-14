package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)

const dayPart1 int = 80
const dayPart2 int = 256
const tableLen int = 9

func main() {
	puzzleInput := readPuzzleInput("input.txt")
	answerPart1 := part1(puzzleInput, dayPart1)
	answerPart2 := part1(puzzleInput, dayPart2)
	printNumberOfFish(answerPart1, dayPart1)
	printNumberOfFish(answerPart2, dayPart2)
}

func printNumberOfFish(numberOfFish int64, day int) {
	fmt.Printf("After %d days, there are a total of %d fish\n", day, numberOfFish)
}

func readPuzzleInput(filePath string) []int {
	file, err := os.Open(filePath)

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	puzzleInput := []int{}

	for scanner.Scan() {
		lineString := scanner.Text()
		lanternfishAppender(lineString, &puzzleInput)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	return puzzleInput
}

func lanternfishAppender(lineString string, puzzleInputArr *[]int) {
	lineStringArr := strings.Split(lineString, ",")

	for _, lineStringItem  := range lineStringArr {
		val, _ := strconv.Atoi(lineStringItem)
		*puzzleInputArr = append(*puzzleInputArr, val)
	}
}

func part1(puzzleInput []int, day int) int64 {
	currentTable := [tableLen]int{}
	backupTable := [tableLen]int{}
	dumpTable(&puzzleInput, &currentTable)

	for i := 0; i < day; i++ {
		nextDay(&currentTable, &backupTable)
	}

	return calculateNumberOfFish(&backupTable)
}

func calculateNumberOfFish (backupTable *[tableLen]int) int64 {
	var result int64 = 0

	for i := 0; i < tableLen; i++ {
		result += int64((*backupTable)[i])
	}

	return result
}

func nextDay(currentTable *[tableLen]int, backupTable *[tableLen]int) {
	cleanTable(&backupTable)
	(*backupTable)[6] += (*currentTable)[0]
	(*backupTable)[8] += (*currentTable)[0]

	for i := 1; i < tableLen; i++ {
		(*backupTable)[i - 1] += (*currentTable)[i]
	}

	*currentTable = *backupTable
}

func cleanTable(backupTable **[tableLen]int) {
	(**backupTable)[0] = 0

	for i := 1; i < tableLen; i *= 2 {
		copy((**backupTable)[i:], (**backupTable)[:i])
	}
}

func dumpTable(puzzleInput *[]int, currentTable *[tableLen]int) {
	for _, val := range *puzzleInput {
		(*currentTable)[val]++
	}
}
