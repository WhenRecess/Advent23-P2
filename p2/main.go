package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {

	const (
		redLimit   = 12
		greenLimit = 13
		blueLimit  = 14
	)

	f, err := os.Open("game_records.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	redReg, err := regexp.Compile("\\d* red")
	greenReg, err := regexp.Compile("\\d* green")
	blueReg, err := regexp.Compile("\\d* blue")

	colorRegM := map[string]*regexp.Regexp{
		"blue":  blueReg,
		"green": greenReg,
		"red":   redReg,
	}

	colorNumM := make(map[string][]int)
	count := 0

	var possibleLines []int
	// regexes := []*regexp.Regexp{redReg, greenReg, blueReg}
	for scanner.Scan() {
		line := scanner.Text()
		count++
		for color, r := range colorRegM {
			matches := r.FindAllString((line), -1)
			// fmt.Printf("For color %v: %v\n", color, matches)

			re := regexp.MustCompile("\\d*")
			var colorNum []int
			for _, match := range matches {
				digit, _ := strconv.Atoi(re.FindString(string(match)))
				colorNum = append(colorNum, digit)
			}

			// fmt.Printf("For color %v: %v\n", color, digits)
			colorNumM[color] = colorNum
			// fmt.Printf("%s", matches)
		}
		// fmt.Printf("%v", colorNumM)
		lineIsImpossible := false
		for _, number := range colorNumM["red"] {
			if number > redLimit {
				lineIsImpossible = true
			}
		}
		for _, number := range colorNumM["blue"] {
			if number > blueLimit {
				lineIsImpossible = true
			}
		}
		for _, number := range colorNumM["green"] {
			if number > greenLimit {
				lineIsImpossible = true
			}
		}
		if !lineIsImpossible {
			fmt.Printf("The line is impossible for %v\n", count)
			possibleLines = append(possibleLines, count)
		}
	}
	sum := 0
	for _, possibleLine := range possibleLines {
		sum += possibleLine
	}
	fmt.Print(sum)
}
