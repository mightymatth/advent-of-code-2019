package main

import (
	"fmt"
	"strconv"
)

func main() {
	from := 153517
	to := 630395
	passCandidatesCount := 0
	passBetterCandidatesCount := 0

	for pass := from; pass <= to; pass++ {
		if matchesCriteria(strconv.Itoa(pass)) {
			passCandidatesCount++
		}

		if matchesBetterCriteria(strconv.Itoa(pass)) {
			passBetterCandidatesCount++
		}
	}

	fmt.Printf("Possible passwords count: %v\n", passCandidatesCount)
	fmt.Printf("More possible passwords count: %v\n", passBetterCandidatesCount)
}

func matchesCriteria(pass string) bool {
	return passSixDigit(pass) &&
		passMultipleDigit(pass) &&
		passNeverDecrease(pass)
}

func matchesBetterCriteria(pass string) bool {
	return passSixDigit(pass) &&
		passExactDoubleDigit(pass) &&
		passNeverDecrease(pass)
}

func passSixDigit(pass string) bool {
	return len(pass) == 6
}

func passMultipleDigit(pass string) bool {
	for i := 0; i < len(pass)-1; i++ {
		if pass[i] == pass[i+1] {
			return true
		}
	}

	return false
}

func passExactDoubleDigit(pass string) bool {
	for i := 0; i < len(pass)-1; i++ {
		if pass[i] == pass[i+1] {
			if !digitsEqual(i, i-1, pass) && !digitsEqual(i, i+2, pass) {
				return true
			}
		}
	}

	return false
}

func digitsEqual(index1 int, index2 int, pass string) bool {
	if index1 >= 0 && index1 < len(pass) &&
		index2 >= 0 && index2 < len(pass) &&
		pass[index1] == pass[index2] {
		return true
	}

	return false
}

func passNeverDecrease(pass string) bool {
	for i := 0; i < len(pass)-1; i++ {
		if pass[i] > pass[i+1] {
			return false
		}
	}

	return true
}
