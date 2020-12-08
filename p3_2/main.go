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
	treeMap, err := getTreeMap(file)
	if err != nil {
		log.Fatalf("Error getting tree map. %v", err)
	}

	var treeFound int = 0

	mapPosition := 0
	modBase := len(treeMap[0])
	for i := 0; i < len(treeMap); i++ {
		if strings.Compare(string(treeMap[i][mapPosition%modBase]), "#") == 0 {
			treeFound++
		}
		// fmt.Printf("Line: %d - Pos: %d - PosMOD: %d - byte: %v\n", i, mapPosition, mapPosition%modBase, treeMap[i][mapPosition%modBase])
		mapPosition = mapPosition + 1
		/*
			The slopes required was
			Right1 Down1
			Right3 Down1
			Right5 Down1
			Right7 Down1
			Right1 Down2

			I hardcoded each one and point the results. Cheated, I know.
		*/
	}

	fmt.Printf("Trees found for p3_2: %d\n", treeFound)
}

func getTreeMap(r io.Reader) ([]string, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)
	var result []string
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}
	return result, scanner.Err()
}
