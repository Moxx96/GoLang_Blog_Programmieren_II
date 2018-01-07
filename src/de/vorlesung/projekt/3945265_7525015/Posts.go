package main
//Matrikelnummern: 3945265, 7525015
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




func createHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("method:", r.Method)
	if r.Method == "GET" {																	//Prüfen auf GET-Anfrage
		t, _ := template.ParseFiles("./ressources/html/createBeitrag.html")
		t.Execute(w, nil)																//Html Maske ausgeben
	} else {																				//Bei Post Anfrage (also beim zweiten Aufruf)
		r.ParseForm()
		text := strings.Join(r.Form["post1"], "")										//Eingegebenen Text speichern
		c, _ := r.Cookie("username")													//Username aus Cookie lesen
		d := string(time.Now().Format("02.01.2006"))									//Aktuelles Datum als String speichern


		v := &posts{Version: "1"}															//Neue Post Referenz anlegen
		v.Svs = append(v.Svs, writepost{text,d,c.Value,"0"})//Daten Anhängen

		output, err := xml.MarshalIndent(v, "  ", "    ")						//XML erzeugen
		if err != nil {
			fmt.Printf("error: %v\n", err)
		}

		files,_ := ioutil.ReadDir("./ressources/storage/")
		filecount := len(files)																//Anzahl der Beiträge auslesen
		path := "./ressources/storage/"+strconv.Itoa(filecount)+".xml"						//Pfad für Beiträge+1 anlegen

		var _,err2 = os.Stat(path)
		if os.IsNotExist(err2) {
			var file, err2 = os.Create(path)												//Neue datei erstellen
			if err2 != nil { return }
			defer file.Close()
		}

		ioutil.WriteFile(path,output,0777)											//Inhalt in XML Datei schreiben
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
	c2,_:= r.Cookie("isAuthor")												//Author Cookie auslesen

	if r.Method == "GET" {
		q := r.URL.Query()
		count := q.Get("count")													//count auslesen (Zuordung des Beitrags)
		cookie := http.Cookie{Name: "count", Value: count, Path: "/comment/"}
		http.SetCookie(w, &cookie)													//cookie mit count wert setzen
		t := template.New("Edit")
		if c2.Value == "0"{
			t, _ = template.ParseFiles("./ressources/html/commentAuthor.html")		//Author Kommentarseite (ohne Namensfeld)
		}else{
			t, _ = template.ParseFiles("./ressources/html/commentGast.html")		//Gast Kommentarseit (mit Namensfeld)
		}
		t.Execute(w,nil)														//Anzeigen
	} else {
		cc,_:= r.Cookie("count")											//Zuordnungscookie auslesen
		count,_ := strconv.Atoi(cc.Value)										//in INT wandeln
		r.ParseForm()
		text := strings.Join(r.Form["post2"], "")							//Eingegeben Text auslesen
		d := string(time.Now().Format("02.01.2006"))						//Aktuelles Datum auslesen
		var name string
		if c2.Value == "0"{														//Entweder Authornamen
			c, _ := r.Cookie("username")
			name= c.Value
		}else{																	//Oder eingebenen Gastnamen auslesen
			name = strings.Join(r.Form["username"], "")
		}
		v:= beitragGen(readPosts(count),count)									//Betroffenen Beitrag laden
		v2 := &posts{Version: "1"}												//Neue Referenz für Beitrag anlegen
		l := len(v.COMMENTS)													//Anzalh der Kommentare lesen
		sum := 0
		v2.Svs = append(v2.Svs, writepost{v.TEXT,v.DATUM,v.AUTHOR,"0"})		//Ursprücglichen Beitrag anhängen
		for sum < l{
			v2.Svs = append(v2.Svs, writepost{v.COMMENTS[sum].TEXT,v.COMMENTS[sum].DATUM,v.COMMENTS[sum].AUTHOR,"1"})	//Ursprüngliche Kommentare anhängen
			sum++
		}
		v2.Svs = append(v2.Svs,writepost{text,d,name,"1"})	//Neuen Kommentar anhängen
		output2, err := xml.MarshalIndent(v2, "  ", "    ")					//XML generieren
		if err != nil {
			fmt.Printf("error: %v\n", err)
		}
		//os.Stdout.Write([]byte(xml.Header))
		//os.Stdout.Write(output2)

		path := "./ressources/storage/"+strconv.Itoa(count)+".xml"
		var err3 = os.Remove(path)															//Alte XML Datei löschen
		if err3 != nil { return }
		var _,err2 = os.Stat(path)
		if os.IsNotExist(err2) {
			var file, err2 = os.Create(path)												//Neue Datei erstellen (gleicher Name)
			if err2 != nil { return }
			defer file.Close()
		}
		ioutil.WriteFile(path,output2,0777)											//Werte in XML schreiben

		responseString := 	"<html>"+
			"<body>"+
			"<h1>Programmieren II - Blog</h1><br>"+
			"Kommentar erfolgreich erstellt "+"<a href='/home'>Bitte Klicken</a>"+
			"</body>"+
			"</html>"
		w.Write([]byte(responseString))

	}
}




