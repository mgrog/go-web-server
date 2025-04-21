package fetchers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
)

func FetchResource[T any](httpClient *http.Client, url string) (*T, error) {
	segments := strings.Split(url, ".com/")
	path := strings.Join(segments[1:], ".com/")
	log.Printf("Fetching resource from %s", path)
	res, err := httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var resource T
	if err := json.Unmarshal(body, &resource); err != nil {
		return nil, err
	}
	return &resource, nil
}

func FetchResourceSlice[T any](httpClient *http.Client, url string) ([]*T, error) {
	resources := []*T{}
	res, err := FetchResource[[]*T](httpClient, url)
	if err != nil {
		return nil, err
	}
	if res != nil {
		resources = *res
	}
	return resources, nil
}
