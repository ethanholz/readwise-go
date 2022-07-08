package readwise

import (
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
func (instance *ReadwiseInstance) GetHighlightsForBook(id int) HighightList {
	URL := "https://readwise.io/api/v2/highlights/?"
	payload := url.Values{}
	payload.Add("book_id", strconv.Itoa(id))
	req, _ := http.NewRequest("GET", URL+payload.Encode(), nil)
	req.Header.Add("Authorization", "Token "+instance.key)
	req.Header.Add("Content-Type", "application/json")
	resp, err := instance.http.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	var highlightList HighightList
	err = json.Unmarshal(body, &highlightList)
	if err != nil {
		panic(err)
	}

	return highlightList
}

func (instance *ReadwiseInstance) GetHighlightList() HighightList {
	URL := "https://readwise.io/api/v2/highlights/"
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Authorization", "Token "+instance.key)
	resp, err := instance.http.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var list HighightList
	err = json.Unmarshal(body, &list)
	if err != nil {
		panic(err)
	}

	return list
}

func (instance *ReadwiseInstance) GetBookList() BookList {
	URL := "https://readwise.io/api/v2/books/"
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Authorization", "Token "+instance.key)
	resp, err := instance.http.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var list BookList
	err = json.Unmarshal(body, &list)
	if err != nil {
		panic(err)
	}

	return list
}
