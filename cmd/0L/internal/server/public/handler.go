package public

import (
	"fmt"
	"net/http"

	"github.com/wb/cmd/0L/internal/serves"
)

func Distribute(w http.ResponseWriter, r *http.Request) {
	data, err := serves.App.Distribute(r.FormValue("kay"))
	if err != nil {
		fmt.Println("err handler:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = w.Write([]byte(data))
	if err != nil {
		fmt.Println("err handler:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
