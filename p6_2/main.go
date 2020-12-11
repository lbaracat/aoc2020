package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
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
		questionsAnsweredYES += len(cdfAnswer)
	}
	fmt.Printf("Sum of all questions answered yes for ALL in a group in p6_2: %v\n", questionsAnsweredYES)
}

func getCDFanswers(r io.Reader) ([]string, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)
	var personAnswers, groupAnswers string
	var tempAnswers []byte
	var newGroup bool = true
	var result []string
	for scanner.Scan() {
		personAnswers = scanner.Text()

		if len(personAnswers) == 0 {
			if len(groupAnswers) != 0 {
				result = append(result, groupAnswers)
			}
			groupAnswers = ""
			newGroup = true
			continue
		}

		if newGroup {
			groupAnswers = personAnswers
			newGroup = false
			continue
		}

		for i := 0; i < len(personAnswers); i++ {
			if strings.Contains(groupAnswers, string(personAnswers[i])) {
				tempAnswers = append(tempAnswers, personAnswers[i])
			}
		}
		groupAnswers = string(tempAnswers)
		tempAnswers = nil

	}
	// last line insert - Needed because last scanner.Text() maybe not an empty line.
	if len(groupAnswers) != 0 {
		result = append(result, groupAnswers)
	}
	return result, scanner.Err()
}
