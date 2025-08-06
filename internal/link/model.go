package link

import (
	"go/http/internal/stat"
	"math/rand"

	"gorm.io/gorm"
)

type Link struct {
	gorm.Model
	Url   string      `json:"url"`
	Hash  string      `json:"hash" gorm:"uniqueIndex"`
	Stats []stat.Stat `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func NewLink(url string) *Link {
	link := &Link{
		Url: url,
	}
	link.RegenerateHash()
	return link
}

func (link *Link) RegenerateHash() {
	link.Hash = RandStringRunes(6)
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
		b[i] = letterRunes[(rand.Intn(len(letterRunes)))]
	}

	return string(b)
}
