package service

import (
	"fmt"
	"go_server/graph/model"
	"go_server/internal/fetchers"
	"net/http"
)

type Photo struct {
	HC *http.Client
}

func (s *Photo) GetAll() ([]*model.Photo, error) {
	url := BASE_URL + "photos"
	return fetchers.FetchResourceSlice[model.Photo](s.HC, url)
}

func (s *Photo) Get(id int) (*model.Photo, error) {
	url := BASE_URL + fmt.Sprintf("photos/%d", id)
	return fetchers.FetchResource[model.Photo](s.HC, url)
}
