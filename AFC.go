package main

import (
	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
	"gonum.org/v1/gonum/mat"
)

//Ces fonctions permettront d'appliquer l'analyse factorielle des correspondances (AFC) à partir de notre matrice terme document
//Les colonnes représentent le nombre de documents et les lignes représentent la catégorie du terme
//Les valeurs représentent l'occurence d'un mot pour chaque document

//Fonction de conversion matrix vers un dataframe
func matrix_to_dataframe(term_doc mat.Matrix) dataframe.DataFrame {
	dataframe := dataframe.LoadMatrix(term_doc)
	return dataframe
}

//Fonction d'application de sommme des valeurs dans un dataframe
var sum = func(s series.Series) series.Series {
	nb := s.Float()
	count := 0.0
	for _, f := range nb {
		count += f
	}
	return series.Floats(count)
}

//Fonction marge des colonnes et des lignes
func marge_rows_columns(df dataframe.DataFrame) (dataframe.DataFrame, dataframe.DataFrame) {
	m_columns := df.Capply(sum) //Marge colonnes
	m_rows := df.Rapply(sum)    //Marge des lignes
	return m_columns, m_rows
}

//Fonction de somme des marges colonnes/lignes
func sum_marge(df dataframe.DataFrame, df2 dataframe.DataFrame) (dataframe.DataFrame, dataframe.DataFrame) {
	m_columns, m_rows := marge_rows_columns(df)
	sum_m_columns := m_columns.Rapply(sum)
	sum_m_rows := m_rows.Capply(sum)
	return sum_m_columns, sum_m_rows
}

//Fonction de calcul du total des marges
func total_marge(df dataframe.DataFrame, df2 dataframe.DataFrame) dataframe.DataFrame {
	new_df := df.Concat(df2)
	marge_total := new_df.Capply(sum)
	return marge_total
}
