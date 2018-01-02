package main

import (
	"log"
	"net/http"
	"html/template"
	"strings"
	"fmt"
	"time"
	"encoding/xml"
	"strconv"
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
			expiration := time.Now().Add(time.Minute*15)
			cookie := http.Cookie{Name: "username", Value: validUser.Name, Expires: expiration, Path: "/"}
			cookie2 := http.Cookie{Name: "isAuthor", Value: validUser.Author, Expires: expiration, Path: "/"}
			cookie3 :=http.Cookie{Name:"timestamp", Value: strconv.FormatInt(time.Now().Unix(), 10), Path: "/", Expires: expiration}
			http.SetCookie(w, &cookie)
			http.SetCookie(w, &cookie2)
			http.SetCookie(w, &cookie3)
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

func guestHandler(w http.ResponseWriter, r *http.Request) {
	expiration := time.Now().Add(time.Hour/4)
	cookie := &http.Cookie{Name: "username", Value: "Guest", Expires: expiration, Path: "/"}
	cookie2 := &http.Cookie{Name: "isAuthor", Value: "1", Expires: expiration , Path: "/"}
	cookie3 := &http.Cookie{Name:"timestamp", Value: strconv.FormatInt(time.Now().Add(15*time.Minute).Unix(), 10), Path: "/", Expires: expiration}
	http.SetCookie(w, cookie)
	http.SetCookie(w, cookie2)
	http.SetCookie(w, cookie3)
	responseString := 	"<html>"+
		"<body>"+
		"<h1>Programmieren II - Blog</h1><br>"+
		"Gastzugang erfolgreich "+"<a href='/home'>Bitte Klicken</a>"+
		"</body>"+
		"</html>"
	w.Write([]byte(responseString))
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	//Funktioniert noch nicht
	c,_ := r.Cookie("username")
	c2,_ := r.Cookie("isAuthor")
	c3,_:= r.Cookie("timestamp")
	c.Expires = time.Now()
	c2.Expires = time.Now()
	c3.Value = strconv.FormatInt(time.Now().Add(-24*time.Hour).Unix(),10)
	c3.Path = "/"
	http.SetCookie(w,c3)
	responseString := 	"<html>"+
		"<body>"+
		"<h1>Programmieren II - Blog</h1><br>"+
		"Logout erfolgreich "+"<a href='/'>Bitte Klicken</a>"+
		"</body>"+
		"</html>"
	w.Write([]byte(responseString))
}


var Users []user
var username string
var password string

func main() {
	http.HandleFunc("/", mainHandler)
	http.HandleFunc("/guest/", guestHandler)
	http.HandleFunc("/home/", homeHandler)
	http.HandleFunc("/logout/",logoutHandler)
	log.Fatalln(http.ListenAndServeTLS(":4443","./ressources/certBlog.pem" ,"./ressources/keyBlog.pem",nil))
}