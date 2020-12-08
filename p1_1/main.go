package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Error opening file. %v", err)
	}
	expense, err := getExpense(file)
	if err != nil {
		log.Fatalf("Error getting expense. %v", err)
	}

	for i := 0; i < len(expense)-1; i++ {
		for j := i + 1; j < len(expense); j++ {
			if (expense[i] + expense[j]) == 2020 {
				log.Printf("Pairs (%d %d) - p1_1 answers is %d", expense[i], expense[j], expense[i]*expense[j])
			}
		}
	}

}

func getExpense(r io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)
	var expense []int
	for scanner.Scan() {
		if v, err := strconv.Atoi(scanner.Text()); err != nil {
			return nil, err
		} else {
			expense = append(expense, v)
		}
	}
	return expense, scanner.Err()
}
