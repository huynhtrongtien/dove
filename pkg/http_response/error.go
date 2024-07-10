package http_response

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Message struct {
	VI string `json:"vi,omitempty"`
	EN string `json:"en,omitempty"`
}

type errorData struct {
	Message *Message `json:"message,omitempty"`
}

func Error(c *gin.Context, code int, err error, msg *Message) {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusBadRequest, &errorData{
			Message: &Message{
				VI: msg.VI,
				EN: msg.EN,
			},
		})
		return
	}

	c.JSON(code, &errorData{
		Message: &Message{
			VI: msg.VI,
			EN: msg.EN,
		},
	})
}