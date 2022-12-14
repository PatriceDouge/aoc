/*
   Elf gives strategy guide on what to play based on what player A plays.
   First column:  A for Rock, B for Paper, and C for Scissors
   Second column: X for Rock, Y for Paper, and Z for Scissors

   Total score is the sum of your scores for each round.
   The score for a single round is the score for the shape you selected
       (1 for Rock, 2 for Paper, and 3 for Scissors)
   + the score for the outcome of the round
       (0 if you lost, 3 if the round was a draw, and 6 if you won
    return total score

    Solution:
    Map each player choice to # (1 for Rock, 2 for Paper, and 3 for Scissors)
    concatenate max from two players to total score

*/

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	filePath := os.Args[1]
	readFile, err := os.Open(filePath)

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	readFile.Close()

	var totalScore int = 0

	scores := make(map[string]int)
	scores["B X"] = 1
	scores["C Y"] = 2
	scores["A Z"] = 3
	scores["A X"] = 4
	scores["B Y"] = 5
	scores["C Z"] = 6
	scores["C X"] = 7
	scores["A Y"] = 8
	scores["B Z"] = 9

	for _, turn := range fileLines {

		fmt.Println("Round total: ", scores[turn])

		totalScore += scores[turn]
	}

	fmt.Println("Total score: ", totalScore)
}

/*B X":1, "C Y":2, "A Z":3, "A X":4,"B Y":5,"C Z":6, "C X":7, "A Y":8, "B Z":9*/
