package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
	"unicode"

	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

func isMn(r rune) bool {
	return unicode.Is(unicode.Mn, r) // Mn: nonspacing marks
}

//func lemmatisation(text string) {
//var data []DictionnaireStruct
//data = data.append(data, DictionnaireStruct{"terme", "lem"})
//result := strings.Split(text, " ")
//for i := range result {
//fmt.Println(result[i])

//}
//}

func main() {
	// 1. Fonction pour importer le document en format iramuteq
	file, err := ioutil.ReadFile("./corpus/new_file.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// 1. Traitement de texte pour les tweets
	min := strings.ToLower(string(file))
	list_reg := `(\d+)|(http\S+)|(www\S+)|(@mention)|(&[a-z])|([.,'’”\/#?!$%\^&\*;:+{}=\-_~()«»])|([\x{1F600}-\x{1F6FF}|[\x{2600}-\x{26FF}])`
	reg := regexp.MustCompile(list_reg)          // Test pour les ponctuations
	res := reg.ReplaceAllString(string(min), "") //Résultat pour Regex

	//Enlever UNIQUEMENT LES ACCENTS DU TEXTE
	t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
	result, _, _ := transform.String(t, string(res)) //Résultat pour remove accents
	// 3. Algo de CHD
	// 4. Retourne les resultats en JSON
	fmt.Println(result)
}
