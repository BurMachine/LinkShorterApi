package httpHandlers

import (
	"burmachine/LinkGenerator/internal/models"
	"burmachine/LinkGenerator/internal/storage"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

var invalidIds = []string{
	"1id@%^&*",
	"id with spaces",
	"id!",
	"long_id_12345678901",
	"shrtid",
	"Inv@lidCh@rs",
	"IDwith*",
	"sp@ci@l!",
	"IDwith#",
	"numb3rs@nd_l3tt3rs",
	"invalid id",
	"not_10_chars",
	"1234567890!",
	"notvalid",
	"abcdefghijklmnopqrst",
	"spec!al_chars",
	"lettersand1",
	"numbersandletters",
	"one two",
	"specialcharacters!",
	"letters_numbers",
	"longerthan10",
	"shorterthan10",
	"lettersand$",
	"letters&numbers",
	"letters_and-numbers",
	"letters_and_symbols",
	"1@3$5^7*9",
	"letters@numbers",
	"letters-numbers",
	"letters_numbers_",
	"numbers_letters",
	"letters+numbers",
	"letters*numbers",
	"letters&numbers&",
	"letters_and.numbers",
	"letters_and_symbols_",
	"letters_and/numbers",
	"letters-and_numbers",
	"letters+and_numbers",
	"letters_and(numbers)",
	"letters_and_symbols*",
	"letters_and_symbols&",
	"letters_and_symbols_and#",
	"letters_and_symbols_and@",
	"letters_and_symbols_and%",
	"letters_and_symbols_and^",
}

func TestGetOriginalLinkValid(t *testing.T) {
	var storageM storage.ServiceStorage
	storageM = storage.NewInMemoryStorageInit()
	var handlers HttpHandlers
	handlers.Storage = &storageM

	var validIds []string

	for _, link := range links {
		body := fmt.Sprintf(`{"url":"%s"}`, link)
		req, err := http.NewRequest("POST", "/generate", bytes.NewBuffer([]byte(body)))
		if err != nil {
			t.Fatal(err)
		}
		w := httptest.NewRecorder()

		handlers.GenerateShortLink(w, req, nil)

		if w.Code != http.StatusOK {
			t.Errorf("Expected status code %d, but got %d", http.StatusOK, w.Code)
		}
		var response models.ResponseBody
		if err := json.NewDecoder(w.Body).Decode(&response); err != nil {
			t.Fatal(err)
		}
		validIds = append(validIds, response.ShortUrl)
	}

	for _, id := range validIds {
		req, err := http.NewRequest("GET", fmt.Sprintf("/?short=%s", id), nil)
		if err != nil {
			t.Fatal(err)
		}
		w := httptest.NewRecorder()

		m := make(map[string]string)
		handlers.GetOriginalUrl(w, req, m)

		if w.Code != http.StatusOK {
			t.Errorf("Expected status code %d, but got %d", http.StatusOK, w.Code)
		}
		var response models.ResponseBody
		if err := json.NewDecoder(w.Body).Decode(&response); err != nil {
			t.Fatal(err)
		}
	}
}

func TestGetOriginalLinkInvalid(t *testing.T) {
	var storageM storage.ServiceStorage
	storageM = storage.NewInMemoryStorageInit()
	var handlers HttpHandlers
	handlers.Storage = &storageM

	for _, id := range invalidIds {
		req, err := http.NewRequest("GET", fmt.Sprintf("/?short=%s", id), nil)
		if err != nil {
			t.Fatal(err)
		}
		w := httptest.NewRecorder()

		handlers.GetOriginalUrl(w, req, map[string]string{})

		if w.Code != http.StatusBadRequest {
			t.Errorf("Expected status code %d, but got %d", http.StatusBadRequest, w.Code)
		}
	}
}
