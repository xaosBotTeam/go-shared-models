package dbAccountInformation

type DbAccountInformation struct {
	ID           int    `json:"id"`
	GameID       int    `json:"game_id"`
	FriendlyName string `json:"friendly_name"`
	Owner        int    `json:"owner"`
	URL          string `json:"url"`
	EnergyLimit  int    `json:"energy_limit"`
}
