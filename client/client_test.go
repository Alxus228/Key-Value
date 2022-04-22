package client

import (
	"flag"
	"io"
	"strconv"
	"strings"
	"testing"
)

var loc = flag.Int("records", 10, "How many records to do")

func BenchmarkSequentially(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < *loc; j++ {
			_, err := Put("key"+strconv.Itoa(j), j)
			if err != nil {
				b.Fatal(err)
			}
		}

		resp, err := GetAll()
		if err != nil {
			b.Fatal(err)
		}

		byteData, _ := io.ReadAll(resp.Body)
		stringData := string(byteData)
		for j := 0; j < *loc; j++ {
			if !strings.Contains(stringData, "key"+strconv.Itoa(j)) {
				b.Fatal(err)
			}
		}
	}
}
