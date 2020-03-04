package results

import ."refactoring-kata-tenis/tenniskata/player"

func NewResultFactory(player1 *Player, player2 *Player) Results {
	if isDrawResult(player1, player2) {
		return NewDrawResult(player1.GetPoints())
	} else if IsAdvantageResult(player1, player2) {
		return NewAdvantageResult(player1, player2)
	} else if IsWinResult(player1, player2) {
		return NewWinResult(player1, player2)
	} else {
		return NewOnGoingResult(player1.GetPoints(), player2.GetPoints())
	}
}
func atLeastHaveFourPoints(score1 int, score2 int) bool {
	if score1 > 3 || score2 > 3 {
		return true
	} else {
		return false
	}
}
