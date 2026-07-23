package model

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CreateRoomRequest struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Password    *string `json:"password"`
}

type UpdateUserDetailRequest struct {
	RequestType string  `json:"type"`
	Username    *string `json:"username"`
	Biography   *string `json:"biography"`
	Pronouns    *string `json:"pronouns"`
	OldPassword *string `json:"oldPassword"`
	NewPassword *string `json:"newPassword"`
}
