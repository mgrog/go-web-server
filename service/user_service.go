package service

import (
	"fmt"
	"go_server/graph/model"
	"go_server/internal/fetchers"
	"net/http"
)

type User struct {
	HC *http.Client
}

func (s *User) GetAll() ([]*model.User, error) {
	url := BASE_URL + "users"
	return fetchers.FetchResourceSlice[model.User](s.HC, url)
}

func (s *User) Get(id int) (*model.User, error) {
	url := BASE_URL + fmt.Sprintf("users/%d", id)
	return fetchers.FetchResource[model.User](s.HC, url)
}

func (s *User) GetPosts(id int) ([]*model.Post, error) {
	url := BASE_URL + fmt.Sprintf("users/%d/posts", id)
	return fetchers.FetchResourceSlice[model.Post](s.HC, url)
}

func (s *User) GetAlbums(id int) ([]*model.Album, error) {
	url := BASE_URL + fmt.Sprintf("users/%d/albums", id)
	return fetchers.FetchResourceSlice[model.Album](s.HC, url)
}
