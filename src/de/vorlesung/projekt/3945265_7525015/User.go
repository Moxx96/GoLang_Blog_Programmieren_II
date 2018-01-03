package main

import (
	"io/ioutil"
	"fmt"
	"os"
	"encoding/xml"
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

func readUsers() []user {
	var users []user

		file, err := os.Open("./ressources/users.xml") // For read access.
		if err != nil {
			fmt.Printf("error: %v", err)
			return users
		}
		defer file.Close()
		data, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Printf("error: %v", err)
			return users
		}
		v := Recurlyusers{}
		err = xml.Unmarshal(data, &v)
		if err != nil {
			fmt.Printf("error: %v", err)
			return users
		}

		fmt.Println(v)

	return v.Svs
}


