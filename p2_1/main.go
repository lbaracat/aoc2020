package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type passwordData struct {
	minPolicy    int
	maxPolicy    int
	stringPolicy string
	password     string
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Error opening file. %v", err)
	}
	passwordDB, err := getPasswdData(file)
	if err != nil {
		log.Fatalf("Error parsing passwd file. %v", err)
	}

	var validPasswords int = 0

	for _, v := range passwordDB {
		if strings.Contains(v.password, v.stringPolicy) {
			stringPolicyCount := strings.Count(v.password, v.stringPolicy)

			if (stringPolicyCount >= v.minPolicy) && (stringPolicyCount <= v.maxPolicy) {
				validPasswords++
			} else {
				// fmt.Printf("%v %v %v %v\n", v.minPolicy, v.maxPolicy, v.stringPolicy, v.password)
			}

		}
	}

	fmt.Printf("Valid passwords for p2_1: %d\n", validPasswords)

}

func getPasswdData(r io.Reader) ([]passwordData, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)
	var lineParsed passwordData
	var result []passwordData
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ":")
		lineParsed.password = strings.TrimSpace(line[1])

		policy := strings.Split(line[0], " ")
		lineParsed.stringPolicy = policy[1]

		minMax := strings.Split(policy[0], "-")
		min, err := strconv.Atoi(minMax[0])
		if err != nil {
			log.Fatalf("Error parsing min policy on line %v. %v", line, err)
		}
		lineParsed.minPolicy = min

		max, err := strconv.Atoi(minMax[1])
		if err != nil {
			log.Fatalf("Error parsing max policy on line %v. %v", line, err)
		}
		lineParsed.maxPolicy = max

		result = append(result, lineParsed)
	}
	return result, scanner.Err()
}
