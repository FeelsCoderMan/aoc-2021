package main

import  (
	"fmt"
	"os"
	"bufio"
	"strconv"
)

func main() {
	var arr = readPuzzleInput("input.txt")
	var answerPart1 = part1(arr)
	var answerPart2 = part2(arr)
	fmt.Printf("Answer of Part 1 is %d\n", answerPart1)
	fmt.Printf("Answer of Part 2 is %d\n", answerPart2)
}

func readPuzzleInput(filePath string) []int {
	file, err := os.Open(filePath)

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	arr := []int{}

	for scanner.Scan() {
		lineString := scanner.Text()
		lineNum, _ := strconv.Atoi(lineString)
		arr = append(arr, lineNum)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	return arr
}

func part1 (arr []int) int{
	countIncrement := 0
	for i := 1; i < len(arr); i++ {
		if arr[i - 1] < arr[i] {
			countIncrement++
		}
	}
	return countIncrement
}

func part2 (arr []int) int {
	tripleSumArray := []int{}

	for i := 0; i < len(arr) - 2; i++ {
		tripleSumArray = append(tripleSumArray, arr[i] + arr[i + 1] + arr[i + 2])
	}

	return part1(tripleSumArray)

}
