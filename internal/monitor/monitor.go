package monitor

import (
	"fmt"
	"net/http"
	"time"
)

func CheckSite(url string) {
	start := time.Now()
	resp, err := http.Get(url)
	duration := time.Since(start)

	if err != nil {
		fmt.Printf("Error when accessing %s: %s\n", url, err)
		return
	}

	defer resp.Body.Close()

	fmt.Printf("test")

	fmt.Printf("URL: %s | Stats: %d | Reponse time: %v\n", url, resp.StatusCode, duration)
}
