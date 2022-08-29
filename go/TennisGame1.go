package tenniskata

type tennisGame1 struct {
	m_score1    int
	m_score2    int
	player1Name string
	player2Name string
}

func TennisGame1(player1Name string, player2Name string) TennisGame {
	game := &tennisGame1{
		player1Name: player1Name,
		player2Name: player2Name}

	return game
}

func (game *tennisGame1) WonPoint(playerName string) {
	if playerName == "player1" {
		game.m_score1 += 1
	} else {
		game.m_score2 += 1
	}
}

func (game *tennisGame1) GetScore() string {
	if game.m_score1 == game.m_score2 {
		switch game.m_score1 {
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
	if game.m_score1 >= 4 || game.m_score2 >= 4 {
		return stringDeucePoints(game.m_score1, game.m_score2)
	}
	return stringPoints(game.m_score1) + "-" + stringPoints(game.m_score2)
}

func stringDeucePoints(score1, score2 int) string {
	minusResult := score1 - score2
	if minusResult == 1 {
		return "Advantage player1"
	}
	if minusResult == -1 {
		return "Advantage player2"
	}
	if minusResult >= 2 {
		return "Win for player1"
	}
	return "Win for player2"
}

func stringPoints(score int) string {
	switch score {
	case 0:
		return "Love"
	case 1:
		return "Fifteen"
	case 2:
		return "Thirty"
	case 3:
		return "Forty"
	}
	return "unknown"
}
