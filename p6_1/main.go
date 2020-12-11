package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Error opening file. %v", err)
	}

	cdfAnswers, err := getCDFanswers(file)
	if err != nil {
		log.Fatalf("Error getting custom declaration forms. %v", err)
	}

	var questionsAnsweredYES int = 0
	for _, cdfAnswer := range cdfAnswers {
		distinctAnswers, err := getDisctinctAnswers(cdfAnswer)
		if err != nil {
			log.Fatalf("Error finding distinct answers. %v", err)
		}
		questionsAnsweredYES += len(distinctAnswers)
	}
	fmt.Printf("Sum of all questions answered yes from anyone in  p6_1: %v\n", questionsAnsweredYES)
}

func getCDFanswers(r io.Reader) ([]string, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)
	var fullLine string
	var result []string
	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			result = append(result, sortString(fullLine))
			fullLine = ""
			continue
		}
		fullLine += scanner.Text()
	}
	// last line insert - Needed because last scanner.Text() maybe not an empty line.
	result = append(result, sortString(fullLine))
	return result, scanner.Err()
}

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func sortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}

func getDisctinctAnswers(s string) (string, error) {
	var distinct []byte

	distinct = append(distinct, s[0])

	for i := 1; i < len(s); i++ {
		switch strings.Compare(string(distinct[len(distinct)-1]), string(s[i])) {
		case -1:
			distinct = append(distinct, s[i])
		case 1:
			return "", errors.New("The input string must be sorted")
		}
	}
	return string(distinct), nil
}
