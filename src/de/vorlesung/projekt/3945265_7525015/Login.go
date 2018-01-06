package main

import (
	"net/http"
	"time"
	"strconv"
	"strings"
	"encoding/xml"
	"crypto/sha256"
	"encoding/hex"
	"html/template"
)

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {														//Prüft ob GET Anfrage
		t, _ := template.ParseFiles("./ressources/html/login.html")	//
		t.Execute(w, nil)													//Wenn ja, dann html Ausgabe
	} else {
		r.ParseForm()															//Wenn nein, dann Post Methode, nach Eingabe der Credentials
		username = strings.Join(r.Form["username"],"")
		password = strings.Join(r.Form["password"],"")						//Username u. Password auslesen
		Users = readUsers()														//Alle vorhandenen User einelesen

		var xmNull xml.Name
		compareUser := user{xmNull,"","","", ""}	//Leeren User zum vergleichen erstellen
		validUser := compareUser														//Variable für einen gültigen User erstellen
		for _, element := range Users{							//Alle User durchiterieren
			if element.Name == username{						//Bei übereinstimmendem Namen
				salted := password + element.Salt				//Berechnen des Hashes von Passwort + Salt
				hash := sha256.New()							//
				hash.Write([]byte(salted)) 						//
				pwdhash := hex.EncodeToString(hash.Sum(nil))	//Konvertieren von Typ byte zu string
				if element.Password == pwdhash{					//Bei übereinstimmenden Passwort
					validUser = element							//Als gültigen Usersetzen
				}
			}
		}
		if validUser != compareUser{							//Falls ein gültiger USer gefunden wurde
			//fmt.Print(validUser.Author)
			expiration := time.Now().Add(time.Minute*15)		//Setze Abluafzeit für Cookies
			cookie := http.Cookie{Name: "username", Value: validUser.Name, Expires: expiration, Path: "/"}
			cookie2 := http.Cookie{Name: "isAuthor", Value: validUser.Author, Expires: expiration, Path: "/"}
			cookie3 :=http.Cookie{Name:"timestamp", Value: strconv.FormatInt(time.Now().Unix(), 10), Path: "/", Expires: expiration}
			http.SetCookie(w, &cookie)
			http.SetCookie(w, &cookie2)
			http.SetCookie(w, &cookie3)				//Setze Cookies mit Name, AuthorTag und timestamp
			//homeHandler(w,r)
			responseString := 	"<html>"+
				"<body>"+
				"<h1>Programmieren II - Blog</h1><br>"+
				"Login erfolgreich "+"<a href='/home'>Bitte Klicken</a>"+
				"</body>"+
				"</html>"
			w.Write([]byte(responseString))
		}else {									//Bei ungültigen Credentials Fehlermeldung
			responseString := 	"<html>"+
				"<body>"+
				"<h1>Programmieren II - Blog</h1><br>"+
				"Falsches Passwort oder Benutzername "+"<a href='/'>Bitte Klicken</a>"+
				"</body>"+
				"</html>"
			w.Write([]byte(responseString))
		}
	}

}

func guestHandler(w http.ResponseWriter, r *http.Request) {			//Wird als anmelden geklickt
	expiration := time.Now().Add(time.Hour/4)
	cookie := &http.Cookie{Name: "username", Value: "Guest", Expires: expiration, Path: "/"}
	cookie2 := &http.Cookie{Name: "isAuthor", Value: "1", Expires: expiration , Path: "/"}
	cookie3 := &http.Cookie{Name:"timestamp", Value: strconv.FormatInt(time.Now().Add(15*time.Minute).Unix(), 10), Path: "/", Expires: expiration}
	http.SetCookie(w, cookie)
	http.SetCookie(w, cookie2)
	http.SetCookie(w, cookie3)					//Setze Cookies zugeschnitten auf Gast Account
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

