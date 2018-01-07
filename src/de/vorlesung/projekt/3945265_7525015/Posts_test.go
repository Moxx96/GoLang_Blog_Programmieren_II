package main
//Matrikelnummern: 3945265, 7525015
import (
	"testing"
	"github.com/stretchr/testify/assert"
	"encoding/xml"
	"strconv"
)

func TestBeitrag(t *testing.T) {
	var p []post
	i:=0

	var xmNull xml.Name

	p = append(p, post{xmNull,"Der Testify Text!", "01.01.2001","Der Testify Author!","0"})
	p = append(p, post{xmNull,"Der Testify Kommentar 1!", "01.01.2002","Der Testify Kommentator!","1"})
	p = append(p, post{xmNull,"Der Testify Kommentar 2!", "01.01.2003", "Der Testify Kommentator Nummer 2!","1"})

	bei := beitragGen(p,i)

	assert.Equal(t,bei.TEXT,p[0].TEXT)
	assert.Equal(t,bei.AUTHOR,p[0].AUTHOR)
	assert.Equal(t,bei.DATUM,p[0].DATUM)
	assert.Equal(t,bei.COUNT,strconv.Itoa(i))

	assert.Equal(t,bei.COMMENTS[0].TEXT,p[1].TEXT)
	assert.Equal(t,bei.COMMENTS[0].AUTHOR,p[1].AUTHOR)
	assert.Equal(t,bei.COMMENTS[0].DATUM,p[1].DATUM)

	assert.Equal(t,bei.COMMENTS[1].TEXT,p[2].TEXT)
	assert.Equal(t,bei.COMMENTS[1].AUTHOR,p[2].AUTHOR)
	assert.Equal(t,bei.COMMENTS[1].DATUM,p[2].DATUM)

}