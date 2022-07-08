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

type HighightList struct {
	Count    int         `json:"count"`
	Next     string      `json:"next"`
	Previous string      `json:"previous"`
	Results  []Highlight `json:"results"`
}

type Highlight struct {
	Text          string `json:"text"`
	ID            int    `json:"id"`
	Location      int    `json:"location"`
	LocationType  string `json:"location_type"`
	HighlightedAt string `json:"highlighted_at"`
	BookID        int    `json:"book_id"`
}
