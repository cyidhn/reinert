package main

import (
	"regexp"
	"strings"
)

//Fonction qui permet de segmenter les mots pour un seul document
func split_segments_words(doc string, segment_size int) []string {
	var tab_doc []string
	sep := strings.Fields(doc)
	for i, word := range sep {
		if i < segment_size {
			tab_doc = append(tab_doc, word)
		}
	}
	return tab_doc
}

func segmentation_sentence(doc string, segment_size int) []string {
	tab_doc := []string{}
	reg := regexp.MustCompile(`[,.?!]`)
	sep := reg.Split(doc, -1)
	for i := range sep {
		if i < segment_size {
			tab_doc = append(tab_doc, sep[i])
		}
	}
	return tab_doc
}

//Fonction de segmentation de text
func segmentation_text(doc string, segment_size int) [][]string { //Un document de 20 mots pour utiliser une taille de 4 segments
	tab_seg := [][]string{} //Stocker 4 segments
	list_tokens := tokens_all(doc)
	seg := []string{}
	size_word := len(list_tokens) / segment_size

	var cpt = 0
	//3. Découpage des segments en fonction de la quantité des mots
	for _, word := range list_tokens {
		seg = append(seg, word)
		cpt++
		if cpt == size_word {
			tab_seg = append(tab_seg, seg)
			seg = []string{}
			cpt = 0
		}
	}

	//S'il contient le reste des mots
	if seg != nil {
		tab_seg = append(tab_seg, seg)
	}

	return tab_seg
}

func tokens_all(doc string) []string {
	list_words := []string{}
	sep := strings.Fields(doc)
	for i := range sep {
		list_words = append(list_words, sep[i])
	}
	return list_words
}

//Fonction qui permet de compter le nombre de vocabulaire sur un dictionnaire
func count_vocabulary(doc string) map[string]int {
	//Trois variables pour permettre de créer une matrice creuse au format CSR
	/*
		var list_values []int //Valeurs des indices non nulles IA = NNN
		var i_indices []int   //valeurs récursives des lignes IA(i+1) = IA(i) + NNNi
		var j_indices []int   //Valeurs récursives des colonnes IA(j+1) = IA(j) + NNNj
	*/
	dict_terme := make(map[string]int)
	liste_words := tokens_all(doc)
	for _, word := range liste_words {
		_, ok := dict_terme[word]
		if ok {
			dict_terme[word] += 1
		} else {
			dict_terme[word] = 1
		}
	}
	return dict_terme
}
