package apis

type ErrorMessage struct {
	VI string `json:"vi,omitempty"`
	EN string `json:"en,omitempty"`
	JP string `json:"jp,omitempty"`
}

type Error struct {
	Domain  string        `json:"domain,omitempty"`
	Reason  string        `json:"reason,omitempty"`
	Message *ErrorMessage `json:"message,omitempty"`
	Data    interface{}   `json:"data,omitempty"`
}

type Empty struct {
}

type CreateResponse struct {
	UUID string `json:"uuid"`
}
