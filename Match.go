package tenniskata

import (
	. "refactoring-kata-tenis/tenniskata/player"
	. "refactoring-kata-tenis/tenniskata/results"
)

type match struct {
	player1 *Player
	player2 *Player
}

func Match(player1Name string, player2Name string) TennisGame {
	player1 := NewPlayer(player1Name, 0)
	player2 := NewPlayer(player2Name, 0)
	return &match{
		player1: player1,
		player2: player2,
	}
}

func (match *match) WonPoint(playerName string) {
	if playerName == match.player1.GetPlayerName() {
		match.player1.WonPoint()
	} else {
		match.player2.WonPoint()
	}
}

func (match *match) GetScore() string {
	result := NewResultFactory(match.player1, match.player2)
	return result.GetResult()
}
