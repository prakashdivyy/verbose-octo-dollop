package main

import (
	"encoding/json"
	"os"

	"github.com/kataras/go-template/html"
	"gopkg.in/iris-contrib/middleware.v4/logger"
	iris "gopkg.in/kataras/iris.v4"
)

// ConfigFile from config.json
type ConfigFile struct {
	URLData string `json:"url_data"`
	URLFoto string `json:"url_foto"`
}

var config ConfigFile

func main() {
	loadConfig()
	startIris()
}

func loadConfig() {
	file, _ := os.Open("config.json")
	parseJSON := json.NewDecoder(file)
	parseJSON.Decode(&config)
	file.Close()
}

func startIris() {
	iris.Config.IsDevelopment = true
	iris.UseTemplate(html.New()).Directory("resources/templates", ".html")
	iris.Static("/js", "resources/js", 1)
	iris.Use(logger.New())
	iris.Get("/", searchPage)
	iris.Post("/", checkQuery)
	iris.Get("/foto/:npm", photoPage)
	iris.Listen(":6969")
}
