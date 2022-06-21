package main

import (
	"fmt"
)

func main() {
	//1. Segmentation des corpus à partir d'un fichier Iramuteq

	var tab_doc = segmentation_text("Le vote devrait être rendu obligatoire si les votes blancs sont comptabilités. C'est une nécessité démocratique pour notre pays et ses citoyens, une obligation impérieuse.", 5)
	fmt.Println(tab_doc)
	matrix := [][]string{{"Le vote devrait être rendu"}, {"obligatoire si les votes blancs"}, {"sont comptabilités. C'est une nécessité"}, {"émocratique pour notre pays et"}, {"ses citoyens, une obligation impérieuse."}}
	fmt.Println(tokens_segmentation(matrix))
	/*
		start := time.Now()
		file, err := ioutil.ReadFile("./corpus/clean_file.txt") //Lecture tout l'intégralité du texte
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		//Nettoyage du texte
		var pro = preprocess(string(file), false)
		//Segmentation du texte
		var tab_doc = segmentation_text(pro, 50)
		fmt.Println(tab_doc)

		//2. Compter le nombre de termes pour chaque élement pour créer par la suite une matrice terme-document
		//fmt.Println(count_vocabulary(pro))

		elapsed := time.Since(start)
		log.Printf("Binomial took %s", elapsed)
	*/

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
}
