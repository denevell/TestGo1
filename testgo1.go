package main

import "net/http"
import "fmt"
import "log"
import "io/ioutil"
import "encoding/json"
import "html/template"

type Project struct {
	Title string
	Description string
	Features []string
	Keywords []string
} 

var templateStr = `
	{{range $i,$v := .projects}}
		<h3>{{.Title}}</h3>
		{{.Description}}<br />
		<br />
		Features:
		<ul>
		{{range $i,$v := .Features}}
			<li>{{$v}}</li>	
		{{end}}
		</ul>
		<div>
		Keywords:
		{{range $i,$v := .Keywords}}{{if $i}}, {{end}}{{$v}}{{end}}
		</div>
		<hr />
	{{end}}
`

func main() {
        tmpl, err := template.New("test").Parse(templateStr)
        if err != nil {
                log.Fatal(err)
        }

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		resp, err := http.Get("https://raw.github.com/denevell/BlogPosts/master/portfolio.json")
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)

		projects := make([]Project, 1)
		err = json.Unmarshal(body, &projects)
		if err!=nil {
			fmt.Fprintf(w, "Problem unmarshaling.")
		}
		fmt.Fprintf(w,"<html>")
                err = tmpl.Execute(w,
                        map[string]interface{}{
                                "projects": projects,
                        })
		if err != nil {
			fmt.Fprintf(w, "Error: "+err.Error())
		}
		fmt.Fprintf(w,"</html>")
	})
	log.Fatal(http.ListenAndServe(":7001", nil))
}
