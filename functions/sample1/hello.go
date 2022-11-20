package sample1

import (
	"io"
	"net/http"
)

func HelloWorld(writer http.ResponseWriter, request *http.Request) {
	io.WriteString(writer, "Hello World")
}