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

type passport struct {
	byr    string
	iyr    string
	eyr    string
	hgt    string
	hcl    string
	ecl    string
	pid    string
	cid    string
	fields int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Error opening file. %v", err)
	}

	passportData, err := getPassportData(file)
	if err != nil {
		log.Fatalf("Error getting passport data. %v", err)
	}

	var passportList []passport

	for _, passportLine := range passportData {
		var passportParsed passport
		passportParsed.fields = 0
		passportFields := strings.Split(passportLine, " ")
		for _, passportField := range passportFields {
			passportFieldSplited := strings.Split(passportField, ":")
			passportParsed.fields++
			switch passportFieldSplited[0] {
			case "byr":
				passportParsed.byr = passportFieldSplited[1]
			case "iyr":
				passportParsed.iyr = passportFieldSplited[1]
			case "eyr":
				passportParsed.eyr = passportFieldSplited[1]
			case "hgt":
				passportParsed.hgt = passportFieldSplited[1]
			case "hcl":
				passportParsed.hcl = passportFieldSplited[1]
			case "ecl":
				passportParsed.ecl = passportFieldSplited[1]
			case "pid":
				passportParsed.pid = passportFieldSplited[1]
			case "cid":
				passportParsed.cid = passportFieldSplited[1]
			}
		}
		passportList = append(passportList, passportParsed)
	}

	var validPassports int = 0

	for _, value := range passportList {
		if (value.fields == 8) || (value.fields == 7 && len(value.cid) == 0) {
			if validateYear(value.byr, 1920, 2002) &&
				validateYear(value.iyr, 2010, 2020) &&
				validateYear(value.eyr, 2020, 2030) &&
				validatePID(value.pid) &&
				validateEyeColor(value.ecl) &&
				validateHairColor(value.hcl) &&
				validateHeight(value.hgt) {
				validPassports++
				continue
			}

			/* 			// Uncomment to show "valid" passports with invalid fields
			   			fmt.Printf("byr: %v | iyr: %v | eyr: %v | hgt: %v | hcl: %v | ecl: %v | pid: %v | cid: %v | Fields: %v\n",
			   				value.byr,
			   				value.iyr,
			   				value.eyr,
			   				value.hgt,
			   				value.hcl,
			   				value.ecl,
			   				value.pid,
			   				value.cid,
			   				value.fields) */
		}

		/* 		// Uncomment to show invalid passports
		   		fmt.Printf("byr: %v | iyr: %v | eyr: %v | hgt: %v | hcl: %v | ecl: %v | pid: %v | cid: %v | Fields: %v\n",
		   			value.byr,
		   			value.iyr,
		   			value.eyr,
		   			value.hgt,
		   			value.hcl,
		   			value.ecl,
		   			value.pid,
		   			value.cid,
		   			value.fields) */

	}

	fmt.Printf("Valid passports from p4_2: %v\n", validPassports)
}

func getPassportData(r io.Reader) ([]string, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)
	var fullLine string
	var result []string
	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			fullLine = strings.TrimPrefix(fullLine, " ")
			result = append(result, fullLine)
			fullLine = ""
			continue
		}
		fullLine = fullLine + " " + scanner.Text()
	}
	// last line insert - Needed because last scanner.Text() maybe not an empty line.
	fullLine = strings.TrimPrefix(fullLine, " ")
	result = append(result, fullLine)
	return result, scanner.Err()
}

func validateYear(year string, minRange int, maxRange int) bool {
	if len(year) != 4 {
		return false
	}
	intYear, err := strconv.Atoi(year)
	if err != nil {
		return false
	}
	if intYear < minRange || intYear > maxRange {
		return false
	}
	return true
}

func validatePID(pid string) bool {
	if len(pid) != 9 {
		return false
	}
	if _, err := strconv.Atoi(pid); err != nil {
		return false
	}
	return true
}

func validateEyeColor(color string) bool {
	switch color {
	case "amb":
		return true
	case "blu":
		return true
	case "brn":
		return true
	case "gry":
		return true
	case "grn":
		return true
	case "hzl":
		return true
	case "oth":
		return true
	default:
		return false
	}
}

func validateHairColor(color string) bool {
	if len(color) != 7 {
		return false
	}
	if strings.Compare(string(color[0]), "#") != 0 {
		return false
	}
	if _, err := strconv.ParseInt(string(color[1:6]), 16, 0); err != nil {
		return false
	}

	return true
}

func validateHeight(height string) bool {
	measureUnit := height[len(height)-2 : len(height)]
	if strings.Compare(measureUnit, "cm") != 0 &&
		strings.Compare(measureUnit, "in") != 0 {
		return false
	}

	intHeight, err := strconv.Atoi(string(height[:len(height)-2]))
	if err != nil {
		return false
	}
	switch measureUnit {
	case "cm":
		if intHeight < 150 || intHeight > 193 {
			return false
		}
	case "in":
		if intHeight < 59 || intHeight > 76 {
			return false
		}
	}

	return true
}
