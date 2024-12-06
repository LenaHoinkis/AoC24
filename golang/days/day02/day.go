package main

import (
	"flag"
	"fmt"
	"log"
	"math"

	"github.com/lenahoinkis/AoC24/utils"
)

var inputFile = flag.String("inputFile", "ex.input", "Relative file path to use as input.")

func main() {
	flag.Parse()

	ignoreChars := map[rune]bool{
		' ': true, // Ignore spaces
	}

	reports, err := utils.ReadIntsMatrix(*inputFile, ignoreChars)
	if err != nil {
		log.Fatalf("Error reading columns: %v", err)
	}

	//Part1
	unsafeReports := [][]int{}
	unsafeReportsCounter := 0
	for _, report := range reports {
		distances := []int{}
		//calc distances
		for i := 0; i < len(report)-1; i++ {
			distances = append(distances, (report[i] - report[i+1]))
		}
		//check if safe
		for _, distance := range distances {
			absDistance := int(math.Abs(float64(distance)))

			if absDistance < 0 || absDistance > 3 {
				unsafeReportsCounter++
				break
			}
			if haveSameSign(distance, distances[0]) == false {
				unsafeReportsCounter++
				break
			}
		}
	}

	fmt.Println(len(reports) - unsafeReportsCounter)

	//Part2
	//find all possible combinations
	newSafeReportsCounter := len(unsafeReports)
	fmt.Println(newSafeReportsCounter)
	for _, unsafeReport := range unsafeReports {
		distances := []int{}
		//calc distances
		for i := 0; i < len(unsafeReport)-1; i++ {
			distances = append(distances, (unsafeReport[i] - unsafeReport[i+1]))
		}
		for _, distance := range distances {
			absDistance := int(math.Abs(float64(distance)))

			if absDistance < 0 || absDistance > 3 {
				newSafeReportsCounter--
				break
			}
			if haveSameSign(distance, distances[0]) == false {
				newSafeReportsCounter--
				break
			}
		}
	}
	fmt.Println(newSafeReportsCounter)
	fmt.Println(len(reports) - unsafeReportsCounter + newSafeReportsCounter)

}

func haveSameSign(a, b int) bool {
	// Check if both numbers are either positive or negative
	return (a > 0 && b > 0) || (a < 0 && b < 0)
}
