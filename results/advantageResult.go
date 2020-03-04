package results

import (
	"math"
	. "refactoring-kata-tenis/tenniskata/player"
)

type advantageResult struct {
	player1 *Player
	player2 *Player
}

// Named Constructor
func NewAdvantageResult(player1 *Player, player2 *Player) *advantageResult {
	return &advantageResult{
		player1: player1,
		player2: player2,
	}
}

func (results *advantageResult) GetResult() string {
	return results.playerIsInAdvantage()
}
func (results *advantageResult) playerIsInAdvantage() string {
	difference := results.player1.GetPoints() - results.player2.GetPoints()
	if difference >= 0 {
		return "Advantage " + results.player1.GetPlayerName()
	} else {
		return "Advantage " + results.player2.GetPlayerName()
	}
}
func IsAdvantageResult(player1 *Player, player2 *Player) bool {
	if !atLeastHaveFourPoints(player1.GetPoints(), player2.GetPoints()) {
		return false
	}
	return isEqualToOne(player1.GetPoints(), player2.GetPoints())
}

func isEqualToOne(score1 int, score2 int) bool {
	difference := score1 - score2
	absoluteDifference := math.Abs(float64(difference))
	if absoluteDifference == 1 {
		return true
	} else {
		return false
	}
}
