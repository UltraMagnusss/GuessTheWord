package ui

import "fmt"

func ScanCat() int {
	var choice int
	for {
		fmt.Print("Enter category number (1-5): ")
		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println("Invalid input, try again.")
			fmt.Scanln() // очищаем буфер
			continue
		}
		if choice >= 1 && choice <= 5 {
			break
		} else {
			fmt.Println("Please enter a number between 1 and 5.")
		}
	}
	fmt.Println()
	return choice
}
