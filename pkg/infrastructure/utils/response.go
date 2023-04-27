package utils

type RegisterResponseModel struct {
	ID           string `json:"id"`
	AccessToken  string `json:"access_token"`
	ExpiresIn    int64  `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
}

type LoginResponseModel struct {
	ID string `json:"id"`
}
