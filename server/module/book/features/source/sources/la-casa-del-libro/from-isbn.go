package lacasadellibro

import (
	"encoding/json"
	"fmt"
	"media/module/book"
	"net/http"
)

func (s *source) FromISBN(isbn string) (*book.SourceData, error) {
	resp, err := http.Get(fmt.Sprintf("https://api.empathy.co/search/v1/query/cdl/isbnsearch?internal=true&query=%s&origin=url:external&start=0&rows=24&instance=cdl&lang=es&scope=desktop&currency=EUR&store=ES", isbn))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var res response
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, err
	}
	if res.Catalog.NumFound == 0 || len(res.Catalog.Content) == 0 {
		return nil, nil
	}

	return toData(*res.Catalog.Content[0]), nil
}
