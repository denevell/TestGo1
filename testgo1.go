package main

import "net/http"
import "fmt"
import "html"
import "log"

func main() {
        http.HandleFunc("/hiya", func(w http.ResponseWriter, r *http.Request) {
                fmt.Fprintf(w, "Hiiiya, %q", html.EscapeString(r.URL.Path))
        })
        log.Fatal(http.ListenAndServe(":8081", nil))
}
