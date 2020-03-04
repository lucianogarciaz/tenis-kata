package player

type Player struct {
	name string
	score int
}
// Named constructor
func NewPlayer(name string, score int) *Player {
	return &Player{name: name, score: score}
}
func (player *Player) WonPoint() {
	player.score += 1
}
func (player *Player) GetPlayerName() string {
	return player.name
}
func (player *Player) GetPoints() int {
	return player.score
}

