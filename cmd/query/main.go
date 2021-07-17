package main

import (
	"context"
	"flag"
	"github.com/blevesearch/bleve"
	wof_bleve "github.com/whosonfirst/go-whosonfirst-search-bleve"
	"log"
	"strings"
)

func main() {

	bleve_uri := flag.String("bleve-uri", "whosonfirst.bleve", "...")

	flag.Parse()

	ctx := context.Background()
	terms := flag.Args()

	bleve_index, err := wof_bleve.NewIndex(ctx, *bleve_uri)

	if err != nil {
		log.Fatalf("Failed to create Bleve index, %v", err)
	}

	qs := strings.Join(terms, " ")
	bleve_query := bleve.NewQueryStringQuery(qs)

	req := bleve.NewSearchRequest(bleve_query)
	rsp, err := bleve_index.Search(req)

	if err != nil {
		log.Fatalf("Failed to query '%s', %v", qs, err)
	}

	for _, doc := range rsp.Hits {

		log.Println(doc.ID)
	}
}
