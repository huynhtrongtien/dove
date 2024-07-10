package apis

type Register struct {
	DisplayName string `json:"displayname,omitempty"`
	Username    string `json:"username,omitempty"`
	Password    string `json:"password,omitempty"`
}

type User struct {
	UUID        string `json:"uuid,omitempty"`
	DisplayName string `json:"displayname,omitempty"`
	Username    string `json:"username,omitempty"`
}

type SelfProfile struct {
	UUID        string `json:"uuid,omitempty"`
	DisplayName string `json:"displayname,omitempty"`
	Username    string `json:"username,omitempty"`
}

type SelfUpdateProfileRequest struct {
	DisplayName string `json:"displayname,omitempty"`
}

type UpdatePasswordRequest struct {
	Password string `json:"password,omitempty"`
}
