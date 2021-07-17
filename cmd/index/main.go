package main

import (
	"context"
	"flag"
	"github.com/whosonfirst/go-whosonfirst-iterate/iterator"
	wof_bleve "github.com/whosonfirst/go-whosonfirst-search-bleve"
	"io"
	"log"
)

func main() {

	bleve_uri := flag.String("bleve-uri", "whosonfirst.bleve", "...")
	iter_uri := flag.String("iterator-uri", "repo://", "...")

	flag.Parse()

	uris := flag.Args()

	ctx := context.Background()

	bleve_index, err := wof_bleve.NewIndex(ctx, *bleve_uri)

	if err != nil {
		log.Fatalf("Failed to create Bleve index, %v", err)
	}

	iter_cb := func(ctx context.Context, fh io.ReadSeeker, args ...interface{}) error {
		return wof_bleve.IndexReader(ctx, bleve_index, fh)
	}

	iter, err := iterator.NewIterator(ctx, *iter_uri, iter_cb)

	if err != nil {
		log.Fatalf("Failed to create new iterator, %v", err)
	}

	err = iter.IterateURIs(ctx, uris...)

	if err != nil {
		log.Fatalf("Failed to iterate URIS, %v", err)
	}
}
