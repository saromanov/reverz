package main

import (
	"log"
	"fmt"
	"net/http"
	"github.com/saromanov/reverz"
)


func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func main(){
	r, _ := reverz.New(&reverz.Config{
		URLs: []string{"localhost:8081"},
	})
	r.Run(handler)
	fmt.Println("RRRR")
	mux := http.NewServeMux()
	finalHandler := http.HandlerFunc(r.Run(handler))
  	mux.Handle("/", finalHandler)

  	log.Println("Listening on :3000...")
  	err := http.ListenAndServe(":3000", mux)
  	log.Fatal(err)
}