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

	for i := 0; i < len(expense)-2; i++ {
		for j := i + 1; j < len(expense)-1; j++ {
			for k := j + 1; k < len(expense); k++ {
				if (expense[i] + expense[j] + expense[k]) == 2020 {
					log.Printf("Triple (%d %d %d) - p1_2 answers is %d", expense[i], expense[j], expense[k], expense[i]*expense[j]*expense[k])
				}
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
