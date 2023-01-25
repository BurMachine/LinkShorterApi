package httpHandlers

import (
	"burmachine/LinkGenerator/internal/models"
	"burmachine/LinkGenerator/internal/service"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func (s *HttpHandlers) GenerateShortLink(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	if r.Method != http.MethodPost {
		println("method post error")
	}

	var body models.RequestBody
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Parse body error: ", err)
		w.Write([]byte("400 Bad Request"))
		return
	}
	err = json.Unmarshal(bodyBytes, &body)
	if err != nil {
		log.Println("Unmarshalling error in parsing post request's body: ", err)
		w.Write([]byte("400 Bad Request"))
		return
	}
	shortLink, err := service.GenerateLink(body.OriginalUrl)
	if err != nil {
		log.Println("Link generation error: ", err)
		w.Write([]byte("400 Bad Request\nLink Exist"))
		return
	}
	err = (*s.Storage).AddShortLink(body.OriginalUrl, shortLink)
	if err != nil {
		log.Println("Link add error: ", err)
		w.Write([]byte("400 Bad Request\nLink Exist"))
		return
	}

	response := models.ResponseBody{ShortUrl: shortLink}
	responseBytes, err := json.Marshal(response)
	w.Write(responseBytes)
}
