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
	list_words := []string{}
	seg := []string{}

	//1. Traitement du pre-processing
	pro := preprocess(doc)
	sep := strings.Fields(pro)
	size_word := len(sep) / segment_size

	//2. Ajout tout les mots
	for i := range sep {
		list_words = append(list_words, sep[i])
	}
	var cpt = 0

	//3. Découpage des segments en fonction de la quantité des mots
	for _, word := range list_words {
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

func dict_doc_terme(tab_seg [][]string) map[string]int {
	dict_terme := make(map[string]int)
	wordlist := tab_seg[:][0]
	for _, word := range wordlist {
		_, ok := dict_terme[word]
		if ok {
			dict_terme[word] += 1
		} else {
			dict_terme[word] = 1
		}
	}
	return dict_terme
}
