package rep

import "embed"

var jsonToGo embed.FS
var jsonFormat embed.FS

func JsonToGoInit(dataWebFS embed.FS) {
	jsonToGo = dataWebFS
}

func GetJsonToGo() embed.FS {
	return jsonToGo
}

func JsonFormatInit(dataWebFS embed.FS) {
	jsonFormat = dataWebFS
}

func GetJsonJsonFormat() embed.FS {
	return jsonFormat
}
