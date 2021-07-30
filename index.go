package bleve

import (
	"context"
	"fmt"
	"github.com/blevesearch/bleve"
	"github.com/tidwall/gjson"
	"io"
	"time"
	"log"
	"strings"
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

	props_rsp := gjson.GetBytes(body, "properties")

	if !props_rsp.Exists() {
		return fmt.Errorf("Document is missing properties")
	}

	id_rsp := props_rsp.Get("wof:id")

	if !id_rsp.Exists() {
		return fmt.Errorf("Missing wof:id property")
	}

	str_id := id_rsp.String()

	// START OF determine what to index
	
	names := make([]string, 0)

	for k,v := range props_rsp.Map() {

		if k == "wof:name" {
			names = append(names, v.String())
			continue
		}

		if strings.HasPrefix(k, "name:"){

			for _, n := range v.Array(){
				names = append(names, n.String())
			}
		}
	}
	
	doc := map[string]interface{}{
		"names": names,
	}

	// END OF determine what to index
	
	t1 := time.Now()
	
	defer func(){
		log.Printf("Time to index '%s', %v\n", str_id, time.Since(t1))
	}()
	
	return index.Index(str_id, doc)
}
