package configFunc

import (
	"math/rand"
	"time"
)

func goscript() (string, error) {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, 24)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b), nil
}

// {"func":"import (\n\"fmt\"\n\"math/rand\"\n)\nfunc goscript() (string, error) {\nreturn fmt.Sprintf(\"%v\", rand.Intn(3000)), nil\n}"}

// "func goscript() (string, error) {\nvar letters = []rune(\"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ\")\nb := make([]rune, 20)\nfor i := range b {\nb[i] = letters[rand.Intn(len(letters))]\n}\nreturn string(b), nil\n}"
