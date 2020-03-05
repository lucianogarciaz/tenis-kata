package match

// Its going to be my Aggregate root of my bounded context
import (
	. "refactoring-kata-tenis/tenniskata/Domain/player"
	. "refactoring-kata-tenis/tenniskata/Domain/results"
)

type Match struct {
	matchId int
	player1 *Player
	player2 *Player
}

func NewMatch(matchId int, player1 *Player, player2 *Player) *Match {
	return &Match{
		matchId: matchId,
		player1: player1,
		player2: player2,
	}
}

// The points should be recorded... Something like $eventbus->record
func (match *Match) WonPoint(playerName string) {
	if match.player1.GetPlayerName() == playerName {
		match.player1.WonPoint()
	} else {
		match.player2.WonPoint()
	}
}

func (match *Match) GetScore() string {
	result := NewResultFactory(match.player1, match.player2)
	return result.GetResult()
}
func (match *Match) GetPlayer1() *Player {
	return match.player1
}
func (match *Match) GetPlayer2() *Player {
	return match.player2
}
func (match *Match) GetId() int {
	return match.matchId
}
