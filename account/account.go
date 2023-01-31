package account

// @Description Model with static info about game account
type Account struct {
	Owner int    `json:"owner"`
	URL   string `json:"url"`
}
