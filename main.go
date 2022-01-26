package main

import (
	"embed"
	"xt/cmd"
	"xt/rep"
)

//go:embed resource/json-to-go
var jsonToGo embed.FS
//go:embed resource/json
var jsonFormat embed.FS

func main() {
	rep.JsonToGoInit(jsonToGo)
	rep.JsonFormatInit(jsonFormat)
	cmd.Run()
}
