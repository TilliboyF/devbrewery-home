package handler

import (
	"log/slog"
	"net/http"
)

func MakeHandler(h func(http.ResponseWriter, *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); nil != err {
			slog.Error(err.Error())
		}
	}
}

func HxRedirect(w http.ResponseWriter, to string) {
	w.Header().Set("HX-REDIRECT", to)
	w.WriteHeader(http.StatusSeeOther)
}
