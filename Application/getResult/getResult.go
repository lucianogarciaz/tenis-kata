package getResult

import . "refactoring-kata-tenis/tenniskata/Domain/match"

type getResult struct {
	repository MatchRepository
}

func GetPoint(repository MatchRepository) *getResult {
	return &getResult{repository: repository}
}
func (service *getResult) Execute(matchId int) string {
	match := service.repository.Get(matchId)
	return match.GetScore()
}
