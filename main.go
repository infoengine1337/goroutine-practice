package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/valyala/fasthttp"
)

func main() {

	before := time.Now()
	fmt.Println("[*] Starting Program...")

	var wg sync.WaitGroup

	for i := 0; i <= 3000; i++ {

		wg.Add(1)
		go func() {

			defer wg.Done()

			_, _, err := fasthttp.Get(nil, "http://localhost:8000")

			if err != nil {
				log.Fatal(err)
			}

		}()

	}

	wg.Wait()

	after := time.Now()

	fmt.Printf("[*] It takes %f second.\n", (after.Sub(before)).Seconds())
}
