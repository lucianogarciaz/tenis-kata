package winPoint

import (
	. "refactoring-kata-tenis/tenniskata/Domain/match"
)

/* Here in a real case I'd have a repository and an event bus to comunicate the changes in my domain.
I'll use an empty struct and a Named constructor because I dont want to other services execute
this service without creating it first. */
type winPoint struct {
	repository MatchRepository
}

func WinPoint(repository MatchRepository) *winPoint {
	return &winPoint{
		repository: repository,
	}
}
func (service *winPoint) Execute(matchId int, playerName string) {
	match := service.repository.Get(matchId)
	match.WonPoint(playerName)
	// here should publish the events of winning a point
	// And if the points should be stored, should make a kind of save points
}
