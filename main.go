package main

import (
	"embed"
	"github.com/xknowledge/xt/cmd"
	"github.com/xknowledge/xt/rep"
)

//go:embed resource/json-to-go
var jsonToGo embed.FS

//go:embed resource/json-format
var jsonFormat embed.FS

func main() {
	rep.JsonToGoInit(jsonToGo)
	rep.JsonFormatInit(jsonFormat)
	cmd.Run()
}
