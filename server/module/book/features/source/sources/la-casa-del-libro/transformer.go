package lacasadellibro

import (
	"encoding/json"
	"media/module/book"
	"strings"
	"time"

	"github.com/samber/lo"
)

type response struct {
	Catalog struct {
		Content  []*content `json:"content"`
		NumFound int        `json:"numFound"`
	} `json:"catalog"`
}

type content struct {
	BoostId    string   `json:"__boostId"`
	ExternalId string   `json:"__externalId"`
	Id         string   `json:"__id"`
	Images     []string `json:"__images"`
	Name       string   `json:"__name"`
	Prices     struct {
		Current struct {
			Value float64 `json:"value"`
		} `json:"current"`
		Previous struct {
			Value float64 `json:"value"`
		} `json:"previous"`
	} `json:"__prices"`
	URL            string   `json:"__url"`
	Authors        []string `json:"authors"`
	Availability   string   `json:"availability"`
	DateRelease    int64    `json:"dateRelease"`
	Description    string   `json:"description"`
	EAN            string   `json:"ean"`
	Editorial      string   `json:"editorial"`
	Encuadernation string   `json:"encuadernation"`
	ExternalScore  string   `json:"externalScore"`
	ExternalVotes  string   `json:"externalVotes"`
	FreeShipping   bool     `json:"freeShipping"`
	IDAssoc        string   `json:"idAssoc"`
	InternalID     string   `json:"internal_id"`
	IsSigned       bool     `json:"isSigned"`
	ISBN           string   `json:"isbn"`
	Price          struct {
		Current  float64 `json:"current"`
		Previous float64 `json:"previous"`
	} `json:"price"`
	PriceOffer  float64 `json:"priceOffer"`
	ProductType string  `json:"productType"`
	Saleable    string  `json:"saleable"`
	Score       float64 `json:"score"`
	Tagging     struct {
		Add2Cart string `json:"add2cart"`
		Checkout string `json:"checkout"`
		Click    string `json:"click"`
	} `json:"tagging"`
	YearPublication string `json:"yearPublication"`
}

func toData(c content) *book.SourceData {
	originalJSON, _ := json.Marshal(c)

	return &book.SourceData{
		Source: book.SourceLaCasaDelLibro,
		Metadata: &book.Metadata{
			Title:       c.Name,
			Description: c.Description,
			Authors:     lo.Filter(c.Authors, func(author string, _ int) bool { return !strings.Contains(author, "/") }),
			ISBN:        book.ToISBN(c.ISBN),
			EAN:         c.EAN,
			Publisher:   c.Editorial,
			PageCount:   0,
			PublishedAt: time.Unix(c.DateRelease, 0),
		},
		Images:   c.Images,
		Original: string(originalJSON),
	}
}
