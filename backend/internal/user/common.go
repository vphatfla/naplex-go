package user

// User struct used for http request payload and response decode/encode ONLY
type User struct {
	ID        int32  `json:"user_id"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Picture   string `json:"picture"`
}
