package bleve

import (
	"bytes"
	"context"
	"errors"
	"github.com/blevesearch/bleve"
	"github.com/whosonfirst/go-whosonfirst-search/filter"
	"github.com/whosonfirst/go-whosonfirst-search/fulltext"
	"github.com/whosonfirst/go-whosonfirst-spr"
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
	return nil, errors.New("Not implemented")
}
