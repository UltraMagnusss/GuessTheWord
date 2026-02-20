package main

import (
	Game "wordGame/internal/game"
)

func main() {
	game := Game.NewGame()
	game.Run()
}

// подключить БД вместо структуры
// Вынести логику в Веб-запросы
