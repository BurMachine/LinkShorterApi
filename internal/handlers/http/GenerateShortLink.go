package httpHandlers

import (
	"net/http"
)

func (s *HttpHandlers) GenerateShortLink(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	println("POST")
}
