package authdto

type LoginResponse struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	FullName string `json:"fullName"`
	Token    string `json:"token"`
}