package lacasadellibro

import (
	"encoding/json"
	"fmt"
	"media/module/book"
	"net/http"
)

func (s *source) Search(query string) ([]*book.SourceData, error) {
	resp, err := http.Get(fmt.Sprintf("https://api.empathy.co/search/v1/query/cdl/search?internal=true&query=%s&origin=url:external&start=0&rows=24&instance=cdl&lang=es&scope=desktop&currency=EUR&store=ES", query))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return nil, nil
}
