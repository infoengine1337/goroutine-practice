package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/imroc/req/v3"
)

func main() {

	before := time.Now()
	fmt.Println("[*] Starting Program...")

	var wg sync.WaitGroup

	for i := 0; i <= 3000; i++ {

		wg.Add(1)
		go func() {

			defer wg.Done()

			//client := req.C()
			//resp, err := client.R().
			//				Get("http://localhost:8000")

			resp := req.MustGet("http://localhost:8000")

			//if err != nil {
			//	log.Fatal(err)
			//}

			fmt.Println(resp)

		}()

	}

	wg.Wait()

	after := time.Now()

	fmt.Printf("[*] It takes %f second.\n", (after.Sub(before)).Seconds())
}
