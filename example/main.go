package main

import (
	"context"
	"os"

	"golang.design/x/code2img"
)

func main() {
	f, err := os.ReadFile("main.go")
	if err != nil {
		panic(err)
	}
	code := string(f)

	// render it!
	b, err := code2img.Render(context.TODO(), code2img.LangGo, code)
	if err != nil {
		panic(err)
	}

	os.WriteFile("code.png", b, os.ModePerm)
}
