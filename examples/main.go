package main

import (
	"log"
	"fmt"
	"net/http"
	"github.com/saromanov/reverz"
)

// test handler
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func main(){
	rev, _ := reverz.New(&reverz.Config{
		URLs: []string{"http://127.0.0.1:8081"},
	})

	handler := func (w http.ResponseWriter, r *http.Request) {
		rev.Proxy(w, r)
		fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	}
	mux := http.NewServeMux()
	finalHandler := http.HandlerFunc(handler)
  	mux.Handle("/", finalHandler)

  	log.Println("Listening on :3000...")
  	err := http.ListenAndServe(":3000", mux)
  	log.Fatal(err)
}
