package models

type ApiBooleanResponse struct {
	IsScuess bool   `json:"is_success"`
	Msg      string `json:"msg"`
	Token    string `json:"token"`
}
