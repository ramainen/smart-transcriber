package main

import (
	"fmt"

	transcriber "github.com/ramainen/smart-transcriber"
)

func main() {

	trans := transcriber.NewTranscriber()
	fmt.Println("apple ->", trans.Transcribe("apple"))
	fmt.Println("makita ->", trans.Transcribe("makita"))
	fmt.Println("bosch ->", trans.Transcribe("bosch"))
	fmt.Println("xiaomi ->", trans.Transcribe("xiaomi"))
	fmt.Println("bergauf ->", trans.Transcribe("bergauf"))
	fmt.Println("knauf ->", trans.Transcribe("knauf"))

	fmt.Println("zelda ->", trans.Transcribe("zelda"))
	fmt.Println("argotech ->", trans.Transcribe("argotech"))
	fmt.Println("rossinka ->", trans.Transcribe("rossinka"))

	fmt.Println("trabadath ->", trans.Transcribe("trabadath"))
	fmt.Println("obs ->", trans.Transcribe("obs"))
	fmt.Println("chs ->", trans.Transcribe("chs"))
	fmt.Println("fbi ->", trans.Transcribe("fbi"))

	fmt.Println("tytan ->", trans.Transcribe("tytan"))
	fmt.Println("qwerty ->", trans.Transcribe("qwerty"))
	fmt.Println("microsoft ->", trans.Transcribe("microsoft"))
	fmt.Println("huawey ->", trans.Transcribe("huawey"))
	fmt.Println("bmw ->", trans.Transcribe("bmw"))

	fmt.Println("wolfenstein ->", trans.Transcribe("wolfenstein"))
	fmt.Println("morgerstern ->", trans.Transcribe("morgerstern"))
	fmt.Println("multfilm ->", trans.Transcribe("multfilm"))
	fmt.Println("bolt ->", trans.Transcribe("bolt"))
	fmt.Println("object ->", trans.Transcribe("object"))
	fmt.Println("avaya ->", trans.Transcribe("avaya"))
	fmt.Println("leroy ->", trans.Transcribe("leroy"))
	fmt.Println("merlin ->", trans.Transcribe("merlin"))
	fmt.Println("matrix  ->", trans.Transcribe("matrix"))
	fmt.Println("philips  ->", trans.Transcribe("philips"))

}
