package handler

import (
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/skip2/go-qrcode"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	url := r.FormValue("url")
	if url == "" {
		http.Error(w, "URL is required", http.StatusBadRequest)
		return
	}

	var png []byte
	png, err := qrcode.Encode(url, qrcode.Medium, 1000)
	if err != nil {
		http.Error(w, "Failed to generate QR code", http.StatusInternalServerError)
		return
	}

	base64Image := base64.StdEncoding.EncodeToString(png)
	imgSrc := "data:image/png;base64," + base64Image

	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "<img src=\"%s\" width=\"250\" height=\"250\"/>", imgSrc)
}
