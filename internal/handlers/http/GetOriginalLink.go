package httpHandlers

import (
	"encoding/json"
	"log"
	"net/http"

	"burmachine/LinkGenerator/internal/models"
)

func (s *HttpHandlers) GetOriginalUrl(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	if r.Method != http.MethodGet {
		println("method get error")
	}

	shortLink := r.URL.Query().Get("short")
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
