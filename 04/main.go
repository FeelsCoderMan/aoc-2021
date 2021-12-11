package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)

const width int = 5

func main() {
	boards, numbersDrawn := readPuzzleInput("input.txt")
	answerPart1 := part1(boards, numbersDrawn)
	answerPart2 := part2(boards, numbersDrawn)
	fmt.Printf("Answer of part1 is %d\n", answerPart1)
	fmt.Printf("Answer of part2 is %d\n", answerPart2)
}

func readPuzzleInput(filePath string) ([][width][width]uint8, []uint8) {
	file, err := os.Open(filePath)

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	boards := [][width][width]uint8{}
	scanner.Scan()
	numbersDrawnStr := scanner.Text()
	numbersDrawnInt := formatDrawnNumbers(numbersDrawnStr)
	currentBoard := 0
	currentRow := 0

	for scanner.Scan() {
		lineString := scanner.Text()
		if len(lineString) == 0 {
			currentBoard++
			currentRow = 0
		} else {
			dumpBoard(&boards, lineString, currentBoard, currentRow)
			currentRow++
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Bingo game has %d boards\n", currentBoard)
	return boards, numbersDrawnInt
}

func dumpBoard(boards *[][width][width]uint8, lineString string, currentBoard int, currentRow int) {
	board := strings.Fields(lineString)

	if (len(*boards) == currentBoard - 1) {
		*boards = append(*boards, [width][width]uint8{})
	}

	for i := 0; i < len(board); i++ {
		val, _ := strconv.Atoi(board[i])
		(*boards)[currentBoard - 1][currentRow][i] = uint8(val)
	}
}

func formatDrawnNumbers(numbersDrawn string) []uint8 {
	numbersDrawnStr := strings.Split(numbersDrawn, ",")
	numbersDrawnInt := []uint8{}

	for _, numberDrawn := range numbersDrawnStr {
		val, _ := strconv.Atoi(numberDrawn)
		numbersDrawnInt = append(numbersDrawnInt, uint8(val))
	}

	return numbersDrawnInt
}

func part1(boards [][width][width]uint8, numbersDrawn []uint8) int {
	savedResult := [][]int{}

	for _, numberDrawn := range numbersDrawn {
		savedUnmarkedAndBoardNum := findAndChange(numberDrawn, &boards)

		if len(savedUnmarkedAndBoardNum) > 0 {
			for i := 0; i < len(savedUnmarkedAndBoardNum); i++ {
				if !isIncluded(&savedResult, savedUnmarkedAndBoardNum[i][0]) {
					savedResult = append(savedResult, savedUnmarkedAndBoardNum[i])
					return savedResult[len(savedResult) - 1][1]
				}
			}
		}
	}

	return -1
}

func findAndChange(numberDrawn uint8,boards *[][width][width]uint8) [][]int {
	savedUnmarkedAndBoardNumber := [][]int{}

	for indexOfBoards := 0; indexOfBoards < len(*boards); indexOfBoards++ {
		indexX, indexY := findNumber(boards, indexOfBoards, numberDrawn)

		if !(indexX < 0 && indexY < 0) {
			changeNumberInBoard(indexX, indexY, numberDrawn, &boards, indexOfBoards)
			isFinished := checkWinCondition(indexX, indexY, boards, indexOfBoards)

			if isFinished {
				sumUnmarkedNumbers := calculateSumUnmarked(boards, indexOfBoards)
				savedUnmarkedAndBoardNumber = append(savedUnmarkedAndBoardNumber, []int{indexOfBoards + 1, sumUnmarkedNumbers * int(numberDrawn)})
			}
		}
	}

	return savedUnmarkedAndBoardNumber
}

func findNumber(boards *[][width][width]uint8, indexOfBoards int, numberDrawn uint8) (int,int) {
	for i := 0; i < len((*boards)[indexOfBoards]); i++ {
		for j := 0; j < len((*boards)[indexOfBoards][0]); j++ {
			if (*boards)[indexOfBoards][i][j] == numberDrawn {
				return i, j
			}
		}
	}

	return -1, -1
}

func changeNumberInBoard(x int, y int, numberDrawn uint8, boards **[][width][width]uint8, currentBoard int) {
	signedNumberDrawn := changeSign(numberDrawn)
	(**boards)[currentBoard][x][y] = signedNumberDrawn
}

func changeSign(num uint8) uint8{
	return num | 0x80
}

func checkWinCondition(x int, y int, boards *[][width][width]uint8, currentBoard int) bool {
	resultY := true
	resultX := true

	for i := 0; i < len((*boards)[currentBoard]); i++ {
		if y == i {
			continue
		}
		if (*boards)[currentBoard][x][i] < 100 {
			resultY = false
			break
		}
	}

	for j := 0; j < len((*boards)[currentBoard][0]); j++ {
		if x == j {
			continue
		}
		if (*boards)[currentBoard][j][y] < 100 {
			resultX = false
			break
		}
	}

	return resultX || resultY
}

func calculateSumUnmarked(boards *[][width][width]uint8, currentBoard int) int{
	currentSum := 0

	for i := 0 ; i < len((*boards)[currentBoard]); i++ {
		for j := 0 ; j < len((*boards)[currentBoard][i]); j++ {
			if (*boards)[currentBoard][i][j] < 100 {
				currentSum += int((*boards)[currentBoard][i][j])
			}
		}
	}

	return currentSum
}

func part2(boards [][width][width]uint8, numbersDrawn []uint8) int {
	savedResult := [][]int{}

	for _, numberDrawn := range numbersDrawn {
		savedUnmarkedAndBoardNum := findAndChange(numberDrawn, &boards)

		if len(savedUnmarkedAndBoardNum) > 0 {
			for i := 0; i < len(savedUnmarkedAndBoardNum); i++ {
				if !isIncluded(&savedResult, savedUnmarkedAndBoardNum[i][0]) {
					savedResult = append(savedResult, savedUnmarkedAndBoardNum[i])
				}
			}
		}
	}

	return savedResult[len(savedResult) - 1][1]
}

func isIncluded(resultArray *[][]int, boardWinner int) bool {
	for i := 0; i < len(*resultArray); i++ {
		if (*resultArray)[i][0] == boardWinner {
			return true
		}
	}

	return false
}
