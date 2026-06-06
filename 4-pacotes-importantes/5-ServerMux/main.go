package main

import "net/http"

func main() {
	mux := http.NewServeMux() //responsável por ataxar url
	// mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Hello World!"))
	// })
	mux.HandleFunc("/", HomeHandler)
	mux.Handle("/blog", blog{title: "My Blog"})
	mux.Handle("/blog2", blog{title: "My Blog 2"})
	http.ListenAndServe(":8080", mux)

	mux2 := http.NewServeMux() //responsável por ataxar url
	mux2.HandleFunc("/", HomeHandler)
	mux2.Handle("/luciana", blog{title: "My Luciana"})
	mux2.Handle("/luciana2", blog{title: "My Luciana 2"})
	http.ListenAndServe(":8080", mux2)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

type blog struct {
	title string
}

func (b blog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(b.title))
}
