package main 

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	//"unicode"
    //"golang.org/x/text/transform"
    //"golang.org/x/text/unicode/norm"
)

//func isMn(r rune) bool {
	//return unicode.Is(unicode.Mn, r) // Mn: nonspacing marks
 //}

func main() {
	// 1. Fonction pour importer le document en format iramuteq
	file, err := ioutil.ReadFile("./corpus/new_file.txt")
	
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// 1. Traitement de texte pour les tweets
	list_reg := `(\d+)|(http\S+)|(www\S+)|(@mention)|(&[a-z])`
	reg := regexp.MustCompile(list_reg) // Test pour les ponctuations
	res := reg.ReplaceAllString(string(file), "")  //Résultat pour Regex 

	//t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
   	//result, _, _ := transform.String(t, string(file))  //Résultat piur remove accents
	
	// But: result + res soit appliqué dans file pour un seul traitement 
	// 3. Algo de CHD
	// 4. Retourne les resultats en JSON
	fmt.Println(res)
	//fmt.Println(result)
}