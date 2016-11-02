package main

import (
	"encoding/json"
	"net/url"

	"net/http"

	iris "gopkg.in/kataras/iris.v4"
)

type JSONData struct {
	Data []DataMhs `json:"data"`
}

type DataMhs struct {
	NPM   string `json:"npm"`
	Nama  string `json:"nama"`
	Email string `json:"email"`
	Fak   string `json:"fakultas"`
}

func searchPage(c *iris.Context) {
	c.MustRender("search.html", nil)
}

func checkQuery(c *iris.Context) {
	qs := c.FormValueString("qs")
	payload := url.Values{}
	payload.Set("qs", qs)
	resp, _ := http.PostForm(config.URLData, payload)
	decodeJSON := json.NewDecoder(resp.Body)
	var data JSONData
	_ = decodeJSON.Decode(&data)
	c.JSON(iris.StatusOK, data)
}
