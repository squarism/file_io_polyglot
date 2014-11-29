package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) int {
	file, err := os.Open(path)
	check(err)
	defer file.Close()

	var telephone = regexp.MustCompile(`\(\d+\)\s\d+-\d+`)
	i := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if telephone.MatchString(scanner.Text()) {
			i += 1
		}
	}
	return i
}

func main() {
	lines := readLines(os.Args[1])
	fmt.Println(lines)
}
