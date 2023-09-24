package main

import (
	"context"
	"log"
	"os"
)

func main() {
	w, err := os.Create("index.html")
	if err != nil {
		log.Fatal(err)
	}
	defer w.Close()
	page("https://www.google.com").Render(context.Background(), w)
}
