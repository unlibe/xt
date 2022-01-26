package rep

import "embed"

var jsonToGo embed.FS

func JsonToGoInit(dataWebFS embed.FS) {
	jsonToGo = dataWebFS
}

func GetJsonToGo() embed.FS {
	return jsonToGo
}
