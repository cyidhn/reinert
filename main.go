package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func main() {
	//1. Ouverture du fichier Iramuteq
	start := time.Now()
	file, err := ioutil.ReadFile("./corpus/clean_file.txt") //Lecture tout l'intégralité du texte
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//2. Nettoyage du texte
	pro := preprocess(string(file), false) //Vrai si l'utilisateur procède la lemmatisation, faux sinon

	//3. Application de Segmentation du texte
	tab_doc := segmentation_text(pro, 50) //Segmentation sur 50 mots

	//4. Compter le nombre de termes pour chaque document pour créer par la suite une matrice terme-document
	new_tab_doc := select_nbdoc(tab_doc, 1, 3)

	//5. Création une matrice terme document choisi par rapport aux nombres de documents
	fmt.Println(matrix_term_doc_(new_tab_doc))

	//5. Application de la méthode de Reinert
	/*
			var read_matrix [][]int
			doc := [...]string{"si le vote blanc soit comptabilité", "C'est nécesssaire  pour notre démocratie !", "une vote impérieuse  doute et nulle", "obligatoire une", "Le vote devrait rendu  obligatoire !"}
			termes := []string{"vote", "une", "est", "obligatoire"}

			read_matrix = matrix_term_doc(doc[:], termes)
			fmt.Println("Matrice Terme document:", read_matrix)
			fmt.Println("Matrice de regroupement des documents:", regroupement_doc(read_matrix))
			var tab_freq_1, tab_freq_2 = tab_frequence(regroupement_doc(read_matrix))
			var chi2 = calcul_chi2(regroupement_doc(read_matrix), tab_freq_1, tab_freq_2)

			//fmt.Println("Classe 1 fréquence:", tab_freq_1, "classe 2 fréquence :", tab_freq_2)
			fmt.Println("chi2=", chi2)
			//Besoin de trouver la maximisation de chi2 pour trouver le regroupement final
		//6. Algo de CHD
		//7. Retourne les resultats en JSON
	*/
	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}
