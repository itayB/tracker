package main

import (
	"io"
	"encoding/json"
	"net/http"
	"fmt"
)

var port = 7070

type Activity struct {
	Title	string
	Summary	string
	Tags	[]string
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hwllo web")
}

func sayBy(w http.ResponseWriter, r *http.Request) {
	activity := Activity{"intro title", "intro summary", []string{"golang", "programming"}}
	js, err := json.Marshal(activity)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func main() {
	fmt.Printf("Tracker web server in listening on port %d ...", port)

	http.HandleFunc("/", sayHello)
	http.HandleFunc("/123", sayBy)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
