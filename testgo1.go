package main

import "net/http"
import "fmt"
import "log"

func main() {
        http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
                fmt.Fprintf(w, "W-w-w-website")
        })
        log.Fatal(http.ListenAndServe(":6002", nil))
}
