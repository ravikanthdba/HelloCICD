package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	// Set up the load test parameters
	numRequests := 100000
	concurrency := 10

	// Create a wait group to wait for all requests to complete
	var wg sync.WaitGroup
	wg.Add(numRequests)

	// Create a channel to receive errors
	errChan := make(chan error)

	// Create a client to make HTTP requests
	client := http.Client{
		Timeout: time.Duration(5 * time.Second),
	}

	// Send requests in parallel
	for i := 0; i < concurrency; i++ {
		go func() {
			for j := 0; j < numRequests/concurrency; j++ {
				// Make an HTTP request
				resp, err := client.Get("http://localhost:1001/")

				// Check for errors
				if err != nil {
					errChan <- err
					continue
				}

				// Check the response status code
				if resp.StatusCode != 200 {
					errChan <- fmt.Errorf("unexpected status code: %d", resp.StatusCode)
				}

				// Close the response body
				resp.Body.Close()

				// Mark the request as completed
				wg.Done()
			}
		}()
	}

	// Wait for all requests to complete
	go func() {
		wg.Wait()
		close(errChan)
	}()

	// Print any errors that occurred during the load test
	for err := range errChan {
		fmt.Println(err)
	}
}
