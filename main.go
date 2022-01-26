package main

import (
	"embed"
	"xt/cmd"
	"xt/rep"
)

//go:embed resource/json-to-go
var jsonToGo embed.FS

func main() {
	rep.JsonToGoInit(jsonToGo)
	cmd.Run()
}
