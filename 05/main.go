package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"strings"
)

const coordinateLen int = 2
const boardLen int = 1010

type line struct {
	x1 int
	y1 int
	x2 int
	y2 int
}


func main() {
	puzzleInput := readPuzzleInput("input.txt")
	answerPart1 := part1(puzzleInput)
	answerPart2 := part2(puzzleInput)
	fmt.Printf("Answer of part1 is %d\n", answerPart1)
	fmt.Printf("Answer of part2 is %d\n", answerPart2)
}

func showBoard(board *[boardLen][boardLen]int) {
	for i := 0; i < len((*board)); i++ {
		for j:= 0; j < len((*board)[i]); j++ {
			if (*board)[i][j] == 0 {
				fmt.Printf(". ")
			} else {
				fmt.Printf("%d ", (*board)[i][j])
			}
		}
		fmt.Println()
	}
}

func part1(puzzleInput []line) int {
	board := [boardLen][boardLen]int{}

	for _, line := range puzzleInput {
		if line.x1 == line.x2 {
			drawVerticalLine(&board, &line)
		} else if line.y1 == line.y2 {
			drawHorizontalLine(&board, &line)
		}
	}

	return calculateOverlap(&board)
}

func part2(puzzleInput []line) int {
	board := [boardLen][boardLen]int{}

	for _, line := range puzzleInput {
		if line.x1 == line.x2 {
			drawVerticalLine(&board, &line)
		} else if line.y1 == line.y2 {
			drawHorizontalLine(&board, &line)
		} else {
			drawDiagonalLine(&board, &line)
		}
	}

	return calculateOverlap(&board)
}


func drawVerticalLine(board *[boardLen][boardLen]int, l *line) {
	if l.y1 > l.y2 {
		l.y2, l.y1 = l.y1, l.y2
	}

	for i := l.y1; i <= l.y2; i++ {
		(*board)[l.x1][i]++
	}
}

func drawHorizontalLine(board *[boardLen][boardLen]int, l *line) {
	if l.x1 > l.x2 {
		l.x2, l.x1 = l.x1, l.x2
	}

	for i := l.x1; i <= l.x2; i++ {
		(*board)[i][l.y1]++
	}

}

func drawDiagonalLine(board *[boardLen][boardLen]int, l *line) {
	if l.x1 > l.x2 {
		l.y2, l.y1 = l.y1, l.y2
		l.x2, l.x1 = l.x1, l.x2
	}

	if l.y1 > l.y2 {
		for i,j := l.x1, l.y1; i <= l.x2 && j >= l.y2; i, j = i+1, j-1 {
			(*board)[i][j]++
		}
	} else {
		for i,j := l.x1, l.y1; i <= l.x2 && j <= l.y2; i, j = i+1, j+1 {
			(*board)[i][j]++
		}
	}

}


func calculateOverlap(board *[boardLen][boardLen]int) int {
	result := 0

	for i := 0; i < len((*board)); i++ {
		for j := 0; j < len((*board)[i]); j++ {
			if (*board)[i][j] > 1 {
				result++
			}
		}
	}

	return result
}


func readPuzzleInput(filePath string) []line {
	file, err := os.Open(filePath)

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	arr := []line{}

	for scanner.Scan() {
		lineString := scanner.Text()
		line := formatLineString(lineString)
		arr = append(arr, line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	return arr
}

func formatLineString(lineString string) line {
	lineCoordinates := splitByArrow(lineString)
	start := splitByComma(lineCoordinates[0])
	end := splitByComma(lineCoordinates[1])
	return line{
		start[0],
		start[1],
		end[0],
		end[1],
	}
}

func splitByArrow(lineString string) []string {
	return strings.Split(lineString, " -> ")
}

func splitByComma(coordinateStr string) [coordinateLen]int {
	coordinate := strings.Split(coordinateStr, ",")
	coordinateArr := [coordinateLen]int{}

	for i := 0; i < coordinateLen; i++ {
		coordinateVal, _ := strconv.Atoi(coordinate[i])
		coordinateArr[i] = coordinateVal
	}

	return coordinateArr
}
