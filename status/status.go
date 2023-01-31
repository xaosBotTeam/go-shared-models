package status

type Status struct {
	GameID       int    `json:"game_id"`
	FriendlyName string `json:"friendly_name"`
	EnergyLimit  int    `json:"energy_limit"`
}
