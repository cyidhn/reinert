package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

type DictionnaireStruct struct {
	Terme         string
	Lemmatisation string
}

func read_csv() {
	csvFile, err := os.Open("./corpus/Lexique383.csv")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened CSV file")
	defer csvFile.Close()

	csvlines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	for _, line := range csvlines {
		dict := DictionnaireStruct{
			Terme:         line[0],
			Lemmatisation: line[1],
		}
		fmt.Println(dict.Terme + " " + dict.Lemmatisation)
	}
}
