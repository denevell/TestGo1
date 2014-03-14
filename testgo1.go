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

var PORTFOLIO_BYTES []byte = nil

func main() {
        tmpl, err := template.ParseFiles("site.tmpl")
        if err != nil {
                log.Fatal(err)
        }

	http.Handle("/res/", http.FileServer(http.Dir("")))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		body, err := getPortfolioFile()
		if err!=nil {
			fmt.Fprintf(w, "Problem getting portfolio file.")
			return
		}
		projects := make([]Project, 1)
		err = json.Unmarshal(body, &projects)
		if err!=nil {
			fmt.Fprintf(w, "Problem unmarshaling.")
			return
		}
                err = tmpl.Execute(w,
                        map[string]interface{}{
                                "projects": projects,
                        })
		if err != nil {
			fmt.Fprintf(w, "Error: "+err.Error())
			return
		}
	})
	log.Fatal(http.ListenAndServe(":7001", nil))
}

func getPortfolioFile() ([]byte, error) {
	if PORTFOLIO_BYTES==nil {
		fmt.Println("Portfolio file doesn't exist")
		resp, err := http.Get("https://raw.github.com/denevell/BlogPosts/master/portfolio.json")
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		bytes, err := ioutil.ReadAll(resp.Body)
		PORTFOLIO_BYTES = bytes
		return bytes, err
	} else {
		return PORTFOLIO_BYTES, nil
	}
}
