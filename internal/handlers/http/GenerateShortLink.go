package httpHandlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"burmachine/LinkGenerator/internal/models"
	"burmachine/LinkGenerator/internal/service"
)

func (s *HttpHandlers) GenerateShortLink(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	if r.Method != http.MethodPost {
		println("method post error")
	}

	var body models.RequestBody
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Parse body error: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}
	err = json.Unmarshal(bodyBytes, &body)
	if err != nil {
		log.Println("Unmarshalling error in parsing post request's body: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	shortLink, err := service.GenerateLink(body.OriginalUrl)
	if err != nil {
		log.Println("Link generation error: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = (*s.Storage).AddShortLink(body.OriginalUrl, shortLink)
	if err != nil {
		log.Println("Link add error: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := models.ResponseBody{ShortUrl: shortLink}
	responseBytes, err := json.Marshal(response)
	w.Write(responseBytes)
}
