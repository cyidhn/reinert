package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func main() {
	//1. Segmentation des corpus à partir d'un fichier Iramuteq
	start := time.Now()
	file, err := ioutil.ReadFile("./corpus/clean_file.txt") //Lecture tout l'intégralité du texte
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var tab_doc = segmentation_text(string(file), 50)

	//2. Traitement du pre-processing
	fmt.Println("Load Lemmatization ...")
	for i := range tab_doc {
		tab_doc[i] = preprocess(tab_doc[i])
	}
	//3. Regroupement des tokens sur une matrice
	// ATTENTION ! Il y a des élements dans la liste qui contiennent des élements vides il faut bien penser à supprimer
	var tab_words = regroupement_tokens(tab_doc)
	fmt.Println("Liste des termes pour chaque document:", tab_words)
	fmt.Println("Taille du tableau du corpus: ", len(tab_words))
	for i := range tab_words {
		if tab_words[i] == nil {
			remove(tab_words[i], i) //Erreur: ça enlève l'index dans l'élement du tableau
			fmt.Println("Liste des termes pour chaque document après le nettoyage:", tab_words)
		}
	}
	//4. Compter le nombre de termes pour chaque élement pour créer par la suite une matrice terme-document

	//5. Application de la méthode de Reinert
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
	//6. Algo de CHD
	//7. Retourne les resultats en JSON
	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}
