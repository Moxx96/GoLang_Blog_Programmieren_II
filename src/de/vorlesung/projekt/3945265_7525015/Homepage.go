package main

import (
	"net/http"
	"html/template"
	"os"
	"fmt"
	"io/ioutil"
	"strconv"
)

type login struct{
	USERNAME string
	MODUS string
}

type beitrag struct{
	TEXT string
	DATUM string
	AUTHOR string
	COMMENTS []comment
}
type comment struct{
	TEXT string
	DATUM string
	AUTHOR string
}




func homeHandler(w http.ResponseWriter, r *http.Request) {
	c,_ := r.Cookie("username")
	c2,_:= r.Cookie("isAuthor")
	t := template.New("Test")
	t, _ = template.ParseFiles("./ressources/html/blog.html")
	var modus string
	if c2.Value == "0"{
		modus = "Author"
	}else{
		modus = "Leser"
	}
	p := login{USERNAME: c.Value, MODUS: modus}
	t.Execute(w,p)

	t, _ = template.ParseFiles("./ressources/html/beitraege.html")

	var post beitrag

	files,_ := ioutil.ReadDir("./ressources/storage/")
	filecount := len(files)
	i:= 0

	for i < filecount{
		dat, err := ioutil.ReadFile("./ressources/storage/"+strconv.Itoa(i)+".txt")
		check(err)
		fmt.Print(string(dat) + "\n")
		f, err := os.Open("./ressources/storage/0.txt")
		check(err)
		b1 := make([]byte, 1000)
		n1, err := f.Read(b1)
		check(err)
		fmt.Printf("%d bytes: %s\n", n1, string(b1))
		post.TEXT = string(b1)

		_, err = f.Seek(int64(1002), 0)


		b2 := make([]byte, 10)
		n2, err := f.Read(b2)
		check(err)
		fmt.Printf("%d bytes: %s\n", n2, string(b2))
		post.DATUM = string(b2)

		_, err = f.Seek(int64(1014), 0)

		b3 := make([]byte, 1)
		n3, err := f.Read(b3)
		check(err)
		fmt.Printf("%d bytes: %s\n", n3, string(b3))

		namelength := int(b3[0]) - 48

		_, err = f.Seek(int64(1016), 0)

		b4 := make([]byte, namelength)
		n4, err := f.Read(b4)
		check(err)
		fmt.Printf("%d bytes: %s\n", n4, string(b4))
		post.AUTHOR = string(b4)

		_, err = f.Seek(int64(1016+namelength+2), 0)

		b0 := make([]byte, 1)
		n0, err := f.Read(b0)
		check(err)
		fmt.Printf("%d bytes: %s\n", n0, string(b0))

		count:=int(b0[0]) - 48
		sum:=0
		offset:=0

		var comments []comment
		for sum < count {
			var com comment
			_, err = f.Seek(int64(1016+namelength+5+offset), 0)

			b5 := make([]byte, 140)
			n5, err := f.Read(b5)
			check(err)
			fmt.Printf("%d bytes: %s\n", n5, string(b5))
			com.TEXT = string(b5)

			_, err = f.Seek(int64(1156+namelength+7+offset), 0)

			b6 := make([]byte, 10)
			n6, err := f.Read(b6)
			check(err)
			fmt.Printf("%d bytes: %s\n", n6, string(b6))
			com.DATUM = string(b6)
			_, err = f.Seek(int64(1166+namelength+9+offset), 0)

			b7 := make([]byte, 1)
			n7, err := f.Read(b7)
			check(err)
			fmt.Printf("%d bytes: %s\n", n7, string(b7))

			commentatorlength := int(b7[0]) - 48

			_, err = f.Seek(int64(1166+namelength+11+offset), 0)

			b8 := make([]byte, commentatorlength)
			n8, err := f.Read(b8)
			check(err)
			fmt.Printf("%d bytes: %s\n", n8, string(b8))
			com.AUTHOR = string(b8)
			comments = append(comments, com)
			sum++
			offset = offset + commentatorlength +140 +10 +8
		}

		post.COMMENTS = comments

		f.Close()



		m := beitrag{TEXT: post.TEXT,
			DATUM: post.DATUM,
			AUTHOR: post.AUTHOR,
			COMMENTS: post.COMMENTS}
		t.Execute(w,m)
		i++
	}


}