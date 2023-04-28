package main

import (
	"bufio"
	"searchEngine/api"
	"searchEngine/internals/searchEngine"
	"strings"
)

func main() {
	doc1 := "new golang project with golang version 1.19"
	x1 := bufio.NewScanner(strings.NewReader(doc1))
	doc2 := "golang is simple but complicated"
	x2 := bufio.NewScanner(strings.NewReader(doc2))

	fastSearchEngine := searchEngine.NewSearchEngine("LinearFastSearchEngine", 1024)
	fastSearchEngine.AddData(x1, 1)
	fastSearchEngine.AddData(x2, 2)

	router := api.NewServer(fastSearchEngine)
	router.GinEngine.Run()

}
