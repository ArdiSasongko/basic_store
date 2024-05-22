package web

type UserRequest struct {
	Name     string `validate:"required" json:"name"`
	Email    string `validate:"required,email" json:"email"`
	Password string `validate:"required" json:"password"`
}

type UserLoginRequest struct {
	Email    string `validate:"required,email" json:"email"`
	Password string `validate:"required" json:"password"`
}

type UserUpdateRequest struct {
	Name string `validate:"required" json:"name"`
}
