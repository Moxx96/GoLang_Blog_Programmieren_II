package main

import (
	"log"
	"net/http"
	"html/template"
	"strings"
	"fmt"
	"time"
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
		compareUser := user{"0","0",false}
		validUser := compareUser
		for _, element := range Users{
			if element.name == username{
				if element.pwd == password{
					validUser = element
				}
			}
		}
		if validUser != compareUser{
			expiration := time.Now().Add(time.Hour/4)
			cookie := http.Cookie{Name: "username", Value: validUser.name, Secure: validUser.isAuthor, Expires: expiration}
			http.SetCookie(w, &cookie)
			responseString := 	"<html>"+
				"<body>"+
				"WELCOME"+
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


var Users []user
var username string
var password string

func main() {
	http.HandleFunc("/", mainHandler)
	log.Fatalln(http.ListenAndServeTLS(":4443","./ressources/certBlog.pem" ,"./ressources/keyBlog.pem",nil))
}