package main

import (
	"log"
	"net/http"
	"html/template"
	"strings"
	"fmt"
	"time"
	"encoding/xml"
)


func mainHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("method:", r.Method) //get request method
	if r.Method == "GET" {
		t, _ := template.ParseFiles("./ressources/html/login.html")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		//fmt.Println("username:", r.Form["username"])
		//fmt.Println("password:", r.Form["password"])
		username = strings.Join(r.Form["username"],"")
		password = strings.Join(r.Form["password"],"")
		fmt.Print(username)
		fmt.Print(password)
		Users = readUsers()

		var xmNull xml.Name
		compareUser := user{xmNull,"","",""}
		validUser := compareUser
		for _, element := range Users{
			if element.Name == username{
				if element.Password == password{
					validUser = element
				}
			}
		}
		if validUser != compareUser{
			fmt.Print(validUser.Author)
			expiration := time.Now().Add(time.Hour/4)
			cookie := http.Cookie{Name: "username", Value: validUser.Name, Expires: expiration}
			cookie2 := http.Cookie{Name: "isAuthor", Value: validUser.Author, Expires: expiration}
			http.SetCookie(w, &cookie)
			http.SetCookie(w, &cookie2)
			//homeHandler(w,r)
			responseString := 	"<html>"+
				"<body>"+
				"<h1>Programmieren II - Blog</h1><br>"+
				"Login erfolgreich "+"<a href='/home'>Bitte Klicken</a>"+
				"</body>"+
				"</html>"
			w.Write([]byte(responseString))
		}else {
			responseString := 	"<html>"+
				"<body>"+
				"Falsches Passwort oder Benutzername"+
				"</body>"+
				"</html>"
			w.Write([]byte(responseString))
		}
	}

}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	c,_ := r.Cookie("username")
	c2,_ := r.Cookie("isAuthor")
	c.Name = "Expired"
	c2.Name = "Expired"
	c.Expires = time.Now()
	c2.Expires = time.Now()
	t, _ := template.ParseFiles("./ressources/html/logout.html")
	t.Execute(w, nil)
}


var Users []user
var username string
var password string

func main() {
	http.HandleFunc("/", mainHandler)
	http.HandleFunc("/home/", homeHandler)
	http.HandleFunc("/logout/",logoutHandler)
	log.Fatalln(http.ListenAndServeTLS(":4443","./ressources/certBlog.pem" ,"./ressources/keyBlog.pem",nil))
}