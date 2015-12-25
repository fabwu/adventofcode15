package main
import (
	"strings"
	"path/filepath"
	"os"
	"bufio"
	"fmt"
)

func main() {
	absPath, _ := filepath.Abs("day05/input.txt")
	f, err := os.Open(absPath)

	checkError(err)

	scanner := bufio.NewScanner(f)

	var oldNiceStrings int
	var newNiceStrings int

	for scanner.Scan() {
		line := scanner.Text()

		if isOldNiceString(line) {
			oldNiceStrings++
		}

		if isNewOldNiceString(line) {
			newNiceStrings++
		}
	}

	checkError(scanner.Err())

	fmt.Printf("Old nice strings: %v\n", oldNiceStrings)
	fmt.Printf("New nice strings: %v\n", newNiceStrings)

}

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func isOldNiceString(str string) bool {
	hasEnoughVowels := countVowels(str) >= 3

	return hasEnoughVowels && hasDoubleLetter(str) && !hasBadString(str)
}

func isNewOldNiceString(str string) bool {
	return hasPair(str) && hasSandwich(str)
}

func countVowels(str string) int {
	var total int
	total += strings.Count(str, "a")
	total += strings.Count(str, "e")
	total += strings.Count(str, "i")
	total += strings.Count(str, "o")
	total += strings.Count(str, "u")
	return total
}

func hasDoubleLetter(str string) bool {
	var hasDoubleLetter bool

	for _, char := range str {
		doubleLetter := string(char) + string(char)

		hasDoubleLetter = strings.Contains(str, doubleLetter)

		if hasDoubleLetter {
			break
		}
	}

	return hasDoubleLetter
}

func hasBadString(str string) bool {
	return strings.Contains(str, "ab") ||
	strings.Contains(str, "cd") ||
	strings.Contains(str, "pq") ||
	strings.Contains(str, "xy")
}

func hasPair(str string) bool {
	var hasPair bool

	for i := 0; i < len(str) - 1; i++ {
		pair := string(str[i]) + string(str[i + 1])
		hasPair = strings.Count(str, pair) >= 2

		if hasPair {
			break
		}
	}

	return hasPair
}

func hasSandwich(str string) bool {
	var hasSandwich bool

	for i := 0; i < len(str) - 2; i++ {
		hasSandwich = str[i] == str[i + 2]

		if hasSandwich {
			break
		}
	}

	return hasSandwich
}