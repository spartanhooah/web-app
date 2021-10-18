package rps

import (
	"math/rand"
	"time"
)

const (
	ROCK          = 0
	PAPER         = 1
	SCISSORS      = 2
	PLAYER_WINS   = 1
	COMPUTER_WINS = 2
	DRAW          = 3
)

func PlayRound(playerValue int) (int, string, string) {
	rand.Seed(time.Now().UnixNano())
	computerValue := rand.Intn(2)
	computerChoice := ""

	switch computerValue {
	case ROCK:
		computerChoice = "Computer chose ROCK"
	case PAPER:
		computerChoice = "Computer chose PAPER"
	case SCISSORS:
		computerChoice = "Computer chose SCISSORS"
	default:
	}

	roundResult := ""
	winner := 0
	if playerValue == computerValue {
		roundResult = "It's a draw"
		winner = DRAW
	} else if playerValue == (computerValue+1)%3 {
		roundResult = "Player wins!"
		winner = PLAYER_WINS
	} else {
		roundResult = "Computer wins!"
		winner = COMPUTER_WINS
	}

	return winner, computerChoice, roundResult
}
