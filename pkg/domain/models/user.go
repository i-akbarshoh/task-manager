package models

type User struct {
	ID        string `json:"id"`
	FullName  string `json:"full_name"`
	Password  string `json:"password"`
	Avatar    string `json:"avatar"`
	Role      string `json:"role"`
	BirthDate string `json:"birth_date"`
	Phone     string `json:"phone"`
	Position  string `json:"position"`
}
