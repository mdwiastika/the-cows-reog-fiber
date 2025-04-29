package dto

type UserData struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	GoogleID string `json:"google_id"`
}
