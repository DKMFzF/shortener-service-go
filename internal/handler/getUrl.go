package handler

import (
	"net/http"
	"strings"
)

func GetUrlHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "not GET", http.StatusBadRequest)
		return
	}

	contentType := r.Header.Get("Content-Type")
	if !strings.Contains(contentType, "text/plain") {
		http.Error(w, "Content-Type not text/plain", http.StatusBadRequest)
		return
	}

	id := r.PathValue("id")
	if id == "" {
		http.Error(w, "Not id", http.StatusBadRequest)
		return
	}

	// TODO: write service logic for get url

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("https://practicum.yandex.ru/"))
}
