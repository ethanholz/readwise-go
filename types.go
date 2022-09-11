package readwise

type Book struct {
	ID              int    `json:"id"`
	Title           string `json:"title"`
	Author          string `json:"author"`
	Category        string `json:"category"`
	Source          string `json:"source"`
	NumHighlights   int    `json:"num_highlights"`
	LastHighlightAt string `json:"last_highlight_at"`
	Updated         string `json:"updated"`
	CoverImageURL   string `json:"cover_image_url"`
	HighlightsURL   string `json:"highlights_url"`
	SourceURL       string `json:"source_url"`
}

type BookList struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []Book `json:"results"`
}

type BookReq struct {
	BookID int `json:"book_id"`
}

type HighlightList struct {
	Count    int         `json:"count"`
	Next     string      `json:"next"`
	Previous string      `json:"previous"`
	Results  []Highlight `json:"results"`
}

type NewHighlights struct {
	Highlights []NewHighlight `json:"highlights"`
}

type NewHighlight struct {
	Text          string `json:"text"`
	Title         string `json:"title"`
	Author        string `json:"author"`
	ImageURL      string `json:"image_url"`
	SourceURL     string `json:"source_url"`
	Category      string `json:"category"`
	Note          string `json:"note"`
	Location      string `json:"location"`
	LocationType  string `json:"location_type"`
	HighlightedAt string `json:"highlighted_at"`
	HighlightURL  string `json:"highlight_url"`
}

type Highlight struct {
	Text          string   `json:"text"`
	ID            int      `json:"id"`
	Note          string   `json:"note"`
	Location      int      `json:"location"`
	LocationType  string   `json:"location_type"`
	HighlightedAt string   `json:"highlighted_at"`
	BookID        int      `json:"book_id"`
	URL           string   `json:"url"`
	Color         string   `json:"color"`
	Updated       string   `json:"updated"`
	Tags          []string `json:"tags"`
}

type listable interface {
	BookList | HighlightList
}

type tagable interface {
	BookTags | HighlightTags
}

type Tags struct {
	Count   int    `json:"count"`
	Next    string `json:"next"`
	Results []Tag  `json:"results"`
}

type Tag struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type (
	BookTags      Tags
	HighlightTags Tags
)
