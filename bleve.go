// package bleve implements the whosonfirst/go-whosonfirst-search interfaces using the Bleve text indexer.
package bleve

import (
	"context"
	"github.com/blevesearch/bleve"
	"os"
)

// NewIndex returns a bleve.Index instance for uri. If uri exists on disk the method will open an
// existing Bleve index at that path.
func NewIndex(ctx context.Context, uri string) (bleve.Index, error) {

	_, err := os.Stat(uri)

	if err != nil && !os.IsNotExist(err) {
		return nil, err
	}

	if err == nil {
		return bleve.Open(uri)
	}

	bleve_mapping := bleve.NewIndexMapping()
	return bleve.New(uri, bleve_mapping)
}
