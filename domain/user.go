package domain

type User struct {
	UserName string `json:"username_name"`
	TenantId int    `json:"tenant_id"`
}

type UserGreetingMessage struct {
	GreetingMessage string `json:"greeting_message"`
	UserDetails     User   `json:"user_details"`
}

type UserRepository interface {
	GetUserDetails() (User, error)
}
