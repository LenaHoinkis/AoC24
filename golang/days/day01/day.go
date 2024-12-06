package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"sort"

	"github.com/lenahoinkis/AoC24/utils"
)

var inputFile = flag.String("inputFile", "ex.input", "Relative file path to use as input.")

func main() {
	flag.Parse()
	col1, col2, err := utils.ReadColumnsOfInt(*inputFile)
	if err != nil {
		log.Fatalf("Error reading columns: %v", err)
	}

	//Part1
	sort.Ints(col1)
	sort.Ints(col2)
	count := 0
	for i := 0; i < len(col1); i++ {
		count += int(math.Abs(float64(col1[i] - col2[i])))
	}
	fmt.Println(count)

	//Part2
	//count the occurrences of each number in the column 2
	counter := make(map[int]int)
	for i := 0; i < len(col2); i++ {
		counter[col2[i]]++
	}

	count = 0
	for i := 0; i < len(col1); i++ {
		count += col1[i] * counter[col1[i]]
	}

	fmt.Println(count)

}
