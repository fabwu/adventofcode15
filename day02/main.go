package main

import (
	"fmt"
	"path/filepath"
	"bufio"
	"os"
	"strings"
	"strconv"
)

type dimension struct {
	length,
	width,
	height int
}

func (d dimension) WrappingPaperArea() int {
	areas := []int{d.length * d.width, d.width * d.height, d.height * d.length}

	area := 2 * areas[0] + 2 * areas[1] + 2 * areas[2]

	area += findMin(areas)

	return area
}

func (d dimension) RibbonLength() int {
	var length int
	lengths := []int{d.length + d.width, d.width + d.height, d.height + d.length}

	length = 2 * findMin(lengths)

	length += d.length * d.width * d.height

	return length
}

func main() {
	absPath, _ := filepath.Abs("day02a/input.txt")
	f, err := os.Open(absPath)

	checkError(err)

	scanner := bufio.NewScanner(f)

	var wrappingPaperArea int
	var ribbonLength int

	for scanner.Scan() {
		line := scanner.Text()
		numbers := strings.Split(line, "x")
		d := dimension{
			length:convertToInt(numbers[0]),
			width: convertToInt(numbers[1]),
			height: convertToInt(numbers[2])}

		wrappingPaperArea += d.WrappingPaperArea()
		ribbonLength += d.RibbonLength()
	}

	checkError(scanner.Err())

	fmt.Printf("Wrapping Paper: %v\n", wrappingPaperArea)
	fmt.Printf("Ribbon: %v", ribbonLength)
}

func convertToInt(s string) int {
	i, err := strconv.Atoi(s)
	checkError(err)
	return i
}

func findMin(numbers []int) int {
	min := numbers[0]
	for _, number := range numbers {
		if number < min {
			min = number
		}
	}
	return min
}

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}
