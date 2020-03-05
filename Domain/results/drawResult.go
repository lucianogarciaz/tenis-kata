package results

import . "refactoring-kata-tenis/tenniskata/Domain/player"

type drawResult struct {
	points int
}

// Named constructor
func NewDrawResult(points int) *drawResult {
	return &drawResult{points: points}
}
func isDrawResult(player1 *Player, player2 *Player) bool {
	if player1.GetPoints() == player2.GetPoints() {
		return true
	} else {
		return false
	}
}
func (result *drawResult) GetResult() string {
	switch result.points {
	case 0:
		return "Love-All"
	case 1:
		return "Fifteen-All"
	case 2:
		return "Thirty-All"
	default:
		return "Deuce"
	}
}
