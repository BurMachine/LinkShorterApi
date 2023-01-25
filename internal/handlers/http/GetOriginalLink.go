package httpHandlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"regexp"

	"burmachine/LinkGenerator/internal/models"
)

func (s *HttpHandlers) GetOriginalUrl(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	if r.Method != http.MethodGet {
		println("method get error")
	}

	shortLink := r.URL.Query().Get("short")
	if shortLink == "" {
		log.Println("Invalid link received(empty)")
		http.Error(w, errors.New("empty link").Error(), http.StatusBadRequest)
		return
	}

	validId := regexp.MustCompile("^[a-zA-Z0-9_]+$")
	if len(shortLink) != 10 && validId.MatchString(shortLink) {
		log.Println("Invalid link received: ", shortLink)
		http.Error(w, errors.New("invalid link").Error(), http.StatusBadRequest)
		return
	}
	originalLink, err := (*s.Storage).GetFullLink(shortLink)
	if err != nil {
		log.Println("Original link getting error: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	originalResponse := models.RequestBody{originalLink}
	response, err := json.Marshal(originalResponse)
	if err != nil {
		log.Println("Original link marshalling error: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(response)
}
