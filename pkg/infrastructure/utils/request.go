package utils

type LoginRequestModel struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type UpdatePasswordRequestModel struct {
	Phone       string `json:"phone"`
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}
