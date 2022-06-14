package main

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"

	"github.com/bbalet/stopwords"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

func lematize(text string) string {
	var tab []string
	var wordtoappend string
	fmt.Println("Load Lemmatization ...")
	dict := csv_to_dict()
	result := strings.Fields(text)
	for word := range result {
		wordtoappend = result[word]
		for _, value := range dict {
			if value.Terme == result[word] {
				wordtoappend = value.Lemmatisation
			}
		}
		tab = append(tab, wordtoappend)
	}
	return strings.Join(tab, " ")
}

// Preprocessing du text
func preprocess(text string) string {
	lower := strings.ToLower(text)
	list_reg := `(\d+)|(http\S+)|(www\S+)|(@mention)|(&[a-z])|([,.?!'’”\/#$%\^&\*;:+{}=\-_~()«»])|([\x{1F600}-\x{1F6FF}|[\x{2600}-\x{26FF}])`
	reg := regexp.MustCompile(list_reg)                                          //Compilation du Regex
	res := reg.ReplaceAllString(lower, "")                                       //Résultats pour Regex
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC) //Enlever les accents du texte
	result, _, _ := transform.String(t, res)                                     //Résultat pour remove accents
	lematizer := lematize(string(result))                                        //Lematization
	delete_words := stopwords.CleanString(lematizer, "fr", true)                 //Stopwords
	return delete_words
}
