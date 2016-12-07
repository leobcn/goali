package app

import (
	"app/interfaces/errs"
	"crypto/sha1"
	"encoding/hex"
)

// NewPharse instances new pharse
func NewPharse(pharse string) (*Pharse, error) {
	var p Pharse
	p.Pharse = pharse
	p.Translations = make(map[string]Translate, 0)

	sum, err := sha1sum(pharse)
	p.Sum = sum

	return &p, err
}

type Pharse struct {
	ID           int    `storm:"id,increment"`
	Sum          string `storm:"unique"`
	Pharse       string
	Translations map[string]Translate
}

func (p *Pharse) AddTranslate(l, t string) {
	p.Translations[l] = Translate{t}
}

type Translate struct {
	Translate string
}

func sha1sum(s string) (string, error) {
	h := sha1.New()
	_, err := h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil)), errs.Wrap(err)
}
