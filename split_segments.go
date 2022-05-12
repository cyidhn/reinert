package main

import (
	"fmt"
	"strings"
)

var tab_doc []string
var matrix [][]int

func split_segments_words(words string, segment_size int) []string {
	sep := strings.Fields(words)
	for i, word := range sep {
		if i < segment_size {
			tab_doc = append(tab_doc, word)
		}
	}
	return tab_doc
}

func matrice_terme_doc(doc []string) [][]int {
	//1) Trouver une condition pour lire les termes qui possèdent le plus dans chaque document
	//2) Combien de termes occurences on va écrire sur la matrice ?
	var terme = "vote"
	var tab_count = []int{}

	for i := range doc {
		sep := strings.Fields(doc[i]) //Lire chaque token de mot
		for _, word := range sep {
			if word == terme {
				tab_count = append(tab_count, 1)
				//En sortie: Lignes: Termes et Colonnes: Documents
			} else {
				tab_count = append(tab_count, 0)
			}
		}
		fmt.Println(sep)
		fmt.Println(tab_count)
	}
	matrix = append(matrix, tab_count) //Trouver un moyen pour insérer plusieurs tableaux dans une matrice. Un tableau = Une document
	//append(matrix,tab1,tab2,tab3,tabn...)

	//fmt.Println(doc)
	return matrix
}

//2) Développer une fonction pour appliquer le découpage de la segmentation du texte
//3) Calcul et traitement de la matrice termes-documents (dtm)
//4) Application de l'algorithme de Rainette (Clustering des mots les plus approchés dans chaque segmentation) CHD
