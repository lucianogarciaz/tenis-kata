package matchTest

import (
	. "refactoring-kata-tenis/tenniskata/Domain/match"
	. "refactoring-kata-tenis/tenniskata/Domain/player"
)

type matchMother struct {
}

func NewMatchMother() *matchMother {
	return &matchMother{}
}
func (matchMother *matchMother) Create(matchId int, player1 *Player, player2 *Player) *Match {
	return NewMatch(matchId, player1, player2)
}
func (matchMother *matchMother) Random() {
	// to implement in the future
}
