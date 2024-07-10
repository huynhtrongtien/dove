package apis

type AuthenticateRequest struct {
	Username string `json:"username" binding:"required" valid:"MaxSize(100)"`
	Password string `json:"password" binding:"required" valid:"MaxSize(100)"`
}

type AuthenticateResponse struct {
	Token string `json:"token"`
}
