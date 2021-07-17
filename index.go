package bleve

import (
	"context"
	"fmt"
	"github.com/blevesearch/bleve"
	"github.com/tidwall/gjson"
	"io"
	"time"
	"log"
)

func IndexReader(ctx context.Context, index bleve.Index, r io.Reader) error {

	select {
	case <- ctx.Done():
		return nil
	default:
		// pass
	}
	
	body, err := io.ReadAll(r)

	if err != nil {
		return err
	}

	rsp := gjson.GetBytes(body, "properties")

	if !rsp.Exists() {
		return fmt.Errorf("Document is missing properties")
	}

	id_rsp := rsp.Get("wof:id")

	if !id_rsp.Exists() {
		return fmt.Errorf("Missing wof:id property")
	}

	str_id := id_rsp.String()
	doc := rsp.Value()

	t1 := time.Now()
	
	defer func(){
		log.Printf("Time to index '%s', %v\n", str_id, time.Since(t1))
	}()
	
	return index.Index(str_id, doc)
}
