package Game

import (
	"fmt"
	"wordGame/internal/data"
)

type Game struct {
	Score int
	Lives int
	Words []data.Word
}

func NewGame() *Game {
	game := Game{
		Score: 0,
		Lives: 5,
		Words: data.GetWrods(),
	}
	return &game
}

func (g *Game) Run() {
	fmt.Println("Welcome to the game")
	fmt.Printf("Guess the hidden word by the given hint\n")
	fmt.Printf(">>>>>Rules<<<<<\n")
	fmt.Printf("1. You have 5 attempts.\n2. Find all words to win.\n3. Have fun and be smart.\n\n")

	for _, word := range g.Words {
		if g.Lives == 0 {
			fmt.Println("You lost!")
			break
		}
		hint := word.Hint
		answer := word.Answer

		fmt.Println("Hint:", hint)

		for g.Lives > 0 {
			var userAnswer string
			fmt.Println("\nEnter your answer")
			fmt.Scan(&userAnswer)

			if userAnswer == answer {
				g.Score++
				fmt.Printf("\nCorrect! Your score for this word: 1\n")
				fmt.Printf("Total score: %v\n", g.Score)
				fmt.Printf("Remaining attempts: %v\n", g.Lives)
				fmt.Println()
				break
			} else {
				fmt.Println("Hint:", hint)
				g.Lives--
				if g.Lives > 0 {
					fmt.Printf("\nWrong answer, Try again!\n")
					fmt.Printf("\nRemaining attempts: %v\n", g.Lives)
				} else {
					fmt.Println("You lost!")
					fmt.Printf("The correct answer was: %v\n", answer)
					return
				}

			}
		}
	}

	if g.Lives > 0 {
		fmt.Printf("Congratulations! You guessed all words. \nFinal score: %v\n", g.Score)
	}
}
