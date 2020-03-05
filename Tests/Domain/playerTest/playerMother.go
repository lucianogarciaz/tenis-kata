package playerTest

import . "refactoring-kata-tenis/tenniskata/Domain/player"

type playerMother struct {
}

func NewPlayerMother() *playerMother {
	return &playerMother{}
}

func (playerMother *playerMother) Create(playerName string, score int) *Player {
	return NewPlayer(playerName, score)
}

// It should be implemented with a class WordRandom that returns a random string and
// In a real project it should be stored in a path something like .. Tests/Shared/Domain
func (playerMother *playerMother) Random() *Player {
	// return NewPlayer(WordRandom.random(), NumberRadom.random())
	return NewPlayer("Player1 1", 0)
}
