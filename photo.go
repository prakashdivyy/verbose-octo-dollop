package main

import (
	"encoding/json"
	"net/http"
	"net/url"

	iris "gopkg.in/kataras/iris.v4"
)

type JSONFoto struct {
	Status  string `json:"status"`
	Reason  string `json:"reason"`
	Nama    string `json:"nama"`
	NPM     string `json:"npm"`
	Jurusan string `json:"jurusan"`
	Base64  string `json:"foto"`
}

func photoPage(c *iris.Context) {
	npm := c.Param("npm")
	payload := url.Values{}
	payload.Set("npm", npm)
	resp, _ := http.PostForm(config.URLFoto, payload)
	decodeJSON := json.NewDecoder(resp.Body)
	var data JSONFoto
	_ = decodeJSON.Decode(&data)
	output := data.Nama
	notErr := true
	if data.Status == "error" {
		output = data.Reason
		notErr = false
	}
	c.MustRender("foto.html", struct {
		Nama     string
		NPM      string
		Jurusan  string
		Foto     string
		NotError bool
	}{Nama: output, NPM: data.NPM, Jurusan: data.Jurusan, Foto: data.Base64, NotError: notErr})
}
