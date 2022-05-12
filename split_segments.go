package main

import (
	"strings"
)

var tab_doc []string
var matrix [][]int
var group_matrix [][]int

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
func matrix_term_doc(doc []string, termes []string) [][]int {
	var tab_binary_1 = []int{}
	var tab_binary_2 = []int{}
	var tab_binary_3 = []int{}
	var tab_binary_4 = []int{}

	for i := range doc {
		if strings.Contains(doc[i], termes[0]) {
			tab_binary_1 = append(tab_binary_1, 1)
		} else {
			tab_binary_1 = append(tab_binary_1, 0)
		}
		if strings.Contains(doc[i], termes[1]) {
			tab_binary_2 = append(tab_binary_2, 1)
		} else {
			tab_binary_2 = append(tab_binary_2, 0)
		}
		if strings.Contains(doc[i], termes[2]) {
			tab_binary_3 = append(tab_binary_3, 1)
		} else {
			tab_binary_3 = append(tab_binary_3, 0)
		}
		if strings.Contains(doc[i], termes[3]) {
			tab_binary_4 = append(tab_binary_4, 1)
		} else {
			tab_binary_4 = append(tab_binary_4, 0)
		}
	}
	matrix = append(matrix, tab_binary_1, tab_binary_2, tab_binary_3, tab_binary_4)
	return matrix
}

func regroupement_doc() [][]int {
	var group1 []int
	var group2 []int
	for element := range matrix {
		group1 = append(group1, matrix[element][0]+matrix[element][1])
		group2 = append(group2, matrix[element][2]+matrix[element][3])
	}
	group_matrix = append(group_matrix, group1, group2)
	return group_matrix
}

func diviseur_khi2() {

}

//2) Développer une fonction pour appliquer le découpage de la segmentation du texte
//3) Calcul et traitement de la matrice termes-documents (dtm)
//4) Application de l'algorithme de Rainette (Clustering des mots les plus approchés dans chaque segmentation) CHD
