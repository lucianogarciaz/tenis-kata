package match

type MatchRepository interface {
	Save(match *Match)
	Get(matchId int) *Match
}
