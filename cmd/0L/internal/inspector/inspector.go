package inspector

import (
	"encoding/json"

	"github.com/wb/cmd/0L/internal/inspector/model"
)

func Check(object []byte) bool {
	var m model.Model

	err := json.Unmarshal(object, &m)
	if err != nil {
		return false
	}

	return true
}
