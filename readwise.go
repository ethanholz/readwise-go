package readwise

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

type ReadwiseInstance struct {
	key  string
	http *http.Client
}

// Creates a new Readwise Instance
// By default uses READWISE_KEY environment variable
func New() *ReadwiseInstance {
	instance := &ReadwiseInstance{
		key:  os.Getenv("READWISE_KEY"),
		http: &http.Client{},
	}

	return instance
}

// Returns a highlight list with a given book id
func (instance *ReadwiseInstance) GetHighlightsForBook(id int) (*HighightList, *error) {
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

	body, err := ioutil.ReadAll(resp.Body)
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

func (instance *ReadwiseInstance) GetHighlightList() (*HighightList, *error) {
	URL := "https://readwise.io/api/v2/highlights/"
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return nil, &err
	}
	req.Header.Add("Authorization", "Token "+instance.key)
	resp, err := instance.http.Do(req)
	if err != nil {
		return nil, &err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, &err
	}

	var list HighightList
	err = json.Unmarshal(body, &list)
	if err != nil {
		return nil, &err
	}

	return &list, nil
}

func (instance *ReadwiseInstance) GetBookList() (*BookList, *error) {
	URL := "https://readwise.io/api/v2/books/"
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return nil, &err
	}
	req.Header.Add("Authorization", "Token "+instance.key)
	resp, err := instance.http.Do(req)
	if err != nil {
		return nil, &err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, &err
	}

	var list BookList
	err = json.Unmarshal(body, &list)
	if err != nil {
		panic(err)
	}

	return &list, nil
}

func (instance *ReadwiseInstance) CreateHighlight(highlight NewHighlight) *error {
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
