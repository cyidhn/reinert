package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"
	"time"
	"unicode"

	"github.com/bbalet/stopwords"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

func isMn(r rune) bool {
	return unicode.Is(unicode.Mn, r) // Mn: nonspacing marks
}

func lematize(text string) string {
	start := time.Now()
	var tab []string
	var wordtoappend string
	dict := csv_to_dict()
	result := strings.Fields(text)
	fmt.Println("Load Lemmatization ...")
	for word := range result {
		wordtoappend = result[word]
		for _, value := range dict {
			if value.Terme == result[word] {
				wordtoappend = value.Lemmatisation
			}
		}
		tab = append(tab, wordtoappend)
	}
	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
	return strings.Join(tab, " ")
}

// Preprocessing du text
func preprocess(text string) string {
	lower := strings.ToLower(text)
	list_reg := `(\d+)|(http\S+)|(www\S+)|(@mention)|(&[a-z])|([.,'’”\/#?!$%\^&\*;:+{}=\-_~()«»])|([\x{1F600}-\x{1F6FF}|[\x{2600}-\x{26FF}])`
	reg := regexp.MustCompile(list_reg)                                  //Compilation du Regex
	res := reg.ReplaceAllString(lower, "")                               //Résultats pour Regex
	t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC) //Enlever les accents du texte
	result, _, _ := transform.String(t, res)                             //Résultat pour remove accents
	lematizer := lematize(string(result))                                //Lematization
	delete_words := stopwords.CleanString(lematizer, "fr", true)         //Stopwords
	//fmt.Println(lematizer)
	return delete_words
}

func main() {
	//1. Fonction pour importer le document en format iramuteq

	//file, err := ioutil.ReadFile("./corpus/clean_file.txt") //Lecture tout l'intégralité du texte
	//if err != nil {
	//	fmt.Println(err)
	//	os.Exit(1)
	//}

	//new_file := []byte(preprocess(string(file))) //Ecrire l'intégralité du texte des mots lemmatisés par des valeurs binaires

	//err2 := ioutil.WriteFile("./corpus/text_lematize.txt", new_file, 0)
	//if err2 != nil {
	//	log.Fatal(err)
	//}

	//tab_doc = split_segments_words("Le vote devrait être rendu obligatoire si les votes blancs sont comptabilités. C'est une nécessité démocratique pour notre pays et ses citoyens, une obligation impérieuse", 10)
	//fmt.Println(tab_doc)

	documents := [...]string{"Le vote devrait rendu obligatoire", "si le vote blanc est comptabilité", "C'est une nécessité démocratique est notre", "une vote impérieuse doute et nulle"}
	termes := []string{"vote", "une", "est", "obligatoire"}
	var read_matrix [][]int
	read_matrix = matrix_term_doc(documents[:], termes)
	fmt.Println("Matrice Terme document:", read_matrix)
	fmt.Println("Matrice de regroupement des documents:", regroupement_doc(read_matrix))
	fmt.Println("Tableaux de fréquence des individus:", tab_frequence(regroupement_doc(read_matrix)))
	calcul_chi2(regroupement_doc(read_matrix), tab_frequence(regroupement_doc(read_matrix)))

	// 3. Algo de CHD
	// 4. Retourne les resultats en JSON
}