func deleteHandler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	count,_:= strconv.Atoi(q.Get("count"))													//Zuordnung auslesen

	files,_ := ioutil.ReadDir("./ressources/storage/")
	filecount := len(files)																		//Anzahl Beiträge auslesen

	for count < filecount-1{
		v:= beitragGen(readPosts(count+1),count+1)												//Gesählten Beitrag+1 laden
		v2 := &posts{Version: "1"}																//Neue Beitragsferenz

		l := len(v.COMMENTS)																	//Anzahl der Kommentare
		sum := 0
		v2.Svs = append(v2.Svs, writepost{v.TEXT,v.DATUM,v.AUTHOR,"0"}) //Beitragstext anhängen
		for sum < l{
			v2.Svs = append(v2.Svs, writepost{v.COMMENTS[sum].TEXT,v.COMMENTS[sum].DATUM,v.COMMENTS[sum].AUTHOR,"1"})	//Kommentare anhängen
			sum++
		}
		output2, err := xml.MarshalIndent(v2, "  ", "    ")					//XML erzeugen
		if err != nil {
			fmt.Printf("error: %v\n", err)
		}
		path := "./ressources/storage/"+strconv.Itoa(count)+".xml"							//Datei löschen
		var err3 = os.Remove(path)
		if err3 != nil { return }
		var _,err2 = os.Stat(path)
		if os.IsNotExist(err2) {
			var file, err2 = os.Create(path)												//Datei erzeugen
			if err2 != nil { return }
			defer file.Close()
		}
		ioutil.WriteFile(path,output2,0777)											//Werte aus dem höheren Beitrag in den zu löschenden schreiben
		count++																				//Solange bis das letzte file erreicht ist
	}

	path := "./ressources/storage/"+strconv.Itoa(filecount-1)+".xml"
	var err3 = os.Remove(path)																//letztes XML File wird gelöscht, da insgesamt einer weniger
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
	if r.Method == "GET" {																	//Prüfen ob GET (1st Aufruf)
		q := r.URL.Query()
		count := q.Get("count")
		cookie := http.Cookie{Name: "count", Value: count, Path: "/edit/"}					//Zuordung lesen und als Cookie speichern
		http.SetCookie(w, &cookie)
		t := template.New("Edit")
		type preview struct{																//Struct für das template erstellen
			TEXT string
		}
		countINT,_:= strconv.Atoi(count)													//Zuordung in INT
		v:= beitragGen(readPosts(countINT),countINT)										//Betroffenen beitrag einlesen
		var pre preview
		pre.TEXT = v.TEXT
		t, _ = template.ParseFiles("./ressources/html/editBeitrag.html")
		t.Execute(w,pre)																	//Beitrag wird ins Textfeld vorgeladen
	} else {
		cc,_:= r.Cookie("count")														//Zuordnungscookie wird ausgelesen
		count,_ := strconv.Atoi(cc.Value)													//in INT konvertiert
		r.ParseForm()
		text := strings.Join(r.Form["post3"], "")										//veränderter Text wird ausgelesen

		v:= beitragGen(readPosts(count),count)												//Betroffener Beitrag geladen
		v2 := &posts{Version: "1"}															//Neue Referenz für Beitrag erstellt

		l := len(v.COMMENTS)																//Anzahl der Kommentare
		sum := 0
		v2.Svs = append(v2.Svs, writepost{text,v.DATUM,v.AUTHOR,"0"})	//Neue Beitragstext wird angehängt
		for sum < l{
			v2.Svs = append(v2.Svs, writepost{v.COMMENTS[sum].TEXT,v.COMMENTS[sum].DATUM,v.COMMENTS[sum].AUTHOR,"1"})	//Kommentare werden angehängt
			sum++
		}
		output2, err := xml.MarshalIndent(v2, "  ", "    ")	//XML wird erzeugt
		if err != nil {
			fmt.Printf("error: %v\n", err)
		}
		//os.Stdout.Write([]byte(xml.Header))
		//os.Stdout.Write(output2)

		path := "./ressources/storage/"+strconv.Itoa(count)+".xml"
		var err3 = os.Remove(path)												//Alte Datei gelöscht
		if err3 != nil { return }
		var _,err2 = os.Stat(path)
		if os.IsNotExist(err2) {
			var file, err2 = os.Create(path)									//Neue Datei (gleicher Name) erstellt
			if err2 != nil { return }
			defer file.Close()
		}
		ioutil.WriteFile(path,output2,0777)								//XML wird in Datei geschrieben

		responseString := 	"<html>"+
			"<body>"+
			"<h1>Programmieren II - Blog</h1><br>"+
			"Beitrag erfolgreich bearbeitet "+"<a href='/home'>Bitte Klicken</a>"+
			"</body>"+
			"</html>"
		w.Write([]byte(responseString))

	}
}


