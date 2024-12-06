package utils

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// ReadLinesOfInt reads line and convert the number to int
func ReadLinesOfInt(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var numbers []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		number, _ := strconv.Atoi(line)
		numbers = append(numbers, number)
	}
	return numbers, scanner.Err()
}

// ReadLinesOfInt reads line
func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	return lines, scanner.Err()
}

// ReadInts reads whitespace-separated ints from r. If there's an error, it
// returns the ints successfully read so far as well as the error value.
func ReadInts(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	var result []int
	intString := scanner.Text()
	for _, v := range strings.Split(intString, ",") {
		x, err := strconv.Atoi(v)
		if err != nil {
			return result, err
		}
		result = append(result, x)

	}
	return result, scanner.Err()
}

// ReadColumnsOfInt reads a file with two integers per line separated by whitespace
// and returns two slices: one for the first column and one for the second column.
func ReadColumnsOfInt(path string) ([]int, []int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	var col1, col2 []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Split the line into fields based on whitespace
		fields := strings.Fields(line)
		if len(fields) != 2 {
			return nil, nil, strconv.ErrSyntax // Ensure exactly two columns
		}

		// Convert each field to an integer
		num1, err1 := strconv.Atoi(fields[0])
		num2, err2 := strconv.Atoi(fields[1])
		if err1 != nil || err2 != nil {
			return nil, nil, strconv.ErrSyntax
		}

		// Append to respective columns
		col1 = append(col1, num1)
		col2 = append(col2, num2)
	}
	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	return col1, col2, nil
}

func ReadIntsWithoutSeperator(path string) ([]int, int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, 0, err
	}
	defer file.Close()

	var result []int
	rows := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		rows = 0
		for _, v := range line {
			number, _ := strconv.Atoi(string(v))
			result = append(result, number)
			rows++
		}
	}
	return result, rows, scanner.Err()
}

// ReadIntsMatrix reads a matrix of integers from a file, optionally ignoring specific characters.
func ReadIntsMatrix(path string, ignore map[rune]bool) ([][]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var result [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Remove characters that need to be ignored, but keep spaces
		for char := range ignore {
			if char != ' ' { // Don't remove spaces
				line = strings.ReplaceAll(line, string(char), "")
			}
		}

		// Split the line into separate parts (numbers)
		parts := strings.Fields(line) // This splits by any whitespace

		var intSlice []int
		for _, part := range parts {
			// Convert each part to an integer
			num, err := strconv.Atoi(part)
			if err != nil {
				return nil, err // Error converting string to integer
			}
			intSlice = append(intSlice, num)
		}

		result = append(result, intSlice)
	}
	return result, scanner.Err()
}
