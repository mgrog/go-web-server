package dataloader

import (
	"context"
	"go_server/graph/model"
	"log"
	"net/http"
)

type PhotoReader struct {
	hc *http.Client
}

func (pr *PhotoReader) GetHttpClient() *http.Client {
	return pr.hc
}

func (pr *PhotoReader) GetEdgeID(resource any) int {
	photo, ok := resource.(*model.Photo)
	if !ok {
		log.Fatalf("PostReader: expected *model.Photo, got %T", resource)
	}
	return photo.AlbumID
}

func (pr *PhotoReader) GetQueryField() string {
	return "albumId"
}

func (pr *PhotoReader) GetPath() string {
	return "photos"
}

func (pr *PhotoReader) getPhotosForAlbums(_ context.Context, albumIds []int) ([][]*model.Photo, []error) {
	log.Printf("Photos dataloader called with albumIds(%v):%v\n", len(albumIds), albumIds)

	return GetNestedResource[model.Photo](pr, albumIds)
}

func LoadAlbumsPhotos(ctx context.Context, ids []int) ([][]*model.Photo, error) {
	loaders := For(ctx)
	return loaders.PhotosLoader.LoadAll(ctx, ids)
}

func LoadAlbumPhotos(ctx context.Context, id int) ([]*model.Photo, error) {
	loaders := For(ctx)
	posts, err := loaders.PhotosLoader.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	return posts, nil
}
