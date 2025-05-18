package main

import (
	"fmt"
	"time"

	"github.com/Gab-Mello/GoPinger/internal/monitor"
)

func main() {

	fmt.Println("Initializing monitoring...")

	url := "https://www.google.com"
	interval := 5 * time.Second

	for {
		monitor.CheckSite(url)
		time.Sleep(interval)
	}

}
