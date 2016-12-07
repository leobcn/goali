package main

import (
	"log"
	"reflect"
	"strings"
	"testing"
)

func AlĞĞĞĞĞ() {
	strings.Contains("a")
	log.Println("selam")
}

func TestUTF(t *testing.T) {
	log.Println(2<<2, 1)
}

func TestString(t *testing.T) {
	adi := "Oğuzhan"
	// adi := "Ali"
	log.Printf("%q", adi[0])
	log.Printf("% x", adi[1])
	log.Printf("%+b", adi)
	log.Println(reflect.TypeOf(adi[0]))
}
