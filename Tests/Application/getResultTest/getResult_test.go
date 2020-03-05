package getResultTest

import (
	. "refactoring-kata-tenis/tenniskata/Application/getResult"
	. "refactoring-kata-tenis/tenniskata/Tests/Domain/matchTest"
	. "refactoring-kata-tenis/tenniskata/Tests/Domain/playerTest"
	. "refactoring-kata-tenis/tenniskata/Tests/Infraestructure"
	"testing"
)

type testSample struct {
	Player1Score  int
	Player2Score  int
	ExpectedScore string
	Player1Name   string
	Player2Name   string
}

func Test_Score_Should_Be_Player_One_Winner(t *testing.T) {
	data := setTest()
	for index, element := range data {
		matchId := 2
		player1 := NewPlayerMother().Create(element.Player1Name, element.Player1Score)
		player2 := NewPlayerMother().Create(element.Player2Name, element.Player2Score)
		matchMock := NewMatchMother().Create(matchId, player1, player2)
		matchMockRepository := NewServiceMockInfraestructure(matchMock)
		service := GetPoint(matchMockRepository)
		score := service.Execute(matchMock.GetId())

		if element.ExpectedScore != score {
			t.Error("The result should be: ", element.ExpectedScore, " instead ", score, index)
		}
	}

}

func setTest() []testSample {
	return []testSample{
		{0, 0, "Love-All", "", ""},
		{1, 1, "Fifteen-All", "", ""},
		{2, 2, "Thirty-All", "", ""},
		{3, 3, "Deuce", "", ""},
		{4, 4, "Deuce", "", ""},

		{1, 0, "Fifteen-Love", "", ""},
		{0, 1, "Love-Fifteen", "", ""},
		{2, 0, "Thirty-Love", "", ""},
		{0, 2, "Love-Thirty", "", ""},
		{3, 0, "Forty-Love", "", ""},
		{0, 3, "Love-Forty", "", ""},
		{4, 0, "Win for Player 1", "Player 1", ""},
		{0, 4, "Win for Fernando", "", "Fernando"},
		{2, 1, "Thirty-Fifteen", "", ""},
		{1, 2, "Fifteen-Thirty", "", ""},
		{3, 1, "Forty-Fifteen", "", ""},
		{1, 3, "Fifteen-Forty", "", ""},
		{4, 1, "Win for Luciano", "Luciano", ""},
		{1, 4, "Win for Fernando", "", "Fernando"},

		{3, 2, "Forty-Thirty", "", ""},
		{2, 3, "Thirty-Forty", "", ""},
		{4, 2, "Win for Luciano", "Luciano", ""},
		{2, 4, "Win for Fernando", "", "Fernando"},

		{4, 3, "Advantage Luciano", "Luciano", ""},
		{3, 4, "Advantage Fernando", "", "Fernando"},
		{5, 4, "Advantage Luciano", "Luciano", ""},
		{4, 5, "Advantage Fernando", "", "Fernando"},
		{15, 14, "Advantage Luciano", "Luciano", ""},
		{14, 15, "Advantage Fernando", "", "Fernando"},

		{6, 4, "Win for Luciano", "Luciano", ""},
		{4, 6, "Win for Fernando", "", "Fernando"},
		{16, 14, "Win for Luciano", "Luciano", ""},
		{14, 16, "Win for Fernando", "", "Fernando"},
		{14, 16, "Win for Fernando", "", "Fernando"},
	}
}
