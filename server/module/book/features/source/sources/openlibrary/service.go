package openlibrary

import (
	"encoding/json"
	"net/http"
	"net/url"
)

func Search(query string) (any, error) {
	resp, err := http.Get("https://openlibrary.org/search.json?lang=es&q=" + url.QueryEscape(query))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result, nil
}
