package config

// @description Model with information about tasks and their parameters
type Config struct {
	ArenaFarming       bool `json:"arena_farming"`
	ArenaUseEnergyCans bool `json:"arena_use_energy_cans"`
	Travelling         bool `json:"travelling"`
	OpenChests		   bool `json::"open_chests"`
}
