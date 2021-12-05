package main

import (
	"fmt"
	"bufio"
	"os"
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

func part1(puzzleInput []string) int64 {
	bits := makeBitArr (puzzleInput)
	gamma, epsilon := calculateGamEp(bits)

	return calcPowerConsumption(gamma, epsilon)
}

func makeBitArr (puzzleInput []string) [12][2]int {
	bits := [12][2]int{}

	for _, val := range puzzleInput {
		for binaryIndex, char := range val {
			digit, _ := strconv.Atoi(string(char))
			bits[binaryIndex][digit]++
		}
	}

	return bits
}

func calculateGamEp(bits [12][2]int) (string, string) {
	gamma := ""
	epsilon := ""

	for i := 0; i < len(bits); i++ {
		if bits[i][0] > bits[i][1] {
			gamma += "0"
			epsilon += "1"
		} else {
			gamma += "1"
			epsilon += "0"
		}
	}

	return gamma, epsilon
}

func calcPowerConsumption(gamma string, epsilon string) int64 {
	valOfGamma := convertStrToInt(gamma)
	valOfEpsilon := convertStrToInt(epsilon)

	return valOfGamma * valOfEpsilon
}

func convertStrToInt(str string) int64 {
	if val, err := strconv.ParseInt(str, 2, 64); err != nil {
		panic(err)
	} else {
		return val
	}
}

func part2 (puzzleInput []string) int64 {
	oxygenRate := calcOxCo2Rate(puzzleInput, true)
	co2Rate := calcOxCo2Rate(puzzleInput, false)

	return oxygenRate * co2Rate
}

func calcOxCo2Rate (puzzleInput []string, isOxygen bool) int64 {
	currentArr := puzzleInput
	currentDigit := 0

	for len(currentArr) != 1 {
		result := compareZerosAndOnes(currentArr, currentDigit)
		currentArr = fixArray(currentArr, result, isOxygen, currentDigit)
		currentDigit++
	}

	return convertStrToInt(currentArr[0])
}

func compareZerosAndOnes (arr []string, currentDigit int) int {
	countZeros := 0
	countOnes := 0

	for i := 0; i < len(arr); i++ {
		digit := convertByteToInt(arr[i][currentDigit])

		if digit == 0 {
			countZeros++
		} else {
			countOnes++
		}
	}

	if countZeros > countOnes {
		return -1
	} else if countZeros == countOnes {
		return 0
	}

	return 1
}

func convertByteToInt(b byte) int64 {
	return convertStrToInt(string(b))
}

func fixArray (currentArr []string, resultOfCompare int, isOxygen bool, currentDigit int) []string {
	newArray := []string{}

	for i := 0; i < len(currentArr); i++ {
		digit := convertByteToInt(currentArr[i][currentDigit])

		if resultOfCompare == 1 || resultOfCompare == 0 {
			if (isOxygen && digit == 1) || (!isOxygen && digit == 0) {
				newArray = append(newArray, currentArr[i])
			}
		} else {
			if (isOxygen && digit == 0) || (!isOxygen && digit == 1) {
				newArray = append(newArray, currentArr[i])
			}
		}
	}

	return newArray
}
