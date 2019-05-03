package proto

type CommandJoin struct {
	GamePIN string `json:"game_pin"`
	UserID  string `json:"user_id"`
	Token   string `json:"token"`
}
