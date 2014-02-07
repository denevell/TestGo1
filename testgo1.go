package main

import "net/http"
import "fmt"
import "html"
import "log"

func main() {
        http.HandleFunc("/hiya", func(w http.ResponseWriter, r *http.Request) {
                fmt.Fprintf(w, "AND A THIRD EDIT, WOO WEE, %q", html.EscapeString(r.URL.Path))
        })
        log.Fatal(http.ListenAndServe(":6002", nil))
}
