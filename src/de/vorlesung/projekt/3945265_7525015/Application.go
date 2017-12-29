package main

import (
	"log"
	"net/http"
)

func mainHandler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	name := q.Get("name")
	if name == "" {
		name = "World"
	}
	responseString := 	"<html>"+
							"<body>"+
								"<form>"+
									"<label for='user'>Username:"+
									"<input id='user' name='user'>"+
									"</label><br>"+
									"<label for='pwd'>Passwort:"+
									"<input type='Password' id='pwd' name='pwd'>"+
									"</label><br>"+
									"<input type='submit' value='Login'>"+
								"</form>"+
							"</body>"+
						"</html>"
	w.Write([]byte(responseString))
}

func main() {
	readUsers()
	http.HandleFunc("/", mainHandler)
	log.Fatalln(http.ListenAndServeTLS(":4443","./ressources/certBlog.pem" ,"./ressources/keyBlog.pem",nil))
}