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

func (g *tennisGame1) WonPoint(playerName string) {
	if playerName == "player1" {
		g.m_score1 += 1
	} else {
		g.m_score2 += 1
	}
}

func (g *tennisGame1) GetScore() string {
	if g.m_score1 == g.m_score2 {
		return g.stringEqualPoints()
	}
	if g.m_score1 >= 4 || g.m_score2 >= 4 {
		return g.stringDeucePoints()
	}
	return g.stringNormalPoints()
}

func (g *tennisGame1) stringEqualPoints() string {
	switch g.m_score1 {
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

func (g *tennisGame1) stringDeucePoints() string {
	minusResult := g.m_score1 - g.m_score2
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

func (g *tennisGame1) stringNormalPoints() string {
	return stringPoints(g.m_score1) + "-" + stringPoints(g.m_score2)
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
