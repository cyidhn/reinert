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

//Trouver un moyen pour insérer plusieurs tableaux dans une matrice. Chaque tableau doit lire la présence/absence du terme
func matrix_term_doc(doc []string) [][]int {
	var tab_binary_1 = []int{}
	var tab_binary_2 = []int{}
	var tab_binary_3 = []int{}
	var tab_binary_4 = []int{}

	for i := range doc {
		if strings.Contains(doc[i], "vote") {
			tab_binary_1 = append(tab_binary_1, 1)
		} else {
			tab_binary_1 = append(tab_binary_1, 0)
		}
		if strings.Contains(doc[i], "une") {
			tab_binary_2 = append(tab_binary_2, 1)
		} else {
			tab_binary_2 = append(tab_binary_2, 0)
		}
		if strings.Contains(doc[i], "est") {
			tab_binary_3 = append(tab_binary_3, 1)
		} else {
			tab_binary_3 = append(tab_binary_3, 0)
		}
		if strings.Contains(doc[i], "obligatoire") {
			tab_binary_4 = append(tab_binary_4, 1)
		} else {
			tab_binary_4 = append(tab_binary_4, 0)
		}
	}
	matrix = append(matrix, tab_binary_1, tab_binary_2, tab_binary_3, tab_binary_4)
	return matrix
}

//La fonction ne marche pas
func matrice_terme_doc(doc []string) [][]int {
	var terme = "vote"
	var tab_count = []int{}
	//Il faut que les lignes: Termes et Colonnes: Documents
	for i := range doc {
		sep := strings.Fields(doc[i]) //Lire chaque token de mot
		for _, word := range sep {
			if word == terme {
				tab_count = append(tab_count, 1)
			} else {
				tab_count = append(tab_count, 0)
			}
		}
		fmt.Println(sep)
		fmt.Println(tab_count)
	}
	matrix = append(matrix, tab_count)
	//fmt.Println(doc)
	return matrix
}

//2) Développer une fonction pour appliquer le découpage de la segmentation du texte
//3) Calcul et traitement de la matrice termes-documents (dtm)
//4) Application de l'algorithme de Rainette (Clustering des mots les plus approchés dans chaque segmentation) CHD
