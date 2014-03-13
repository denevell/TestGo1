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
	Images []string
} 

func main() {
        tmpl, err := template.ParseFiles("site.tmpl")
        if err != nil {
                log.Fatal(err)
        }

	http.Handle("/res/", http.FileServer(http.Dir("")))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		resp, err := http.Get("https://raw.github.com/denevell/BlogPosts/master/portfolio.json")
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)

		projects := make([]Project, 1)
		err = json.Unmarshal(body, &projects)
		if err!=nil {
			fmt.Fprintf(w, "Problem unmarshaling.")
		}
                err = tmpl.Execute(w,
                        map[string]interface{}{
                                "projects": projects,
                        })
		if err != nil {
			fmt.Fprintf(w, "Error: "+err.Error())
		}
	})
	log.Fatal(http.ListenAndServe(":7001", nil))
}
