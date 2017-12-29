package main

import (
	"net/http"
	"html/template"
)

type login struct{
	USERNAME string
	MODUS string
}

type beitrag struct{
	TEXT string
	DATUM string
	AUTHOR string
	COMMENTS []string
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

	t, _ = template.ParseFiles("./ressources/html/beitraege.html")
	m := beitrag{TEXT: "Ich bin Blog Text! Ich bin Blog Text! Ich bin Blog Text! Ich bin Blog Text! Ich bin Blog Text! Ich bin Blog Text! Ich bin Blog Text! Ich bin Blog Text! Ich bin Blog Text! Ich bin Blog Text!",
				DATUM: "29.12.2017",
				AUTHOR: "Author",
				COMMENTS: []string{"Kommentator: Ich bin ein Kommentar!\n","Kommentator 2: Ich bin noch ein Kommentar :P\n"}}
	t.Execute(w,m)



}