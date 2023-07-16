package main

import (
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"math/rand"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

var urlMap = make(map[string]string)

type request struct {
	URL string `json:"url"`
}

type response struct {
	ShortURL string `json:"short_url"`
}

func shortenHandler(w http.ResponseWriter, r *http.Request) {
	var req resquest
	json.NewDecoder(r.Body).Decode(&req)
	shortUrl := generateID(req.URL)
	urlMap[shortUrl] = req.URL
	res := response{ShortURL: shortUrl}
	json.NewEncoder(w).Encode(res)

}
func generateID(url string) string {
	h := sha1.New()
	h.Write([]byte(url))
	hash := h.Sum(nil)
	return base64.URLEncoding.EncodeToString(hash[:6])
}
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/shorten", shortenHandler)
	http.ListenAndServe()(":8080", r)
}

// To continue:

// Implement a redirect handler
// Lookup short code in urlMap
// Redirect to original URL
// Add routes for both handlers
