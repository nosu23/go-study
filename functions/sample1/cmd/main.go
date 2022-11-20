package main

import (
	"context"
	"log"
	"github.com/nosu23/go-study/functions/sample1"
	"github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
)

func main() {

	ctx := context.Background()

	hello := sample1.HelloWorld

	// 実行する関数の登録
	if err := funcframework.RegisterHTTPFunctionContext(ctx, "/", hello); err != nil {
		log.Fatalf("funcframework.RegisterHTTPFunctionContext: %v\n", err)
	}

	port := "8080"
	if err := funcframework.Start(port); err != nil {
		log.Fatalf("funcframework.Start: %v\n", err)
	}
}
