package main

import (
	"net/http"
	"html/template"
	"encoding/xml"
	"os"
	"fmt"
	"io/ioutil"
	"strconv"
	"time"
)

type login struct{
	USERNAME string
	MODUS string
}

type Recurlyposts struct {
	XMLName     xml.Name `xml:posts"`
	Version     string   `xml:"version,attr"`
	Svs         []post   `xml:"post"`
}

type post struct {
	XMLName  xml.Name    `xml:"post"`
	TEXT string          `xml:"TEXT"`
	DATUM string         `xml:"DATUM"`
	AUTHOR   string      `xml:"AUTHOR"`
	COMMENT   string     `xml:"COMMENT"`
}



type beitrag struct{
	TEXT string
	DATUM string
	AUTHOR string
	COUNT string
	COMMENTS []comment
}
type comment struct{
	TEXT string
	DATUM string
	AUTHOR string
}




func homeHandler(w http.ResponseWriter, r *http.Request) {
	c3,_:= r.Cookie("timestamp")
	timeint, _ := strconv.ParseInt(c3.Value, 10, 0)
	if time.Unix(timeint, 0).Before(time.Unix(timeint, 0).Add(time.Minute*15)){
		c,_ := r.Cookie("username")
		c2,_:= r.Cookie("isAuthor")
		t := template.New("Page")
		t2 := template.New("Comment")

		var modus string
		if c2.Value == "0"{
			t, _ = template.ParseFiles("./ressources/html/blogAuthor.html")
			modus = "Author"
		}else{
			t, _ = template.ParseFiles("./ressources/html/blogGast.html")
			modus = "Leser"
		}
		p := login{USERNAME: c.Value, MODUS: modus}
		t.Execute(w,p)

		if c2.Value == "0"{
			t, _ = template.ParseFiles("./ressources/html/beitraegeAuthor.html")
		}else{
			t, _ = template.ParseFiles("./ressources/html/beitraegeGast.html")
		}

		t2, _ = template.ParseFiles("./ressources/html/comments.html")


		files,_ := ioutil.ReadDir("./ressources/storage/")
		filecount := len(files)-1
		var j int
		for 0 <= filecount {
			posts := readPosts(filecount)
			m := beitragGen(posts, filecount)
			t.Execute(w, m)
			j = len(m.COMMENTS)-1
			for 0 <= j{
				t2.Execute(w,comment{m.COMMENTS[j].TEXT,m.COMMENTS[j].DATUM,m.COMMENTS[j].AUTHOR})
				j--
			}
			filecount--
		}
	}else{
		responseString := 	"<html>"+
			"<body>"+
			"<h1>Programmieren II - Blog</h1><br>"+
			"Zugang verweigert "+"<a href='/'>Bitte Klicken</a>"+
			"</body>"+
			"</html>"
		w.Write([]byte(responseString))
	}


}

func readPosts(x int) []post{
	var posts []post

	file, err := os.Open("./ressources/storage/"+strconv.Itoa(x)+".xml") // For read access.
	if err != nil {
		fmt.Printf("error: %v", err)
		return nil
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("error: %v", err)
		return nil
	}
	v := Recurlyposts{}
	err = xml.Unmarshal(data, &v)
	if err != nil {
		fmt.Printf("error: %v", err)
		return nil
	}

	//fmt.Println(v)


	posts = v.Svs

	return posts
}

func beitragGen(p []post, x int) beitrag{
	var bei beitrag
	count :=len(p)
	var com comment

	j:=0
	i:=0

	for count > i{
		if p[i].COMMENT == "0"{
			bei.TEXT = p[i].TEXT
			bei.AUTHOR = p[i].AUTHOR
			bei.DATUM = p[i].DATUM
			bei.COUNT = strconv.Itoa(x)
		}else if p[i].COMMENT == "1"{
			com.TEXT = p[i].TEXT
			com.DATUM = p[i].DATUM
			com.AUTHOR = p[i].AUTHOR
			bei.COMMENTS = append(bei.COMMENTS, com)
			j++
		}

		i++
}

	return bei
}