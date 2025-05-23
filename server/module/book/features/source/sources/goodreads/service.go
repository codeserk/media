package goodreads

import (
	"fmt"
	"io"
	"net/http"
)

func Search(query string) (any, error) {
	res, err := http.Get(fmt.Sprintf("https://www.goodreads.com/search?q=%s", query))
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	fmt.Printf("%s", body)

	return nil, nil
}
