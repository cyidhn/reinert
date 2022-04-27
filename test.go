package main 

import (
	"github.com/bbalet/stopwords"
	
)

func main() {
//Example with 2 strings containing P html tags
//"la", "un", etc. are (stop) words without lexical value in French
string1 := []byte("<p>la fin d'un bel après-midi d'été</p>")
string2 := []byte("<p>cet été, nous avons eu un bel après-midi</p>")

//Return a string where HTML tags and French stop words has been removed
cleanContent := stopwords.CleanString(string1, "fr", true)

//Get two (Sim) hash representing the content of each string
hash1 := stopwords.Simhash(string1, "fr", true)
hash2 := stopwords.Simhash(string2, "fr", true)

//Hamming distance between the two strings (diffference between contents)
distance := stopwords.CompareSimhash(hash1, hash2)

//Clean the content of string1 and string2, compute the Levenshtein Distance
stopwords.LevenshteinDistance(string1, string2, "fr", true)

}