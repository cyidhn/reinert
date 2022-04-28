package main 

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
)

func main() {
	// 1. Fonction pour importer le document en format iramuteq
	file, err := ioutil.ReadFile("./corpus/new_file.txt")
	
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// 1. Traitement de texte pour les tweets
	list_reg := `([^\w\s])|(\d+)|(http\S+)|(www\S+)|(@mention)|(&[a-z])`
	reg := regexp.MustCompile(list_reg) // Test pour les ponctuations
	res := reg.ReplaceAllString(string(file), "")
	// 3. Algo de CHD
	// 4. Retourne les resultats en JSON
	fmt.Println(res)
}