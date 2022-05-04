package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func csv_to_dict() []map[string]string {
	csvFile, err := os.Open("./corpus/Lexique383.csv")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened CSV file")
	defer csvFile.Close()

	csvlines := csv.NewReader(csvFile)
	if err != nil {
		fmt.Println(err)
	}

	rows := []map[string]string{}
	var header []string

	for {
		record, err := csvlines.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if len(header) == 0 {
			header = record
		} else {
			dict := map[string]string{}
			for i := range header {
				dict[header[i]] = record[i]
			}
			rows = append(rows, dict)
		}
	}
	//fmt.Println(rows)
	return rows
}
