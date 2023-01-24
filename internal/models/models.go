package models

type RequestBody struct {
	OriginalUrl string `json:"url"`
}

type ResponseBody struct {
	ShortUrl string `json:"short_url"`
}
