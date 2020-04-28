package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Post struct {
	Id    int
	Title string
	Body  string
}

func main() {

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {

		post := Post{Id: 1, Title: "Unamed Post", Body: "No content"}

		if title := r.FormValue("title"); title != "" {
			post.Title = title
		}

		t := template.Must(template.ParseFiles("templates/index.html"))

		if err := t.ExecuteTemplate(rw, "index.html", post); err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println(http.ListenAndServe(":9999", nil))
}
