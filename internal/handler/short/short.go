package short

import (
	"io"
	"net/http"
	"net/url"
	"shortener/internal/service"
	"strings"
)

func ShortHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST", http.StatusBadRequest)
		return
	}

	contentType := r.Header.Get("Content-Type")
	if !strings.Contains(contentType, "text/plain; charset=utf-8") {
		http.Error(w, "Content-Type must be text/plain; charset=utf-8", http.StatusUnsupportedMediaType)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Not URL", http.StatusBadGateway)
		return
	}

	originUrl := string(body)
	if originUrl == "" {
		http.Error(w, "not url", http.StatusBadRequest)
		return
	}

	if _, err := url.ParseRequestURI(originUrl); err != nil {
		http.Error(w, `It's not URL`, http.StatusBadRequest)
		return
	}

	shortCode := service.ResizeUrl(originUrl)
	resStr := "http://localhost:8080/" + shortCode

	w.Header().Set("Content-type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(resStr))
}
