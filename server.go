package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"github.com/eacolina/Google_Home_TTC/models/dialogflow/Request"
)

var count int
var mu sync.Mutex

func main() {
	fmt.Println("Starting server...")
	fmt.Println("Listening on localhost:8080")
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	http.HandleFunc("/actions", actions)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))

}

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q", r.URL.Path)
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprint(w, "Count: %d", count)
	mu.Unlock()
}
func actions(w http.ResponseWriter, r *http.Request) {
	reader := r.Body
	defer reader.Close()
	var body Request.WebhookRequest  // create empty Webhook object
	b, _ := ioutil.ReadAll(reader) // create a byte array from the reader stream
	json.Unmarshal(b,&body) // parse JSON byte array into body
	w.WriteHeader(http.StatusOK)

}
