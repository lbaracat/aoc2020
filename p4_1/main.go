package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
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
		if value.fields == 8 {
			validPassports++
			continue
		}

		if value.fields == 7 && len(value.cid) == 0 {
			validPassports++
			continue
		}
		/*
			// Uncomment to show invalid passports
			fmt.Printf("byr: %v | iyr: %v | eyr: %v | hgt: %v | hcl: %v | ecl: %v | pid: %v | cid: %v | Fields: %v\n",
				value.byr,
				value.iyr,
				value.eyr,
				value.hgt,
				value.hcl,
				value.ecl,
				value.pid,
				value.cid,
				value.fields)
		*/
	}

	fmt.Printf("Valid passports from p4_1: %v\n", validPassports)
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
