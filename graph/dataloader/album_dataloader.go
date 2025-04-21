package dataloader

import (
	"context"
	"go_server/graph/model"
	"log"
	"net/http"
)

type AlbumReader struct {
	hc *http.Client
}

func (pr *AlbumReader) GetHttpClient() *http.Client {
	return pr.hc
}

func (pr *AlbumReader) GetEdgeID(resource any) int {
	album, ok := resource.(*model.Album)
	if !ok {
		log.Fatalf("PostReader: expected *model.Album, got %T", resource)
	}
	return album.UserID
}

func (pr *AlbumReader) GetQueryField() string {
	return "userId"
}

func (pr *AlbumReader) GetPath() string {
	return "albums"
}

func (pr *AlbumReader) getAlbumsForUsers(_ context.Context, userIds []int) ([][]*model.Album, []error) {
	log.Printf("Albums dataloader called with userIds(%v):%v\n", len(userIds), userIds)

	return GetNestedResource[model.Album](pr, userIds)
}

func LoadUsersAlbums(ctx context.Context, ids []int) ([][]*model.Album, error) {
	loaders := For(ctx)
	return loaders.AlbumsLoader.LoadAll(ctx, ids)
}

func LoadUserAlbums(ctx context.Context, id int) ([]*model.Album, error) {
	loaders := For(ctx)
	posts, err := loaders.AlbumsLoader.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	return posts, nil
}
