package main

import (
	"fmt"
	"net/http"
	"time"
)

func checkURL(url string) {
	start := time.Now()

	response, err := http.Get(url)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	defer response.Body.Close()

	latency := time.Since(start)

	fmt.Println("URL:", url)
	fmt.Println("Status:", response.StatusCode)
	fmt.Println("Latency:", latency)
	fmt.Println("--------------------")
}

func main() {
	checkURL("https://google.com")
	checkURL("https://github.com")
	checkURL("https://example.com")
}
