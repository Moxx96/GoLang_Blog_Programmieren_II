package main
//Matrikelnummern: 3945265, 7525015
import (
	"crypto/md5"
	"encoding/hex"
	"crypto/sha256"
)

func createsalt (pwd string) (string){						//Generiert ein Salt aus dem PW
	//Wird beim erstellen von Accounts oder dem ändern des PW benötigt
	hash := md5.New()
	hash.Write([]byte(pwd))
	salt := hex.EncodeToString(hash.Sum(nil))
	return salt

}

func createHash (pwd string, salt string) (string){			//Generiert Passworthash aus passwort + Salt
	//Wird beim erstellen von Accounts oder dem ändern des PW benötigt
	salted := pwd + salt
	hash := sha256.New()
	hash.Write([]byte(salted))
	pwdhash := hex.EncodeToString(hash.Sum(nil))

	return pwdhash
}
