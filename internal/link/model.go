package link

import (
	"math/rand"

	"gorm.io/gorm"
)

type Link struct {
	gorm.Model
	Url  string `json:"url"`
	Hash string `json:"hash" gorm:"uniqueIndex"`
}

func NewLink(url string) *Link {
	return &Link{
		Url:  url,
		Hash: RandStringRunes(6),
	}
}

var letterRunes []rune

func initializeRunes() {
	for r := 'a'; r <= 'z'; r++ {
		letterRunes = append(letterRunes, r)
	}
	for r := 'A'; r <= 'Z'; r++ {
		letterRunes = append(letterRunes, r)
	}
}

func init() {
	initializeRunes()
}

func RandStringRunes(n int) string {
	b := make([]rune, n)

	for i := range b {
		b[i] = rune(rand.Intn(len(letterRunes)))
	}

	return string(b)
}
