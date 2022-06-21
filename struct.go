package main

type LemmatisationStruct struct {
	Terme         string
	Lemmatisation string
}

type Dictionary struct {
	Key   string
	Value int
}

type DictionaryList []Dictionary

type tokens_doc struct {
	number_doc  int
	list_tokens []string
}
