package Infraestructure

import (
	. "refactoring-kata-tenis/tenniskata/Domain/match"
)

type matchMock struct {
	match *Match
}

func NewServiceMockInfraestructure(match *Match) *matchMock {
	return &matchMock{
		match: match,
	}
}
func (matchMock *matchMock) Save(match *Match) {

}
func (matchMock *matchMock) Get(matchId int) *Match {
	return matchMock.match
}
