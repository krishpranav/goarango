package requests

type JWTAuth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (r *JWTAuth) Path() string {
	return "/_open/auth"
}

func (r *JWTAuth) Method() string {
	return "POST"
}
