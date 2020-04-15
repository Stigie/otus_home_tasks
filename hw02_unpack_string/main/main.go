package main

import (
	unt "github.com/Stigie/otus_home_tasks/hw02_unpack_string"
	"log"
)

func main() {
	str, err := unt.Unpack("dab0cvv0o")
	if err != nil {
		log.Println(err)
	}
	log.Println(str)
	//var rune rune
	//fmt.Println(1,string(rune),1)
}
