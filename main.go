package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	//1. Segmentation des corpus + Traitement du pre-processing à partir d'un fichier Iramuteq

	file, err := ioutil.ReadFile("./corpus/clean_file.txt") //Lecture tout l'intégralité du texte
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var tab_doc = segmentation_text(string(file), 10)
	for i := range tab_doc {
		tab_doc[i] = preprocess(tab_doc[i])
	}

	fmt.Println(tab_doc)
	fmt.Println("Taille du tableau des corpus: ", len(tab_doc))
	var tab_words = regroupement_tokens(tab_doc)
	fmt.Println(tab_words)

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
		//Besoin de trouver la maximisation de chi2 pour trouver le regroupement final
	*/
	//4. Algo de CHD
	//5. Retourne les resultats en JSON
}
