package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
	"unicode"

	"github.com/bbalet/stopwords"

	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

var dict []map[string]string

func isMn(r rune) bool {
	return unicode.Is(unicode.Mn, r) // Mn: nonspacing marks
}

func lematize(text string) string {
	var tab []string
	var wordtoappend string
	dict := csv_to_dict()
	result := strings.Fields(text)
	fmt.Println("Load Lemmatization ...")
	for word := range result {
		wordtoappend = result[word]
		for _, value := range dict {
			if value["1_ortho"] == result[word] {
				wordtoappend = value["3_lemme"]
			}
		}
		tab = append(tab, wordtoappend)
	}
	return strings.Join(tab, " ")
}

// Preprocessing du text
func preprocess(text string) string {
	lower := strings.ToLower(text)
	list_reg := `(\d+)|(http\S+)|(www\S+)|(@mention)|(&[a-z])|([.,'’”\/#?!$%\^&\*;:+{}=\-_~()«»])|([\x{1F600}-\x{1F6FF}|[\x{2600}-\x{26FF}])`
	reg := regexp.MustCompile(list_reg)                                  //Compilation du Regex
	res := reg.ReplaceAllString(lower, "")                               //Résultat pour Regex
	t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC) //Enlever les accents du texte
	result, _, _ := transform.String(t, res)                             //Résultat pour remove accents
	lematizer := lematize(string(result))                                //Lematization (Environ 15 -20 min du temps d'exécution)
	delete_words := stopwords.CleanString(lematizer, "fr", true)         //Stopwords
	//fmt.Println(lematizer)
	return delete_words
}

func main() {
	// 1. Fonction pour importer le document en format iramuteq
	file, err := ioutil.ReadFile("./corpus/clean_file.txt") //Permet de lire tout l'intégralité du texte
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(preprocess(string(file)))
	//remove_words(file)
	// 3. Algo de CHD
	// 4. Retourne les resultats en JSON
}
