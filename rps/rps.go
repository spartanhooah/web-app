package rps

import (
	"math/rand"
	"time"
)

const (
	ROCK     = 0
	PAPER    = 1
	SCISSORS = 2
)

var messages = map[string][]string{
	"player":   {"Good for you!", "Way to go!", "You win!"},
	"computer": {"Aw, too bad.", "Try again?", "Oops"},
	"draw":     {"Great minds think alike.", "It's a draw.", "Let's try that again."},
}

type Round struct {
	Message        string `json:"message"`
	ComputerChoice string `json:"computerChoice"`
	RoundResult    string `json:"roundResult"`
}

func PlayRound(playerValue int) Round {
	rand.Seed(time.Now().UnixNano())
	computerValue := rand.Intn(2)
	var result Round

	switch computerValue {
	case ROCK:
		result.ComputerChoice = "Computer chose ROCK"
	case PAPER:
		result.ComputerChoice = "Computer chose PAPER"
	case SCISSORS:
		result.ComputerChoice = "Computer chose SCISSORS"
	default:
	}

	if playerValue == computerValue {
		result.RoundResult = "It's a draw"
		result.Message = messages["draw"][rand.Intn(2)]
	} else if playerValue == (computerValue+1)%3 {
		result.RoundResult = "Player wins!"
		result.Message = messages["player"][rand.Intn(2)]
	} else {
		result.RoundResult = "Computer wins!"
		result.Message = messages["computer"][rand.Intn(2)]
	}

	return result
}
