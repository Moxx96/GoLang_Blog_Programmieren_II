package main

import (
	"strings"
	"time"
	"encoding/xml"
	"fmt"
	"os"
	"net/http"
	"html/template"
	"io/ioutil"
	"strconv"
)

type posts struct {
	XMLName     xml.Name `xml:posts"`
	Version     string   `xml:"version,attr"`
	Svs         []writepost   `xml:"post"`
}

type writepost struct {
	TEXT string          `xml:"TEXT"`
	DATUM string         `xml:"DATUM"`
	AUTHOR   string      `xml:"AUTHOR"`
	COMMENT   string     `xml:"COMMENT"`
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



func createHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("./ressources/html/createBeitrag.html")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		text := strings.Join(r.Form["post1"], "")
		c, _ := r.Cookie("username")
		d := string(time.Now().Format("01.02.2006"))


		v := &posts{Version: "1"}
		v.Svs = append(v.Svs, writepost{text,d,c.Value,"0"})

		output, err := xml.MarshalIndent(v, "  ", "    ")
		if err != nil {
			fmt.Printf("error: %v\n", err)
		}
		//os.Stdout.Write([]byte(xml.Header))
		//os.Stdout.Write(output)

		files,_ := ioutil.ReadDir("./ressources/storage/")
		filecount := len(files)
		path := "./ressources/storage/"+strconv.Itoa(filecount)+".xml"

		var _,err2 = os.Stat(path)
		if os.IsNotExist(err2) {
			var file, err2 = os.Create(path)
			if err2 != nil { return }
			defer file.Close()
		}

		ioutil.WriteFile(path,output,0777)
		//fmt.Print(err)


		responseString := 	"<html>"+
			"<body>"+
			"<h1>Programmieren II - Blog</h1><br>"+
			"Beitrag erfolgreich erstellt "+"<a href='/home'>Bitte Klicken</a>"+
			"</body>"+
			"</html>"
		w.Write([]byte(responseString))

	}
}

func commentHandler(w http.ResponseWriter, r *http.Request) {
	c2,_:= r.Cookie("isAuthor")

	if r.Method == "GET" {
		q := r.URL.Query()
		count := q.Get("count")
		cookie := http.Cookie{Name: "count", Value: count, Path: "/comment/"}
		http.SetCookie(w, &cookie)
		t := template.New("Edit")
		if c2.Value == "0"{
			t, _ = template.ParseFiles("./ressources/html/commentAuthor.html")
		}else{
			t, _ = template.ParseFiles("./ressources/html/commentGast.html")
		}
		t.Execute(w,nil)
	} else {
		cc,_:= r.Cookie("count")
		count,_ := strconv.Atoi(cc.Value)
		r.ParseForm()
		text := strings.Join(r.Form["post2"], "")
		d := string(time.Now().Format("01.02.2006"))
		var name string
		if c2.Value == "0"{
			c, _ := r.Cookie("username")
			name= c.Value
		}else{
			name = strings.Join(r.Form["username"], "")
		}
		v:= beitragGen(readPosts(count),count)
		v2 := &posts{Version: "1"}
		l := len(v.COMMENTS)
		sum := 0
		v2.Svs = append(v2.Svs, writepost{v.TEXT,v.DATUM,v.AUTHOR,"0"})
		for sum < l{
			v2.Svs = append(v2.Svs, writepost{v.COMMENTS[sum].TEXT,v.COMMENTS[sum].DATUM,v.COMMENTS[sum].AUTHOR,"1"})
			sum++
		}
		v2.Svs = append(v2.Svs,writepost{text,d,name,"1"})
		output2, err := xml.MarshalIndent(v2, "  ", "    ")
		if err != nil {
			fmt.Printf("error: %v\n", err)
		}
		//os.Stdout.Write([]byte(xml.Header))
		//os.Stdout.Write(output2)

		path := "./ressources/storage/"+strconv.Itoa(count)+".xml"
		var err3 = os.Remove(path)
		if err3 != nil { return }
		var _,err2 = os.Stat(path)
		if os.IsNotExist(err2) {
			var file, err2 = os.Create(path)
			if err2 != nil { return }
			defer file.Close()
		}
		ioutil.WriteFile(path,output2,0777)

		responseString := 	"<html>"+
			"<body>"+
			"<h1>Programmieren II - Blog</h1><br>"+
			"Kommentar erfolgreich erstellt "+"<a href='/home'>Bitte Klicken</a>"+
			"</body>"+
			"</html>"
		w.Write([]byte(responseString))

	}
}


func passwordHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("./ressources/html/changePW.html")
		t.Execute(w, nil)
	}else{
		r.ParseForm()
		PW1 := strings.Join(r.Form["password1"],"")
		PW2 := strings.Join(r.Form["password2"],"")
		if PW1 != PW2{
			responseString := 	"<html>"+
				"<body>"+
				"<h1>Programmieren II - Blog</h1><br>"+
				"Passwörter stimmen nicht überein "+"<a href='/home'>Bitte Klicken</a>"+
				"</body>"+
				"</html>"
			w.Write([]byte(responseString))
		}else{
			c,_:= r.Cookie("username")
			username = c.Value
			salt := createsalt(PW1)
			hash := createHash(PW1, salt)

			actualUsers := readUsers()
			v2 := &users{Version: "1"}
			i:= 0
			for i < len(actualUsers){
				if actualUsers[i].Name != username{
					v2.Svs = append(v2.Svs, writeuser{actualUsers[i].Name,actualUsers[i].Password,actualUsers[i].Author,actualUsers[i].Salt})
				}else{
					v2.Svs = append(v2.Svs, writeuser{actualUsers[i].Name,hash,actualUsers[i].Author,salt})
				}
				i++
			}

			output2, err := xml.MarshalIndent(v2, "  ", "    ")
			if err != nil {
				fmt.Printf("error: %v\n", err)
			}
			//os.Stdout.Write([]byte(xml.Header))
			//os.Stdout.Write(output2)

			path := "./ressources/users.xml"
			var err3 = os.Remove(path)
			if err3 != nil { return }
			var _,err2 = os.Stat(path)
			if os.IsNotExist(err2) {
				var file, err2 = os.Create(path)
				if err2 != nil { return }
				defer file.Close()
			}
			ioutil.WriteFile(path,output2,0777)


			//Hier muss das Passwort für den User aus dem Cookie in der users.xml geändert werden

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

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	count,_:= strconv.Atoi(q.Get("count"))

	files,_ := ioutil.ReadDir("./ressources/storage/")
	filecount := len(files)

	for count < filecount-1{
		v:= beitragGen(readPosts(count+1),count+1)
		v2 := &posts{Version: "1"}

		l := len(v.COMMENTS)
		sum := 0
		v2.Svs = append(v2.Svs, writepost{v.TEXT,v.DATUM,v.AUTHOR,"0"})
		for sum < l{
			v2.Svs = append(v2.Svs, writepost{v.COMMENTS[sum].TEXT,v.COMMENTS[sum].DATUM,v.COMMENTS[sum].AUTHOR,"1"})
			sum++
		}
		output2, err := xml.MarshalIndent(v2, "  ", "    ")
		if err != nil {
			fmt.Printf("error: %v\n", err)
		}
		path := "./ressources/storage/"+strconv.Itoa(count)+".xml"
		var err3 = os.Remove(path)
		if err3 != nil { return }
		var _,err2 = os.Stat(path)
		if os.IsNotExist(err2) {
			var file, err2 = os.Create(path)
			if err2 != nil { return }
			defer file.Close()
		}
		ioutil.WriteFile(path,output2,0777)
		count++
	}

	path := "./ressources/storage/"+strconv.Itoa(filecount-1)+".xml"
	var err3 = os.Remove(path)
	if err3 != nil { return }

	responseString := 	"<html>"+
		"<body>"+
		"<h1>Programmieren II - Blog</h1><br>"+
		"Beitrag erfolgreich gelöscht "+"<a href='/home'>Bitte Klicken</a>"+
		"</body>"+
		"</html>"
	w.Write([]byte(responseString))



}



func editHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		q := r.URL.Query()
		count := q.Get("count")
		cookie := http.Cookie{Name: "count", Value: count, Path: "/edit/"}
		http.SetCookie(w, &cookie)
		t := template.New("Edit")
		type preview struct{
			TEXT string
		}
		countINT,_:= strconv.Atoi(count)
		v:= beitragGen(readPosts(countINT),countINT)
		var pre preview
		pre.TEXT = v.TEXT
		t, _ = template.ParseFiles("./ressources/html/editBeitrag.html")
		t.Execute(w,pre)
	} else {
		//Edit magic
		cc,_:= r.Cookie("count")
		count,_ := strconv.Atoi(cc.Value)
		r.ParseForm()
		text := strings.Join(r.Form["post3"], "")

		v:= beitragGen(readPosts(count),count)
		v2 := &posts{Version: "1"}

		l := len(v.COMMENTS)
		sum := 0
		v2.Svs = append(v2.Svs, writepost{text,v.DATUM,v.AUTHOR,"0"})
		for sum < l{
			v2.Svs = append(v2.Svs, writepost{v.COMMENTS[sum].TEXT,v.COMMENTS[sum].DATUM,v.COMMENTS[sum].AUTHOR,"1"})
			sum++
		}
		output2, err := xml.MarshalIndent(v2, "  ", "    ")
		if err != nil {
			fmt.Printf("error: %v\n", err)
		}
		//os.Stdout.Write([]byte(xml.Header))
		//os.Stdout.Write(output2)

		path := "./ressources/storage/"+strconv.Itoa(count)+".xml"
		var err3 = os.Remove(path)
		if err3 != nil { return }
		var _,err2 = os.Stat(path)
		if os.IsNotExist(err2) {
			var file, err2 = os.Create(path)
			if err2 != nil { return }
			defer file.Close()
		}
		ioutil.WriteFile(path,output2,0777)

		responseString := 	"<html>"+
			"<body>"+
			"<h1>Programmieren II - Blog</h1><br>"+
			"Beitrag erfolgreich bearbeitet "+"<a href='/home'>Bitte Klicken</a>"+
			"</body>"+
			"</html>"
		w.Write([]byte(responseString))

	}
}




