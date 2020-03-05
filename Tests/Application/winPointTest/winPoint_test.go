package winPointTest

import (
	. "refactoring-kata-tenis/tenniskata/Application/winPoint"
	. "refactoring-kata-tenis/tenniskata/Tests/Domain/matchTest"
	. "refactoring-kata-tenis/tenniskata/Tests/Domain/playerTest"
	. "refactoring-kata-tenis/tenniskata/Tests/Infraestructure"
	"testing"
)

type testSample struct {
	PointsPlayer1 int
	PointsPlayer2 int
}

func setTest() []testSample {
	return []testSample{
		{0, 0},
		{1, 1},
		{2, 2},
		{3, 3},
		{4, 4},

		{1, 0},
		{0, 1},
		{2, 0},
		{0, 2},
		{3, 0},
		{0, 3},
		{4, 0},
		{0, 4},
		{2, 1},
		{1, 2},
		{3, 1},
		{1, 3},
		{4, 1},
		{1, 4},

		{3, 2},
		{2, 3},
		{4, 2},
		{2, 4},

		{4, 3},
		{3, 4},
		{5, 4},
		{4, 5},
		{15, 14},
		{14, 15},

		{6, 4},
		{4, 6},
		{16, 14},
		{14, 16},
		{14, 16},
	}
}

func Test_Check_If_The_Points_Are_Correct_After_Winning_A_Point(t *testing.T) {
	for _, data := range setTest() {
		supposedPointsPlayer1 := data.PointsPlayer1 + 1
		supposedPointsPlayer2 := data.PointsPlayer2 + 1
		player1Mock := NewPlayerMother().Create("player 1", data.PointsPlayer1)
		player2Mock := NewPlayerMother().Create("player 2", data.PointsPlayer2)
		matchMock := NewMatchMother().Create(2, player1Mock, player2Mock)
		matchMockRepository := NewServiceMockInfraestructure(matchMock)
		winPoint := WinPoint(matchMockRepository)
		winPoint.Execute(matchMock.GetId(), "player 1")
		winPoint.Execute(matchMock.GetId(), "player 2")
		if matchMock.GetPlayer1().GetPoints() != supposedPointsPlayer1 {
			t.Error("The number of points should be: ", supposedPointsPlayer1, " instead :", matchMock.GetPlayer1().GetPoints())
		}
		if matchMock.GetPlayer2().GetPoints() != supposedPointsPlayer2 {
			t.Error("The number of points of the seconds player should be:", supposedPointsPlayer2, " instead :", matchMock.GetPlayer2().GetPoints())
		}
	}
}
