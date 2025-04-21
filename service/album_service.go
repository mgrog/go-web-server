package service

import (
	"fmt"
	"go_server/graph/model"
	"go_server/internal/fetchers"
	"net/http"
)

type Album struct {
	HC *http.Client
}

func (s *Album) GetAll() ([]*model.Album, error) {
	url := BASE_URL + "albums"
	return fetchers.FetchResourceSlice[model.Album](s.HC, url)
}

func (s *Album) Get(id int) (*model.Album, error) {
	url := BASE_URL + fmt.Sprintf("albums/%d", id)
	return fetchers.FetchResource[model.Album](s.HC, url)
}

func (s *Album) GetPhotos(id int) ([]*model.Photo, error) {
	url := BASE_URL + fmt.Sprintf("albums/%d/photos", id)
	return fetchers.FetchResourceSlice[model.Photo](s.HC, url)
}
