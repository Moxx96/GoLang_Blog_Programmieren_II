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
		os.Stdout.Write([]byte(xml.Header))

		os.Stdout.Write(output)

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

