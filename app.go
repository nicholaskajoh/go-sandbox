package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	res, _ := http.Get("https://alphacoder.xyz/blog/index.xml")
	bytes, _ := ioutil.ReadAll(res.Body)
	bodyStr := string(bytes)
	fmt.Println(bodyStr)
	res.Body.Close()

	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8081", nil)
}
