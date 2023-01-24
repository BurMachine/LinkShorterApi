package httpHandlers

import (
	"burmachine/LinkGenerator/internal/models"
	"encoding/json"
	"log"
	"net/http"
)

func (s *HttpHandlers) GetOriginalUrl(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	if r.Method != http.MethodGet {
		println("method get error")
	}

	shortLink := r.URL.Query().Get("short")
	originalLink, err := (*s.Storage).GetFullLink(shortLink)
	if err != nil {
		log.Println("Original link getting error: ", err)
		return
	}
	originalResponse := models.RequestBody{originalLink}
	response, err := json.Marshal(originalResponse)
	if err != nil {
		log.Println("Original link marshalling error: ", err)
		return
	}
	w.Write(response)
}
