package results

import (
	"math"
	. "refactoring-kata-tenis/tenniskata/player"
)

type winResult struct {
	player1 *Player
	player2 *Player
}

func NewWinResult(player1 *Player, player2 *Player) *winResult {
	return &winResult{
		player1: player1,
		player2: player2,
	}
}
func (results *winResult) GetResult() string {
	return results.playerIsWinner()
}
func (results *winResult) playerIsWinner() string {
	difference := results.player1.GetPoints() - results.player2.GetPoints()
	if difference >= 0 {
		return "Win for " + results.player1.GetPlayerName()
	} else {
		return "Win for " + results.player2.GetPlayerName()
	}
}

func IsWinResult(player1 *Player, player2 *Player) bool {
	if !atLeastHaveFourPoints(player1.GetPoints(), player2.GetPoints()) {
		return false
	}
	return isGreaterThanOne(player1.GetPoints(), player2.GetPoints())
}

func isGreaterThanOne(score1 int, score2 int) bool {
	difference := score1 - score2
	absoluteDifference := math.Abs(float64(difference))
	if absoluteDifference > 1 {
		return true
	} else {
		return false
	}
}
