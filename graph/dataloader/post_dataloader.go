package dataloader

import (
	"context"
	"go_server/graph/model"
	"log"
	"net/http"
)

type PostReader struct {
	hc *http.Client
}

func (pr *PostReader) GetHttpClient() *http.Client {
	return pr.hc
}

func (pr *PostReader) GetEdgeID(resource any) int {
	post, ok := resource.(*model.Post)
	if !ok {
		log.Fatalf("PostReader: expected *model.Post, got %T", resource)
	}
	return post.UserID
}

func (pr *PostReader) GetQueryField() string {
	return "userId"
}

func (pr *PostReader) GetPath() string {
	return "posts"
}

func (pr *PostReader) getPostsForUsers(_ context.Context, userIds []int) ([][]*model.Post, []error) {
	log.Printf("Posts dataloader called with userIds(%v):(%v)\n", len(userIds), userIds)

	return GetNestedResource[model.Post](pr, userIds)
}

func LoadUsersPosts(ctx context.Context, ids []int) ([][]*model.Post, error) {
	loaders := For(ctx)
	return loaders.PostsLoader.LoadAll(ctx, ids)
}

func LoadUserPosts(ctx context.Context, id int) ([]*model.Post, error) {
	loaders := For(ctx)
	posts, err := loaders.PostsLoader.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	return posts, nil
}
