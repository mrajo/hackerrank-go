package main

import (
	"fmt"
	"bufio"
	"os"
	"regexp"
)

func ReadInputIntoArray() []string {
	scanner := bufio.NewScanner(os.Stdin)

	var lines []string;

	for scanner.Scan() {
		lines = append(lines, scanner.Text());
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Reading standard input:", err);
	}

	return lines[1:];
}

func main() {
	lines := ReadInputIntoArray();

	for _, line := range lines {
		matched, _ := regexp.MatchString(`\([+-]?(([1-8]?[0-9])(\.[0-9]+)*|90(\.0+)*), [+-]?(([0-9]?[0-9]|1[0-7][0-9])(\.[0-9]+)*|180(\.0+)*)\)`, line);
		if matched {
			fmt.Println("Valid");
		} else {
			fmt.Println("Invalid");
		}
	}
}