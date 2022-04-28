package main 

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"unicode"
    "golang.org/x/text/transform"
    "golang.org/x/text/unicode/norm"
)

func isMn(r rune) bool {
	return unicode.Is(unicode.Mn, r) // Mn: nonspacing marks
 }

func preprocessing(string text)
	list_reg := `(\d+)|(http\S+)|(www\S+)|(@mention)|(&[a-z])`
	reg := regexp.MustCompile(list_reg)
	res := reg.ReplaceAllString(text, "")

	t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
	result, _, _ := transform.String(t, text)
 	return text

func main() {
	// 1. Fonction pour importer le document en format iramuteq
	file, err := ioutil.ReadFile("./corpus/new_file.txt")
	
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// 1. Traitement de texte pour les tweets
	
	// 3. Algo de CHD
	// 4. Retourne les resultats en JSON
	fmt.Println(preprocessing(file))
}