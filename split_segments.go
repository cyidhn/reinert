package main

import (
	"fmt"
	"math"
	"strings"
)

var tab_doc []string
var matrix [][]int
var group_matrix [][]int

//Fonction qui permet de segmenter les mots pour un seul document
func split_segments_words(doc string, segment_size int) []string {
	sep := strings.Fields(doc)
	for i, word := range sep {
		if i < segment_size {
			tab_doc = append(tab_doc, word)
		}
	}
	return tab_doc
}

//Création d'une matrice terme-document pour un exemple sur 4 documents et 4 termes
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

//Fonction regroupement des documents sur la matrice terme-documents
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

func methode_reinert() {
	var marge_rows []int
	var marge_columns []int
	var J1, J2, J3, J4, marge_total int
	var tab_freq []float64
	var freq1, freq2 float64

	//Calcul marge ligne
	for i, rows := range group_matrix {
		rowsum := 0
		for j := range rows {
			rowsum = rowsum + group_matrix[i][j]
		}
		marge_rows = append(marge_rows, rowsum)
	}

	//Calcul marge colonnes/lignes (Le boucle for ne marche pas puisque la variable de la matrice group_matrix ne possède pas les mêmes dimensions)
	J1 = group_matrix[0][0] + group_matrix[1][0]
	J2 = group_matrix[0][1] + group_matrix[1][1]
	J3 = group_matrix[0][2] + group_matrix[1][2]
	J4 = group_matrix[0][3] + group_matrix[1][3]
	marge_columns = append(marge_columns, J1, J2, J3, J4)

	sum_m_column := 0
	sum_m_rows := 0
	for k := range marge_columns {
		sum_m_column += (marge_columns[k])
	}

	for l := range marge_rows {
		sum_m_rows += (marge_rows[l])
	}

	marge_total = sum_m_column + sum_m_rows
	fmt.Println("Marge de colonne:", marge_columns)
	fmt.Println("Marge de lignes:", marge_rows)
	fmt.Println("Marge total:", marge_total)

	//Calcul de la fréquence pour chaque individu
	for n := 0; n <= 3; n++ {
		freq1 = ((float64(marge_columns[n] * marge_rows[0])) / float64(marge_total))
		freq2 = ((float64(marge_columns[n] * marge_rows[1])) / float64(marge_total))
		tab_freq = append(tab_freq, freq1, freq2)
	}
	fmt.Println("Tableaux de fréquence:", tab_freq)

	//Calcul du chi2
	var chi2 float64
	chi2 = math.Pow(float64(group_matrix[0][0])-tab_freq[0], 2) / tab_freq[0]
	fmt.Println("Premier chi2:", chi2)
}

//2) Développer une fonction pour appliquer le découpage de la segmentation du texte
//3) Calcul et traitement de la matrice termes-documents (dtm)
//4) Application de l'algorithme de Rainette (Clustering des mots les plus approchés dans chaque segmentation) CHD
