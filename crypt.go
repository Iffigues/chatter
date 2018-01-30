package main

import(
	"golang.org/x/crypto/bcrypt"
)

func crypte(pass string)(io []byte,err error){
	return bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
}

func decrypt(a []byte,b []byte)(vrai bool){
	if err := bcrypt.CompareHashAndPassword(a, b); err != nil {
		return false
	}
	return true
}
