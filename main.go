package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
	"time"
	"unicode"

	_ "embed"

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
	list_reg := `(\d+)|(http\S+)|(www\S+)|(@mention)|(&[a-z])|(['’”\/#$%\^&\*;:+{}=\-_~()«»])|([\x{1F600}-\x{1F6FF}|[\x{2600}-\x{26FF}])`
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
	//1. Segmentation des corpus à partir d'un fichier Iramuteq

	file, err := ioutil.ReadFile("./corpus/clean_file.txt") //Lecture tout l'intégralité du texte
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var tab_doc = segmentation_text(string(file), 50)
	fmt.Println(tab_doc)
	//2. Traitement de pre-processing

	/*
		file, err := ioutil.ReadFile("./corpus/clean_file.txt") //Lecture tout l'intégralité du texte
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		new_file := []byte(preprocess(string(file))) //Ecrire l'intégralité du texte des mots lemmatisés par des valeurs binaires

		err2 := ioutil.WriteFile("./corpus/text_lematize_3.txt", new_file, 0)
		if err2 != nil {
			log.Fatal(err)
		}

		fmt.Println(string(new_file))
	*/

	//3. Application de la méthode de Reinert
	/*
		var read_matrix [][]int
		doc := [...]string{"si le vote blanc soit comptabilité", "C'est nécesssaire  pour notre démocratie !", "une vote impérieuse  doute et nulle", "obligatoire une", "Le vote devrait rendu  obligatoire !"}
		termes := []string{"vote", "une", "est", "obligatoire"}

		read_matrix = matrix_term_doc(doc[:], termes)
		fmt.Println("Matrice Terme document:", read_matrix)
		fmt.Println("Matrice de regroupement des documents:", regroupement_doc(read_matrix))
		var tab_freq_1, tab_freq_2 = tab_frequence(regroupement_doc(read_matrix))
		var chi2 = calcul_chi2(regroupement_doc(read_matrix), tab_freq_1, tab_freq_2)

		//fmt.Println("Classe 1 fréquence:", tab_freq_1, "classe 2 fréquence :", tab_freq_2)
		fmt.Println("chi2=", chi2)
	*/
	//Besoin de trouver la maximisation de chi2 pour trouver le regroupement final

	//4. Algo de CHD

	//5. Retourne les resultats en JSON
}
