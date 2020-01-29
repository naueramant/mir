package server

import (
	"net/http"
	"strconv"

	"github.com/gobuffalo/packr"
)

const (
	Port = 8109
)

func Start() {
	box := packr.NewBox("../../assets/html")
	http.Handle("/", http.FileServer(box))

	http.ListenAndServe(":"+strconv.Itoa(Port), nil)
}
