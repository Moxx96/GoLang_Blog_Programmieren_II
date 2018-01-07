package main
//Matrikelnummern: 3945265, 7525015
import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSalt(t *testing.T) {
	pwd:= "test"
	salt:= createsalt(pwd)
	assert.Equal(t,salt,"098f6bcd4621d373cade4e832627b4f6")
}

func TestHash(t *testing.T) {
	pwd:= "test"
	salt:= "098f6bcd4621d373cade4e832627b4f6"
	hash:= createHash(pwd, salt)
	assert.Equal(t,hash,"1ff557be460525377cf810a16766d4a5825446ef831549c61c77a0f01069c37b")
}