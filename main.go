package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"sync"
)

type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ =
			template.Must(template.ParseFiles(t.filename))
	})
	t.templ.Execute(w, r)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := newRoom()
	http.Handle("/", &templateHandler{filename: "index.html"})
	http.Handle("/room", r)

	go r.run()

	log.Print("listen :" + port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
}
