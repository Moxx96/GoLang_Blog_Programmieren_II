package main
/*
import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"crypto/sha256"
)

func main() {
	pwd := "test"

	hash := md5.New()
	hash.Write([]byte(pwd))
	salt := hex.EncodeToString(hash.Sum(nil))

	fmt.Println(salt)



	salted := pwd + salt
	bash := sha256.New()
	bash.Write([]byte(salted))
	pwdhash := hex.EncodeToString(bash.Sum(nil))

	fmt.Println(pwdhash)
} */
