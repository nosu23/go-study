package gcp_sample_func

import (
	"io"
	"fmt"
	"net/http"
)

func HelloWorld(writer http.ResponseWriter, request *http.Request) {
	//cloud functionsのログに文字列出力
	fmt.Println("ossu, ossu",)
	//cloud functionsのhttpレスポンスに文字列出力。
	io.WriteString(writer, "Hello World gen2    !!   ")
}