package handler

import (
	"net/http"

	"github.com/TilliboyF/devbrewery-home/view/home"
)

func HandleGetHome(w http.ResponseWriter, r *http.Request) error {
	return home.Index().Render(r.Context(), w)
}
