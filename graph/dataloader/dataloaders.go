package dataloader

import (
	"context"
	"go_server/graph/model"
	"go_server/internal/fetchers"
	"go_server/service"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vikstrous/dataloadgen"
)

type ctxKey string

const loadersKey = ctxKey("dataloaders")

// We use dataloaders to batch nested resource requests
// and avoid N+1 queries
type Loaders struct {
	PostsLoader  *dataloadgen.Loader[int, []*model.Post]
	AlbumsLoader *dataloadgen.Loader[int, []*model.Album]
	PhotosLoader *dataloadgen.Loader[int, []*model.Photo]
}

func NewLoaders(hc *http.Client) *Loaders {
	pr := &PostReader{hc}
	ar := &AlbumReader{hc}
	phr := &PhotoReader{hc}
	return &Loaders{
		PostsLoader:  dataloadgen.NewLoader(pr.getPostsForUsers, dataloadgen.WithWait(5*time.Millisecond)),
		AlbumsLoader: dataloadgen.NewLoader(ar.getAlbumsForUsers, dataloadgen.WithWait(5*time.Millisecond)),
		PhotosLoader: dataloadgen.NewLoader(phr.getPhotosForAlbums, dataloadgen.WithWait(5*time.Millisecond)),
	}
}

func Middleware(hc *http.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		loaders := NewLoaders(hc)

		ctx := context.WithValue(c.Request.Context(), loadersKey, loaders)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func For(ctx context.Context) *Loaders {
	return ctx.Value(loadersKey).(*Loaders)
}

// ResourceReader is an interface
// used by the dataloaders
// to fetch resources from the API
type ResourceReader interface {
	GetHttpClient() *http.Client
	GetEdgeID(resource any) int
	GetQueryField() string
	GetPath() string
}

// GetNestedResource is a generic function that fetches
// resources of the parent resource
// and returns them in the order of the ids parameter
func GetNestedResource[T any](rdr ResourceReader, ids []int) ([][]*T, []error) {
	rurl := service.BASE_URL + rdr.GetPath()
	v := url.Values{}
	for _, id := range ids {
		v.Add(rdr.GetQueryField(), strconv.Itoa(id))
	}
	rurl = rurl + "?" + v.Encode()
	posts, err := fetchers.FetchResourceSlice[T](rdr.GetHttpClient(), rurl)
	if err != nil {
		errs := make([]error, 0, len(ids))
		for range ids {
			errs = append(errs, err)
		}
		return nil, errs
	}
	// Resources are returned sorted by edge id
	// we need to chunk them and
	// return them in the order of request
	posMap := make(map[int]int, len(ids))
	for i, id := range ids {
		posMap[id] = i
	}
	// Start with slice size of request ids
	chunks := make([][]*T, len(ids))
	for _, post := range posts {
		// Place post in the original order position
		pos := posMap[rdr.GetEdgeID(post)]
		chunks[pos] = append(chunks[pos], post)
	}

	return chunks, nil

}
