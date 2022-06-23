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
func preprocess(text string, lema bool) string {
	var list_reg, result, lematizer, delete_words, res, clean_text string
	var reg *regexp.Regexp
	var accent transform.Transformer

	list_reg = `|(\d+)|(http\S+)|(www\S+)|(@mention)|(/[^\w\s]/gi)|(/[~!@#$%^&*()_|+\-=?;:'",.<>\{\}\[\]\\\/]/gi)|([\x{1F600}-\x{1F6FF}|[\x{2600}-\x{26FF}])`
	reg = regexp.MustCompile(list_reg)                                               //Compilation du Regex
	res = reg.ReplaceAllString(text, "")                                             //Résultats pour Regex
	accent = transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC) //Enlever les accents du texte
	result, _, _ = transform.String(accent, res)                                     //Résultat pour remove accents
	//Condition d'application de lematisation
	if lema {
		lematizer = lematize(string(result))                        //Lematization
		delete_words = stopwords.CleanString(lematizer, "fr", true) //Stopwords
		clean_text = remove_words(delete_words)
		return clean_text
	} else {
		delete_words = stopwords.CleanString(result, "fr", true)
		clean_text = remove_words(delete_words)
		return clean_text
	}
}

func remove_words(text string) string {
	stringToRemove := []string{"ca", "ve", "n", "s", "d", "l", "j", "y", "c", "e", "m", "h", "quelqu", "cht", "lr", "oas", "qu", "ll", "yu", "an", "g", "TRUE", "aujourd", "-", "--", "er"}
	string_filtered := make([]string, 0)
	for _, word := range strings.Split(text, " ") {
		shouldAppend := true
		lowered := strings.ToLower(word)
		for _, word2 := range stringToRemove {
			if lowered == word2 {
				shouldAppend = false
				break
			}
		}
		if shouldAppend {
			string_filtered = append(string_filtered, lowered)
		}
	}
	resultString := strings.Join(string_filtered, " ")
	return resultString
}