func readPosts(x int) []post{					//Dient dazu XML mit EINEM Beitrag einzulesen
	var posts []post

	file, err := os.Open("./ressources/storage/"+strconv.Itoa(x)+".xml") //Öffnet Datei, x bestimmt die Zuordnung
	if err != nil {
		fmt.Printf("error: %v", err)
		return nil
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)											//Ganzer Inhalt wird geladen
	if err != nil {
		fmt.Printf("error: %v", err)
		return nil
	}
	v := Recurlyposts{}															//Container zum speichern der Werte
	err = xml.Unmarshal(data, &v)												//XML wird eingelesen
	if err != nil {
		fmt.Printf("error: %v", err)
		return nil
	}

	//fmt.Println(v)


	posts = v.Svs

	return posts																//Rückgabe des Beitrags (ohne Kommentardifferenzierung)
}

func beitragGen(p []post, x int) beitrag{										//Konvertiert einen eingelesenen Beitrag in einen verwendbaren Beitrag
	var bei beitrag																//Beitragsvariable
	count :=len(p)																//Anzahl der Elemente
	var com comment																//Neue Kommentarvariable

	j:=0
	i:=0

	for count > i{																//Iteriert durch Posts
		if p[i].COMMENT == "0"{													//Ist es kein Kommentar wird es als Beitrag hinzugefügt
			bei.TEXT = p[i].TEXT
			bei.AUTHOR = p[i].AUTHOR
			bei.DATUM = p[i].DATUM
			bei.COUNT = strconv.Itoa(x)
		}else if p[i].COMMENT == "1"{											//Steht das Kommentarflag, wird ein Kommentar angehängt
			com.TEXT = p[i].TEXT
			com.DATUM = p[i].DATUM
			com.AUTHOR = p[i].AUTHOR
			bei.COMMENTS = append(bei.COMMENTS, com)
			j++
		}

		i++
	}

	return bei																//Der fertige Beitrag wird zurückgegeben
}

