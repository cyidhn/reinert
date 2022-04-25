package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	// 1. Fonction pour importer le document en format iramuteq
	file, err := ioutil.ReadFile("./corpus/all.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(string(file))
	// 2. Traitement du texte importe Traitement de texte par Rainette

	// 4. Retourne les resultats en JSON
}
