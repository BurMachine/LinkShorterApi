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

var links = []string{
	"https://www.google.ru/search?q=test",
	"https://www.yandex.ru/news",
	"http://ozon.ru/test",
	"https://www.example.ru/",
	"http://www.example.su/",
	"https://ru.wikipedia.org/wiki/Main_Page",
	"http://su.wikipedia.org/wiki/Main_Page",
	"https://github.com/ru",
	"http://github.com/su",
	"https://www.example.ru/test?q=test",
	"http://www.example.su/test?q=test",
	"http://www.google.com",
	"https://www.facebook.com",
	"https://www.youtube.com",
	"https://www.instagram.com",
	"https://www.linkedin.com",
	"https://www.amazon.com",
	"https://www.twitter.com",
	"https://www.apple.com",
	"https://www.github.com",
	"https://www.microsoft.com",
	"https://www.spotify.com",
	"https://www.tiktok.com",
	"https://www.pinterest.com",
	"https://www.reddit.com",
	"https://www.dropbox.com",
	"https://www.whatsapp.com",
	"https://www.quora.com",
	"https://www.medium.com",
	"https://www.hulu.com",
	"https://www.zoom.us",
	"https://www.airbnb.com",
	"https://www.uber.com",
	"https://www.lyft.com",
	"https://www.tinder.com",
	"https://www.grindr.com",
	"https://www.buzzfeed.com",
	"https://www.nytimes.com",
	"https://www.washingtonpost.com",
	"https://www.cnn.com",
	"https://www.bbc.com",
	"https://www.alibaba.com",
	"https://www.tencent.com",
	"https://www.xiaomi.com",
	"https://www.huawei.com",
	"https://www.bytedance.com",
	"https://www.paypal.com",
	"https://www.stripe.com",
	"https://www.square.com",
	"https://www.salesforce.com",
	"https://www.zendesk.com",
	"https://www.shopify.com",
	"https://www.squarespace.com",
	"https://www.wix.com",
	"https://www.weebly.com",
	"https://www.godaddy.com",
}

func TestGenerateShortLink(t *testing.T) {

	var storageM storage.ServiceStorage
	storageM = storage.NewInMemoryStorageInit()
	var handlers HttpHandlers
	handlers.Storage = &storageM

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
	}
}
