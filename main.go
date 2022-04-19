package main

import (
	"fmt"
	"io"
	"net/http"
)

func bonjour(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "anselbrandt.dev<br/><br/><a href=\"/headers\">headers</a><br/><br/><a href=\"/kanye\">kanye quotes</a>")
}

func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func api(w http.ResponseWriter, req *http.Request) {

	resp, err := http.Get("https://anselbrandt.com/api")
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", resp.Header.Get("Content-Type"))
	w.Header().Set("Content-Length", resp.Header.Get("Content-Length"))
	io.Copy(w, resp.Body)
	resp.Body.Close()
}

func kanye(w http.ResponseWriter, req *http.Request) {

	resp, err := http.Get("https://api.kanye.rest")
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.Copy(w, resp.Body)
	resp.Body.Close()
}

func main() {
	fmt.Println("server is running...")

	http.HandleFunc("/", bonjour)
	http.HandleFunc("/api", api)
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/kanye", kanye)

	http.ListenAndServe(":8080", nil)
}
