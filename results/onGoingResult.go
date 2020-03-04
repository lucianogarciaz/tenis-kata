package results

type onGoingResult struct {
	score1 int
	score2 int
}
// NamedConstructor
func NewOnGoingResult(score1 int, score2 int) *onGoingResult {
	return &onGoingResult{
		score1: score1,
		score2: score2,
	}
}

func (results *onGoingResult) GetResult() string{
	return namedScore(results.score1) + "-" + namedScore(results.score2)
}
func namedScore(score int) string {
	switch score {
	case 1:
		return "Fifteen"
	case 2:
		return "Thirty"
	case 3:
		return "Forty"
	default:
		return "Love"
	}
}