package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/subkaitaku/connpass/event"
)

func main() {
	fmt.Println("listen and serve on :5000")
	http.HandleFunc("/", event.RenderEvents)
	http.HandleFunc("/register", event.RegisterBlock)
	err := http.ListenAndServe("127.0.0.1:6000", nil)

	if err != nil {
		fmt.Printf("error: %v", err)
		os.Exit(1)
	}
}
