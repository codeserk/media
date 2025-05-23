package book

type SourceType string

const (
	SourceGoogle         SourceType = "Google"
	SourceLaCasaDelLibro SourceType = "LaCasaDelLibro"
	SourceLibraryThing   SourceType = "LibraryThing"
	SourceStoryGraph     SourceType = "StoryGraph"
	SourceGoodreads      SourceType = "Goodreads"
)

type Source interface {
	Search(query string) ([]*SourceData, error)
	FromISBN(isbn string) (*SourceData, error)
}

type SourceGoogleConfig struct {
	IsEnabled bool   `json:"isEnabled"`
	ApiKey    string `json:"apiKey"`
}

type SourceLaCasaDelLibroConfig struct {
	IsEnabled bool `json:"isEnabled"`
}

type SourceLibraryThingConfig struct {
	IsEnabled bool   `json:"isEnabled"`
	ApiKey    string `json:"apiKey"`
}

type SourceStoryGraphConfig struct {
	IsEnabled bool `json:"isEnabled"`
}

type SourceGoodreadsConfig struct {
	IsEnabled bool `json:"isEnabled"`
}

type SourceConfig struct {
	Google         SourceGoogleConfig         `json:"google"`
	LaCasaDelLibro SourceLaCasaDelLibroConfig `json:"laCasaDelLibro"`
	LibraryThing   SourceLibraryThingConfig   `json:"libraryThing"`
	StoryGraph     SourceStoryGraphConfig     `json:"storyGraph"`
	Goodreads      SourceGoodreadsConfig      `json:"goodreads"`
}

type SourceService interface {
	Search(query string) ([]*SourceData, error)
	FromISBN(isbn ISBN) (SourceMultiData, error)
}

type SourceData struct {
	Source   SourceType
	Metadata *Metadata
	Images   []string
	Original any
}

type SourceMultiData []*SourceData

func (s SourceMultiData) Metadata() Metadata {
	var result Metadata
	for _, source := range s {
		result = result.Merge(source.Metadata)
	}

	return result
}

func (s SourceMultiData) Images() []string {
	for i := len(s) - 1; i >= 0; i-- {
		if len(s[i].Images) > 0 {
			return s[i].Images
		}
	}

	return []string{}
}
