package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type numRange struct {
	start int
	end   int
}

func main() {
	contents := readFileToCSV("day2/input.txt")
	r := parseRangesFromRecordStrings(contents)
	fmt.Printf("isAllPatterned: %d", isAllPatterned(r))

	//fmt.Printf("isPatterned: %t\n", part2(111111111111))
}

func readFileToCSV(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	record, _ := reader.Read()
	return record
}

func parseRangesFromRecordStrings(record []string) []numRange {
	var numRanges []numRange

	for _, field := range record {
		parts := strings.Split(field, "-")
		if len(parts) == 2 {
			startInt, _ := strconv.Atoi(parts[0])
			endInt, _ := strconv.Atoi(parts[1])

			numRanges = append(numRanges, numRange{
				start: startInt,
				end:   endInt,
			})
		}

	}

	return numRanges
}

func isAllPatterned(numRanges []numRange) int {
	totalPatterned := 0
	for _, nr := range numRanges {
		for i := nr.start; i <= nr.end; i++ {
			if part2(i) {
				totalPatterned += i
			}
		}
	}

	return totalPatterned
}

func part1(num int) bool {
	s := strconv.Itoa(num)

	if len(s)%2 != 0 {
		return false
	}
	mid := len(s) / 2
	left := s[:mid]
	right := s[mid:]
	return left == right
}

func part2(num int) bool {
	s := strconv.Itoa(num)

	for sectionSize := range len(s) {
		if sectionSize < 1 || len(s)%sectionSize != 0 {
			continue
		}

		var parts []string
		tempS := s
		for len(tempS) > 0 {
			parts = append(parts, tempS[len(tempS)-sectionSize:])
			tempS = tempS[:len(tempS)-sectionSize]
		}

		if allEqual(parts) {
			return true
		}
	}

	return false
}

func allEqual(parts []string) bool {
	if len(parts) == 1 {
		return false
	}

	newParts := slices.Compact(parts)
	return len(newParts) == 1
}
