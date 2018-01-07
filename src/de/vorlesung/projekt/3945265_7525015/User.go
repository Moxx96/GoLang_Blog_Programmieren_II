package main
//Matrikelnummern: 3945265, 7525015
import (
	"io/ioutil"
	"fmt"
	"os"
	"encoding/xml"
	"net/http"
	"strings"
	"html/template"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}


type Recurlyusers struct {
	XMLName     xml.Name `xml:"users"`
	Version     string   `xml:"version,attr"`
	Svs         []user   `xml:"user"`
}

type user struct {
	XMLName    xml.Name 	`xml:"user"`
	Name 		string      `xml:"Name"`
	Password   	string 	`xml:"Password"`
	Author  	string   	`xml:"Author"`
	Salt 		string		`xml:"Salt"`
}

type users struct {
	XMLName     xml.Name `xml:"users"`
	Version     string   `xml:"version,attr"`
	Svs         []writeuser   `xml:"user"`
}

type writeuser struct {
	Name 		string      `xml:"Name"`
	Password   	string 	    `xml:"Password"`
	Author  	string   	`xml:"Author"`
	Salt 		string		`xml:"Salt"`
}

func readUsers() []user {	//Liest alle vorhanden User aus der XML ein
	var users []user																//Neues User Array wird angelegt

		file, err := os.Open("./ressources/users.xml") 						//Datei wird geöffnet
		if err != nil {
			fmt.Printf("error: %v", err)
			return users
		}
		defer file.Close()
		data, err := ioutil.ReadAll(file)											//Datei wird ausgelesen
		if err != nil {
			fmt.Printf("error: %v", err)
			return users
		}
		v := Recurlyusers{}															//Container zum speichern der Werte
		err = xml.Unmarshal(data, &v)												//XML einlesen
		if err != nil {
			fmt.Printf("error: %v", err)
			return users
		}
	return v.Svs																	//Users zurückgeben
}

func passwordHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("./ressources/html/changePW.html")		//Passwort ändern Maske aufrufen
		t.Execute(w, nil)
	}else{
		r.ParseForm()
		PW1 := strings.Join(r.Form["password1"],"")
		PW2 := strings.Join(r.Form["password2"],"")							//Eingegebene Passwörter auslesen
		if PW1 != PW2{																//Prüfen ob diese übereinstimmen
			responseString := 	"<html>"+
				"<body>"+
				"<h1>Programmieren II - Blog</h1><br>"+
				"Passwörter stimmen nicht überein "+"<a href='/home'>Bitte Klicken</a>"+
				"</body>"+
				"</html>"
			w.Write([]byte(responseString))
		}else{																		//Wenn ja, dann
			c,_:= r.Cookie("username")										//Betroffenen Usernamen auslesen
			username = c.Value
			salt := createsalt(PW1)
			hash := createHash(PW1, salt)											//Hash u. Salt generieren

			actualUsers := readUsers()												//Alle vorhandenen User speichen
			v2 := &users{Version: "1"}												//Neue User Referenz anlegen
			i:= 0
			for i < len(actualUsers){												//Alle User durchiterieren
				if actualUsers[i].Name != username{
					v2.Svs = append(v2.Svs, writeuser{actualUsers[i].Name,actualUsers[i].Password,actualUsers[i].Author,actualUsers[i].Salt}) //Alte User u. Passwörrter übernehmen
				}else{
					v2.Svs = append(v2.Svs, writeuser{actualUsers[i].Name,hash,actualUsers[i].Author,salt})			//User mit neuem Passwort ändern
				}
				i++
			}

			output2, err := xml.MarshalIndent(v2, "  ", "    ")  //XML erzeugen
			if err != nil {
				fmt.Printf("error: %v\n", err)
			}

			path := "./ressources/users.xml"
			var err3 = os.Remove(path)												//Altes User File löschen
			if err3 != nil { return }
			var _,err2 = os.Stat(path)
			if os.IsNotExist(err2) {
				var file, err2 = os.Create(path)									//Neue Datei erstellen
				if err2 != nil { return }
				defer file.Close()
			}
			ioutil.WriteFile(path,output2,0777)								//Neue User XML schreiben


			responseString := 	"<html>"+
				"<body>"+
				"<h1>Programmieren II - Blog</h1><br>"+
				"Passwort erfolgreich geändert "+"<a href='/home'>Bitte Klicken</a>"+
				"</body>"+
				"</html>"
			w.Write([]byte(responseString))
		}
	}
}


