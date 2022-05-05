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

func clean_words(text string) {
	var tab []string
	liste_terme := []string{"ve", "n", "s", "d", "l", "j", "y", "c", "e", "m", "h", "quelqu", "cht", "lr", "oas", "qu", "ll", "yu", "an", "g", "TRUE", "jadot", "avectaubira", "zemmourcroissance", "zemmourlille", "cdanslair", "taubirasorbonne", "emmanuel", "bfmpolitique", "aujourd", "macron"}
	result := strings.Fields(text)
	for word := range result {
		if result[word] != liste_terme[word] {
			tab = append(tab)
		}
	}
}

func lematize(text string) string {
	var tab []string
	var wordtoappend string
	dict := csv_to_dict()
	result := strings.Fields(text) //Tableau
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
	reg := regexp.MustCompile(list_reg)                                  // Test pour les ponctuations
	res := reg.ReplaceAllString(lower, "")                               //Résultat pour Regex
	t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC) //Enlever UNIQUEMENT LES ACCENTS DU TEXTE
	result, _, _ := transform.String(t, res)                             //Résultat pour remove accents
	delete_words := stopwords.CleanString(result, "fr", true)            //Stopwords
	//lematizer := lematize(string(delete_words))
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
	// 3. Algo de CHD
	// 4. Retourne les resultats en JSON
}
