package utilities

import (
	"fmt"

	"github.com/google/uuid"
)

func UUID() {
	id := uuid.New()
	fmt.Println("Generated UUID:")
	fmt.Println(id.String())
}
