package main
//Matrikelnummern: 3945265, 7525015
import (
	"log"
	"net/http"
	"os"
	"flag"
)

var timeout int64

func main() {
	portPtr := flag.String("port","4443","define port")
	timeoutPtr := flag.Int64("timeout",15,"timeout in minutes")
	flag.Parse()
	timeout = *timeoutPtr
	if _, err := os.Stat("./ressources/storage"); os.IsNotExist(err) {
		os.Mkdir("./ressources/storage", os.ModePerm)
	}
	http.HandleFunc("/", loginHandler) 		 	 	//Öffnet Login Page  (Login.go)
	http.HandleFunc("/guest/", guestHandler) 		//Setzt Gast Cookie und leitet auf home weiter (Login.go)
	http.HandleFunc("/home/", homeHandler)   		//Hier werden alle Beiträge angezeigt (Homepage.go)
	http.HandleFunc("/logout/",logoutHandler)		//Entfernt alle gültigen Cookies und leitet auf Login Page (Login.go)
	http.HandleFunc("/create/",createHandler)		//Neuen Beitrag erstellen (Posts.go)
	http.HandleFunc("/comment/",commentHandler) 		//Einen Beitrag Kommentieren (Posts.go)
	http.HandleFunc("/edit/",editHandler)			//Einen Beitrag berabeiten (Posts.go)
	http.HandleFunc("/delete/",deleteHandler)		//Einen Beitrag löschen (Posts.go)
	http.HandleFunc("/changePW/",passwordHandler)	//Passwort ändern (User.go)
	log.Fatalln(http.ListenAndServeTLS(":"+*portPtr,"./ressources/certBlog.pem" ,"./ressources/keyBlog.pem",nil))  //Server mit HTTPS starten
}