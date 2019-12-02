package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	moduleMasses := readData("tasks/day1/input.txt")

	result := sumCalculation(moduleMasses)

	fmt.Printf("Result: %#v \n", result)
}

func sumCalculation(moduleMasses []int) int {
	massSum := 0
	for _, mass := range moduleMasses {
		massSum += calculateFuel(mass)
	}

	return massSum
}

func readData(filePath string) []int {
	file, err := os.Open(filePath)

	if err != nil {
		panic("Cannot read the file!")
	}
	defer file.Close()

	outputData := make([]int, 0)

	scanner := bufio.NewScanner(file)


	for scanner.Scan() {
		value, err := strconv.Atoi(scanner.Text())

		if err != nil {
			panic("Cannot read a line.")
		}
		outputData = append(outputData, value)
	}

	return outputData
}

func calculateFuel(mass int) int {
	return mass/3 - 2
}
