package flyweight

type PlayerFactory struct {
	existingPlayers map[string]Player
}

func NewPlayerFactory() *PlayerFactory {
	return &PlayerFactory{existingPlayers: make(map[string]Player)}
}

func (p *PlayerFactory) GetPlayer(playerType string) Player {
	if player, exists := p.existingPlayers[playerType]; exists {
		return player
	}

	var player Player
	switch playerType {
	case "T":
		player = &Terrorist{
			task:   "Plant a bomb",
			weapon: "AK-47",
		}
	case "CT":
		player = &CounterTerrorist{
			task:   "Defuse bomb",
			weapon: "M4A1",
		}
	}
	p.existingPlayers[playerType] = player
	return player
}
