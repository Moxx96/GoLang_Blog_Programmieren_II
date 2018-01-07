package main
//Matrikelnummern: 3945265, 7525015
import (
	"testing"
	"github.com/stretchr/testify/assert"
	"encoding/xml"
	"crypto/sha256"
	"encoding/hex"
)

func TestRead(t *testing.T) {
	username:= "TestUser"
	password:= "Password123"

	Users:= readUsers()

	var xmNull xml.Name
	compareUser := user{xmNull,"","","", ""}	//Leeren User zum vergleichen erstellen
	validUser := compareUser														//Variable für einen gültigen User erstellen
	for _, element := range Users{							//Alle User durchiterieren
		if element.Name == username{						//Bei übereinstimmendem Namen
			salted := password + element.Salt				//Berechnen des Hashes von Passwort + Salt
			hash := sha256.New()							//
			hash.Write([]byte(salted)) 						//
			pwdhash := hex.EncodeToString(hash.Sum(nil))	//Konvertieren von Typ byte zu string
			if element.Password == pwdhash{					//Bei übereinstimmenden Passwort
				validUser = element							//Als gültigen Usersetzen
			}
		}
	}
	assert.Equal(t,validUser.Name, username)
	assert.Equal(t,validUser.Password,createHash(password,createsalt(password)))
}

