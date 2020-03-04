package tenniskata

import (
	"math"
	"testing"
)

type testDataSample struct {
	player1Score  int
	player2Score  int
	expectedScore string
}

var testData = []testDataSample{
	{0, 0, "Love-All"},
	{1, 1, "Fifteen-All"},
	{2, 2, "Thirty-All"},
	{3, 3, "Deuce"},
	{4, 4, "Deuce"},

	{1, 0, "Fifteen-Love"},
	{0, 1, "Love-Fifteen"},
	{2, 0, "Thirty-Love"},
	{0, 2, "Love-Thirty"},
	{3, 0, "Forty-Love"},
	{0, 3, "Love-Forty"},
	{4, 0, "Win for Luciano"},
	{0, 4, "Win for Fernando"},

	{2, 1, "Thirty-Fifteen"},
	{1, 2, "Fifteen-Thirty"},
	{3, 1, "Forty-Fifteen"},
	{1, 3, "Fifteen-Forty"},
	{4, 1, "Win for Luciano"},
	{1, 4, "Win for Fernando"},

	{3, 2, "Forty-Thirty"},
	{2, 3, "Thirty-Forty"},
	{4, 2, "Win for Luciano"},
	{2, 4, "Win for Fernando"},

	{4, 3, "Advantage Luciano"},
	{3, 4, "Advantage Fernando"},
	{5, 4, "Advantage Luciano"},
	{4, 5, "Advantage Fernando"},
	{15, 14, "Advantage Luciano"},
	{14, 15, "Advantage Fernando"},

	{6, 4, "Win for Luciano"},
	{4, 6, "Win for Fernando"},
	{16, 14, "Win for Luciano"},
	{14, 16, "Win for Fernando"},
	{14, 16, "Win for Fernando"},
}

func runSuiteOnGame(t *testing.T, factory func(player1Name string, player2Name string) TennisGame) {
	for _, sample := range testData {
		game := factory("Luciano", "Fernando")
		highestScore := int(math.Max(float64(sample.player1Score), float64(sample.player2Score)))

		for i := 0; i < highestScore; i++ {
			if i < sample.player1Score {
				game.WonPoint("Luciano")
			}
			if i < sample.player2Score {
				game.WonPoint("Fernando")
			}
		}

		result := game.GetScore()
		if sample.expectedScore != result {
			t.Logf("Expected \"%s\" but was \"%s\" for %d:%d", sample.expectedScore, result, sample.player1Score, sample.player2Score)
			t.Fail()
		}
	}
}

func TestMatchTennis(t *testing.T) {
	runSuiteOnGame(t, func(player1Name string, player2Name string) TennisGame { return Match(player1Name, player2Name) })
}
