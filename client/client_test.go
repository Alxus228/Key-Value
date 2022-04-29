package client

import (
	"flag"
	"io"
	"strconv"
	"strings"
	"sync"
	"testing"
)

// We use flag records in order to set the amount of
// PUT requests to be created during one benchmark iteration.
var loc = flag.Int("records", 10, "How many records to do")

// Test BenchmarkSequentially tries to send a PUT request into the server *loc times.
//
// Afterwards, it receives a response from GET all request and compares data
// from the response body to data it has written while sending PUT requests.
func BenchmarkSequentially(b *testing.B) {
	for i := 0; i < b.N; i++ {
		//Putting j element into the key-values storage *loc times.
		for j := 0; j < *loc; j++ {
			_, err := Put("key"+strconv.Itoa(j), j)
			if err != nil {
				b.Fatal(err)
			}
		}

		//GET all request
		resp, err := GetAll()
		if err != nil {
			b.Fatal(err)
		}

		//Asserting results
		byteData, _ := io.ReadAll(resp.Body)
		stringData := string(byteData)
		for j := 0; j < *loc; j++ {
			if !strings.Contains(stringData, "key"+strconv.Itoa(j)) {
				b.Fatal(stringData)
			}
		}
	}
}

// Test BenchmarkConcurrently does exactly the same what BenchmarkSequentially does,
// but each PUT request is created in a new goroutine.
func BenchmarkConcurrently(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		for j := 0; j < *loc; j++ {
			wg.Add(1)
			go func(key int) {
				defer wg.Done()
				_, err := Put("key"+strconv.Itoa(key), key)
				if err != nil {
					b.Fatal(err)
				}
			}(j)
		}

		wg.Wait()

		resp, err := GetAll()
		if err != nil {
			b.Fatal(err)
		}

		byteData, _ := io.ReadAll(resp.Body)
		stringData := string(byteData)
		for j := 0; j < *loc; j++ {
			if !strings.Contains(stringData, "key"+strconv.Itoa(j)) {
				b.Fatal(stringData)
			}
		}
	}
}
