package main

import (
	"flag"
	"fmt"
	"os"

	"gopkg.in/olivere/elastic.v3"
)

func main() {
	var (
		url = flag.String("url", "", "Elasticsearch URL")
	)
	flag.Parse()

	client, err := elastic.NewClient(elastic.SetURL(*url))
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	indices, err := client.IndexNames()
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	for _, name := range indices {
		fmt.Println(name)
	}
}
