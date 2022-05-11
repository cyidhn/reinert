package main

import (
	"fmt"
	"strings"
)

var tab_doc []string
var matrix []int

func split_segments_words(words string, segment_size int) []string {
	sep := strings.Fields(words)
	for i, word := range sep {
		if i < segment_size {
			tab_doc = append(tab_doc, word)
		}
	}
	return tab_doc
}

func matrice_terme_doc(doc []string) []int {
	//1) Trouver une condition pour lire les termes qui possèdent le plus dans chaque document
	//2) Combien de termes occurences on va écrire sur la matrice ?
	var id_word = "vote"
	for i := range doc { //Pour chaque terme = ligne  !=document
		sep := strings.Fields(doc[i])
		for _, word := range sep { //Pour chaque document = colonne !=termes
			count := strings.Count(doc[i], word) //Compte les mots par document => très embêtant pour régler la taille du texte
			if word == id_word {
				matrix = append(matrix, count) //En sortie: Lignes: Termes et Colonnes: Documents
			}
		}
		fmt.Println(sep)
	}
	//fmt.Println(doc)
	return matrix
}

//2) Développer une fonction pour appliquer le découpage de la segmentation du texte
//3) Calcul et traitement de la matrice termes-documents (dtm)
//4) Application de l'algorithme de Rainette (Clustering des mots les plus approchés dans chaque segmentation) CHD
