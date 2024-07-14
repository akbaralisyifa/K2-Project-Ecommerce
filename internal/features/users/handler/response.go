package handler

type LoginResponse struct {
	Token string `json:"token"`
}

func ToLoginResponse(token string) LoginResponse {
	return LoginResponse{
		Token: token,
	}
}