package main

import (
	"net/http"
	"html/template"
)

type login struct{
	USERNAME string
	MODUS string
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	c,_ := r.Cookie("username")
	c2,_:= r.Cookie("isAuthor")
	t := template.New("Test")
	t, _ = template.ParseFiles("./ressources/html/blog.html")
	var modus string
	if c2.Value == "0"{
		modus = "Author"
	}else{
		modus = "Leser"
	}
	p := login{USERNAME: c.Value, MODUS: modus}
	t.Execute(w,p)



}