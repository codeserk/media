package goodreads

import (
	"fmt"
	"io"
	"media/module/book"
	"net/http"
)

func (s *source) Search(query string) ([]*book.SourceData, error) {
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
