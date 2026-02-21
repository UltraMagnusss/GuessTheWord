package Game

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"wordGame/internal/data"
)

type Game struct {
	Score        int
	Lives        int
	Words        []data.Word
	CurrentIndex int
}

func NewGame(CategroyId int) *Game {
	game := Game{
		Score: 0,
		Lives: 5,
		Words: data.GetWords(CategroyId),
	}
	return &game
}

func (g *Game) GetCurrentHint() string {
	if g.CurrentIndex >= len(g.Words) {
		return ""
	}
	return g.Words[g.CurrentIndex].Hint

}

func (g *Game) CheckAnswer(userAnswer string) bool {
	correctAnswer := g.Words[g.CurrentIndex].Answer
	if g.IsGameOver() {
		return false
	}

	if strings.TrimSpace(strings.ToLower(userAnswer)) == strings.ToLower(correctAnswer) {
		g.Score++

		return true
	} else {
		g.Lives--
		return false
	}
}

func (g *Game) IsGameOver() bool {
	return g.Lives <= 0 || g.CurrentIndex >= len(g.Words)
}

func (g *Game) Run() {
	reader := bufio.NewReader(os.Stdin)

	for !g.IsGameOver() {
		hint := g.GetCurrentHint()
		fmt.Println("Hint: ", hint)
		fmt.Println("Enter your answer")

		userAnswer, _ := reader.ReadString('\n')
		userAnswer = strings.TrimSpace(userAnswer)

		if g.CheckAnswer(userAnswer) {
			fmt.Printf("Correct! Score: %v, Lives: %v\n\n", g.Score, g.Lives)
			g.CurrentIndex++
		} else {
			fmt.Printf("Wrong! Remaining attempts: %v\n\n", g.Lives)
		}
	}

	if g.Lives == 0 {
		fmt.Println("You lost!")
		fmt.Println("The correct answer was: ", g.Words[g.CurrentIndex].Answer)
	} else {
		fmt.Println("Congratulations! You guessed all words!")

	}
}

// func (g *Game) Run() {
// 	fmt.Println("Welcome to the game")
// 	fmt.Printf("Guess the hidden word by the given hint\n")
// 	fmt.Printf(">>>>>Rules<<<<<\n")
// 	fmt.Printf("1. You have 5 attempts.\n2. Find all words to win.\n3. Have fun and be smart.\n\n")

// 	for _, word := range g.Words {
// 		if g.Lives == 0 {
// 			fmt.Println("You lost!")
// 			break
// 		}
// 		hint := word.Hint
// 		answer := word.Answer

// 		fmt.Println("Hint:", hint)

// 		for g.Lives > 0 {
// 			var userAnswer string
// 			fmt.Println("\nEnter your answer")
// 			fmt.Scan(&userAnswer)

// 			if userAnswer == answer {
// 				g.Score++
// 				fmt.Printf("\nCorrect! Your score for this word: 1\n")
// 				fmt.Printf("Total score: %v\n", g.Score)
// 				fmt.Printf("Remaining attempts: %v\n", g.Lives)
// 				fmt.Println()
// 				break
// 			} else {
// 				fmt.Println("Hint:", hint)
// 				g.Lives--
// 				if g.Lives > 0 {
// 					fmt.Printf("\nWrong answer, Try again!\n")
// 					fmt.Printf("\nRemaining attempts: %v\n", g.Lives)
// 				} else {
// 					fmt.Println("You lost!")
// 					fmt.Printf("The correct answer was: %v\n", answer)
// 					return
// 				}

// 			}
// 		}
// 	}

// 	if g.Lives > 0 {
// 		fmt.Printf("Congratulations! You guessed all words. \nFinal score: %v\n", g.Score)
// 	}
// }
