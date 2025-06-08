package auth

type GoogleUserInfo struct {
	ID        string `json:"id"`
	FirstName string `json:"given_name"`
	LastName  string `json:"family_name"`
	Name      string `json:"name"`
	Picture   string `json:"picture"`
}
