package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// 1. Fonction pour importer le document en format iramuteq
	file, err := ioutil.ReadFile("./corpus/new_file.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// 2. Traitement du texte importe Traitement de texte par Rainette
	fmt.Print(string(output))
	// 4. Retourne les resultats en JSON
}
