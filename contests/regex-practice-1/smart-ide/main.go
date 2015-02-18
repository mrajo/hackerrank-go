package main

import (
	"fmt"
	"bufio"
	"os"
	"regexp"
	"strings"
)

func ReadInputIntoString() string {
	scanner := bufio.NewScanner(os.Stdin)

	var lines string;

	for scanner.Scan() {
		lines += scanner.Text() + "\n";
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Reading standard input:", err);
	}

	return lines;
}

func ScanComment(reS *regexp.Regexp, reB *regexp.Regexp, target string, startAt int) ([]int, []int) {
	mS := reS.FindStringIndex(target[startAt:]);
	if mS != nil {
		mS[0] += startAt;
		mS[1] += startAt;
	}

	mB := reB.FindStringIndex(target[startAt:]);
	if mB != nil {
		mB[0] += startAt;
		mB[1] += startAt;
	}

	return mS, mB;
}

func TrimInteriorNewLineSpaces(s string) string {
	re := regexp.MustCompile(`\n\s+`);
	return re.ReplaceAllString(s, "\n");
}

func main() {
	strLines := ReadInputIntoString();
	searchFrom := 0;
	reSingle := regexp.MustCompile(`//.*`);
	reBlock := regexp.MustCompile(`(?s:/\*.+?\*/)`);
	output := "";
	matchSingle, matchBlock := ScanComment(reSingle, reBlock, strLines, searchFrom);

	for matchSingle != nil || matchBlock != nil {
		if matchSingle == nil || (matchBlock != nil && matchBlock[0] < matchSingle[0]) {
			output += strings.TrimSpace(TrimInteriorNewLineSpaces(strLines[matchBlock[0]:matchBlock[1]])) + "\n";
			searchFrom = matchBlock[1];
			matchSingle, matchBlock = ScanComment(reSingle, reBlock, strLines, searchFrom);
		} else if matchBlock == nil || (matchSingle != nil && matchSingle[0] < matchBlock[0]) {
			output += strings.TrimSpace(strLines[matchSingle[0]:matchSingle[1]]) + "\n";
			searchFrom = matchSingle[1];
			matchSingle, matchBlock = ScanComment(reSingle, reBlock, strLines, searchFrom);
		}
	}

	fmt.Println(strings.TrimSpace(output));
}