package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	moduleMasses := readData("tasks/day1/input.txt")

	result1 := sumCalculation(moduleMasses, roughCalc)
	result2 := sumCalculation(moduleMasses, preciseCalc)

	fmt.Printf("Rough result: %#v \n", result1)
	fmt.Printf("Precise result: %#v \n", result2)
}

func sumCalculation(moduleMasses []int, calculator func(mass int) int) int {
	massSum := 0
	for _, mass := range moduleMasses {
		massSum += calculator(mass)
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

func roughCalc(mass int) int {
	return mass/3 - 2
}

func preciseCalc(mass int) int {
	fuel := mass/3 - 2

	if fuel <= 0 {
		return 0
	}

	return fuel + preciseCalc(fuel)
}
