package goodreads

import (
	"fmt"
	"log"
	"media/module/book"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/samber/lo"
)

func (s *source) FromISBN(isbn string) (*book.SourceData, error) {
	res, err := http.Get(fmt.Sprintf("https://www.goodreads.com/search?q=%s", isbn))
	if err != nil {
		return nil, err
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	title := doc.Find("h1.Text__title1").Text()
	description := doc.Find(".BookPageMetadataSection__description .DetailsLayoutRightParagraph__widthConstrained").Text()
	authors := doc.Find(".BookPageMetadataSection__contributor .ContributorLink__name").Map(func(i int, s *goquery.Selection) string {
		return s.Text()
	})
	genres := lo.Filter(doc.Find(".BookPageMetadataSection__genres .Button__labelItem").Map(func(i int, s *goquery.Selection) string {
		return s.Text()
	}), func(item string, _ int) bool {
		return item != "" && !strings.Contains(item, "...more")
	})
	img, _ := doc.Find(".BookCover__image img").Attr("src")
	var images []string
	if img != "" {
		images = append(images, img)
	}

	return &book.SourceData{
		Source: book.SourceGoodreads,
		Metadata: &book.Metadata{
			Title:       title,
			Description: description,
			Authors:     authors,
			ISBN:        isbn,
			Publisher:   "",
			Tags:        genres,
		},
		Images: images,
	}, nil
}
