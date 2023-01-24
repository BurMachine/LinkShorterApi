package httpHandlers

import (
	"log"
	"net/http"
)

func (s *HttpHandlers) GetOriginalUrl(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	log.Println("GET")
	w.Write([]byte("privet2"))
	// ...
}
