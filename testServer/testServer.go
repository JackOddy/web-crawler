package testServer

import (
	"net/http"
	"testing"
)

func TestServer(t *testing.T) func() {
	// func main() {
	server := http.Server{Addr: ":3000"}
	http.Handle("/", http.FileServer(http.Dir("./testServer/html")))

	go func() {
		if err := server.ListenAndServe(); err != nil {
			t.Fatalf("Error: %s", err)
		}
	}()

	return func() {
		if err := server.Shutdown(nil); err != nil {
			panic(err)
		}
	}
}
