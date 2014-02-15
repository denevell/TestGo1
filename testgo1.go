package main

import "net/http"
import "fmt"
import "log"
import "text/template"

type Feeling struct {
	HowMuch string
	What    string
}

var feeling = []Feeling{
	Feeling{"moderately", "tired"},
	Feeling{"somewhat", "exuberant"},
}

var templateStr = `I am feeling {{range $i, $e := .feeling}}{{$e.HowMuch}} {{$e.What}}{{if $i}}.{{else}}, {{end}}{{end}}`

func main() {
	tmpl, err := template.New("test").Parse(templateStr)
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := tmpl.Execute(w,
			map[string]interface{}{
				"feeling": feeling,
			})
		if err != nil {
			fmt.Fprintf(w, "Error: "+err.Error())
		}
	})
	log.Fatal(http.ListenAndServe(":7001", nil))
}
