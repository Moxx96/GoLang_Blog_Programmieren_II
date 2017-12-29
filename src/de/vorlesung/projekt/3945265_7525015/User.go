package main

import (
	"io/ioutil"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type user struct{
	name string
	pwd string
	isAuthor bool
}

func readUsers(){
	sum := 0
	offset := 3
	dat, err := ioutil.ReadFile("./ressources/users.txt")
	check(err)
	fmt.Print(string(dat)+"\n")
	f, err := os.Open("./ressources/users.txt")
	check(err)

	b0 := make([]byte, 1)
	n0, err := f.Read(b0)
	check(err)
	fmt.Printf("%d bytes: %s\n", n0, string(b0))
	count := int(b0[0])-48

	for sum < count{

		_, err = f.Seek(int64(offset), 0)

		b1 := make([]byte, 1)
		n1, err := f.Read(b1)
		check(err)
		fmt.Printf("%d bytes: %s\n", n1, string(b1))

		namelength := int(b1[0])-48
		_, err = f.Seek(int64(offset+2), 0)
		check(err)

		b2 := make([]byte,1)
		n2, err := f.Read(b2)
		check(err)
		fmt.Printf("%d bytes: %s\n", n2, string(b2))

		passlength := int(b2[0])-48

		_, err = f.Seek(int64(offset+4), 0)

		b3 := make([]byte, namelength)
		n3, err := f.Read(b3)
		check(err)
		fmt.Printf("%d bytes: %s\n", n3, string(b3))

		_, err = f.Seek(int64(offset+5+namelength), 0)

		b4 := make([]byte, passlength)
		n4, err := f.Read(b4)
		check(err)
		fmt.Printf("%d bytes: %s\n", n4, string(b4))

		_, err = f.Seek(int64(offset+6+namelength+passlength), 0)

		b5 := make([]byte, 1)
		n5, err := f.Read(b5)
		check(err)
		fmt.Printf("%d bytes: %s\n\n", n5, string(b5))


		offset = offset+6+namelength+passlength+3
		sum++
	}
	f.Close()

}
