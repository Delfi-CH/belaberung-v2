package model

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Domain string `json:"domain"`
}

type CreateRoomRequest struct {
	Name string `json:"name"`
	Description string `json:"description"`
	Domain string `json:"domain"`
}