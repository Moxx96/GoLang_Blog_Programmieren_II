package main

import "net/http"

func homeHandler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	name := q.Get("name")
	if name == "" {
		name = "World"
	}
	responseString := 	"<html>"+
		"<body>"+
		"WELCOME"+
		"</body>"+
		"</html>"
	w.Write([]byte(responseString))
}