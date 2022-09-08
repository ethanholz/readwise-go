package readwise

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

type Instance struct {
	key  string
	http *http.Client
}

// Creates a new Readwise Instance
// By default uses READWISE_KEY environment variable
func New() *Instance {
	instance := &Instance{
		key:  os.Getenv("READWISE_KEY"),
		http: &http.Client{},
	}

	return instance
}

type listable interface {
	BookList | HighightList
}

func getList[T listable](instance *Instance) (*T, *error) {
	var list T
	var endpoint string
	switch any(list).(type) {
	case BookList:
		endpoint = "https://readwise.io/api/v2/books/"
	case HighightList:
		endpoint = "https://readwise.io/api/v2/highlights/"
	}
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, &err
	}
	req.Header.Add("Authorization", "Token "+instance.key)
	resp, err := instance.http.Do(req)
	if err != nil {
		return nil, &err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, &err
	}

	err = json.Unmarshal(body, &list)
	if err != nil {
		return nil, &err
	}

	return &list, nil
}

func (instance *Instance) GetHighlightList() (*HighightList, *error) {
	return getList[HighightList](instance)
}

func (instance *Instance) GetBookList() (*BookList, *error) {
	return getList[BookList](instance)
}

// Returns a highlight list with a given book id
func (instance *Instance) GetHighlightsForBook(id int) (*HighightList, *error) {
	URL := "https://readwise.io/api/v2/highlights/?"
	payload := url.Values{}
	payload.Add("book_id", strconv.Itoa(id))
	req, _ := http.NewRequest("GET", URL+payload.Encode(), nil)
	req.Header.Add("Authorization", "Token "+instance.key)
	req.Header.Add("Content-Type", "application/json")
	resp, err := instance.http.Do(req)
	if err != nil {
		return nil, &err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, &err
	}
	var highlightList HighightList
	err = json.Unmarshal(body, &highlightList)
	if err != nil {
		return nil, &err
	}

	return &highlightList, nil
}

func (instance *Instance) CreateHighlight(highlight NewHighlight) *error {
	URL := "https://readwise.io/api/v2/books/"
	body, err := json.Marshal(highlight)
	if err != nil {
		return &err
	}
	req, err := http.NewRequest("POST", URL, bytes.NewBuffer(body))
	if err != nil {
		return &err
	}
	req.Header.Add("Authorization", "Token "+instance.key)
	resp, err := instance.http.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	return nil
}
