package main

import (
	"fmt"
	"interface/testing"
)

func getRetriever() retriever {
	return testing.Retriever{}
}

type retriever interface {
	Get(string) string
}

func main() {
	r := getRetriever()
	bytes := r.Get("https://imooc.com")
	fmt.Println(bytes)
}
