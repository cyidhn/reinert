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

func segmentation_text(doc string, segment_size int) []string {
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
