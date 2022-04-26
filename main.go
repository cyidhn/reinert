package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// 1. Fonction pour importer le document en format iramuteq
	file, err := ioutil.ReadFile("./corpus/all.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// 2. Traitement du texte importe Traitement de texte par Rainette
	output := bytes.Replace(file, []byte("*"), []byte(" "), -1)

	if err = ioutil.WriteFile("./corpus/new_file.txt", output, 0666); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Print(string(output))
	// 4. Retourne les resultats en JSON
}
