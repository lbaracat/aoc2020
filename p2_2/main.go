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
	position1Policy int
	position2Policy int
	stringPolicy    string
	password        string
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
			position1Match := strings.Compare(string(v.password[v.position1Policy-1]), v.stringPolicy)
			position2Match := strings.Compare(string(v.password[v.position2Policy-1]), v.stringPolicy)

			if (position1Match == 0 || position2Match == 0) && (position1Match != position2Match) {
				validPasswords++
				// fmt.Printf("Valid: %v %v\n", position1Match, position2Match)
			} else {
				// fmt.Printf("%v %v %v %v\n", v.position1Policy, v.position2Policy, v.stringPolicy, v.password)
				// fmt.Printf("Invalid: %v %v\n", position1Match, position2Match)
			}
		}
	}

	fmt.Printf("Valid passwords for p2_2: %d\n", validPasswords)

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
		lineParsed.position1Policy = min

		max, err := strconv.Atoi(minMax[1])
		if err != nil {
			log.Fatalf("Error parsing max policy on line %v. %v", line, err)
		}
		lineParsed.position2Policy = max

		result = append(result, lineParsed)
	}
	return result, scanner.Err()
}
