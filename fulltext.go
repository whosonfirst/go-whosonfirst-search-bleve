package bleve

import (
	"bytes"
	"context"
	"fmt"
	"github.com/blevesearch/bleve"
	"github.com/whosonfirst/go-whosonfirst-search/filter"
	"github.com/whosonfirst/go-whosonfirst-search/fulltext"
	"github.com/whosonfirst/go-whosonfirst-spr/v2"
	"net/url"
)

type BleveFullTextDatabase struct {
	fulltext.FullTextDatabase
	index bleve.Index
}

func init() {
	ctx := context.Background()
	fulltext.RegisterFullTextDatabase(ctx, "bleve", NewBleveFullTextDatabase)
}

func NewBleveFullTextDatabase(ctx context.Context, str_uri string) (fulltext.FullTextDatabase, error) {

	u, err := url.Parse(str_uri)

	if err != nil {
		return nil, err
	}

	idx, err := NewIndex(ctx, u.Path)

	if err != nil {
		return nil, err
	}

	ftdb := &BleveFullTextDatabase{
		index: idx,
	}

	return ftdb, nil
}

func (ftdb *BleveFullTextDatabase) Close(ctx context.Context) error {
	return nil
}

func (ftdb *BleveFullTextDatabase) IndexFeature(ctx context.Context, f []byte) error {
	br := bytes.NewReader(f)
	return IndexReader(ctx, ftdb.index, br)
}

func (ftdb *BleveFullTextDatabase) QueryString(ctx context.Context, q string, filters ...filter.Filter) (spr.StandardPlacesResults, error) {

	bleve_query := bleve.NewQueryStringQuery(q)

	req := bleve.NewSearchRequest(bleve_query)
	rsp, err := ftdb.index.Search(req)

	if err != nil {
		return nil, fmt.Errorf("Failed to query '%s', %w", q, err)
	}

	for  range rsp.Hits {

		// TBD:
		// Fetch doc with go-reader.Reader instance or ... ?
		// Create stripped down ID-only SPR interface ... ?
		
		// log.Println(doc.ID)
	}

	return nil, fmt.Errorf("Not implemented")
}
