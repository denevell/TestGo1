package main

import "net/http"
import "fmt"
import "html"
import "log"

func main() {
        http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
                fmt.Fprintf(w, "W-w-w-website, %q", html.EscapeString(r.URL.Path))
        })
        log.Fatal(http.ListenAndServe(":6002", nil))
}
