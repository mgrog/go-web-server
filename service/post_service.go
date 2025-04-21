package service

import (
	"fmt"
	"go_server/graph/model"
	"go_server/internal/fetchers"
	"net/http"
)

type Post struct {
	HC *http.Client
}

func (s *Post) GetAll() ([]*model.Post, error) {
	url := BASE_URL + "posts"
	return fetchers.FetchResourceSlice[model.Post](s.HC, url)
}

func (s *Post) Get(id int) (*model.Post, error) {
	url := BASE_URL + fmt.Sprintf("posts/%d", id)
	return fetchers.FetchResource[model.Post](s.HC, url)
}
