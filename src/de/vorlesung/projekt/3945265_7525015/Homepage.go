package main
//Matrikelnummern: 3945265, 7525015
import (
	"net/http"
	"html/template"
	"io/ioutil"
)

type login struct{
	USERNAME string
	MODUS string
}


func homeHandler(w http.ResponseWriter, r *http.Request) {
	c,_ := r.Cookie("username")						//Entsprechende Cookies werden überprüft
	c2,_:= r.Cookie("isAuthor")
	t := template.New("Page")
	t2 := template.New("Comment")

		var modus string
		if c2.Value == "0"{
			t, _ = template.ParseFiles("./ressources/html/blogAuthor.html") 	//Ist das Author flag gesetzt wird die entprechende Seite geladen
			modus = "Author"
		}else{
			t, _ = template.ParseFiles("./ressources/html/blogGast.html")		//Ansonsten Gast Seite
			modus = "Leser"
		}
		p := login{USERNAME: c.Value, MODUS: modus}										//Dynamische Werte werden als Struct gespeichert und als Template geladen
		t.Execute(w,p)

		if c2.Value == "0"{
			t, _ = template.ParseFiles("./ressources/html/beitraegeAuthor.html")		//Ebenso beim Laden der Beiträge, hier sehen Autoren mehr Buttons
		}else{
			t, _ = template.ParseFiles("./ressources/html/beitraegeGast.html")			//Gäste sehen nur den Kommentarbutton (siehe html File)
		}

		t2, _ = template.ParseFiles("./ressources/html/comments.html")					//Kommentare werden für alle gleich geladen

		files,_ := ioutil.ReadDir("./ressources/storage/")							//Es wird gelesen wie viele Beiträge vorhanden sind
		filecount := len(files)-1
		var j int
		for 0 <= filecount {																//Für die Anzalh dieser wird durchiteriert
			posts := readPosts(filecount)													//Post wird eingelesen
			m := beitragGen(posts, filecount)
			if m.AUTHOR == c.Value{
				t, _ = template.ParseFiles("./ressources/html/beitraegeAuthor.html")
			}else {
				t, _ = template.ParseFiles("./ressources/html/beitraegeGast.html")
			}																				//In einen template Kompatiblen Struct konvertiert
			t.Execute(w, m)																	//Und ausgegeben
			j = len(m.COMMENTS)-1															//Anzahl der Kommentare auslesen
			for 0 <= j{
				t2.Execute(w,comment{m.COMMENTS[j].TEXT,m.COMMENTS[j].DATUM,m.COMMENTS[j].AUTHOR})	//Kommentare ausgeben
				j--
			}
			filecount--
		}



}

