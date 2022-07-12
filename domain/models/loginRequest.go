package domain

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type VerifyRequest struct {
	AuthToken string `json:"authtoken"`
}
