package cache

import (
	"slices"
	"sync"

	"github.com/pheynnx/pheynnx.com/internal/database"
	"github.com/pheynnx/pheynnx.com/internal/model"
	"github.com/pheynnx/pheynnx.com/internal/model/dto"
)

// naming of the cache field will change
// also this cache store might be moved under the model package into is respective type
type PostCache struct {
	Mx        sync.Mutex
	PostsMap  map[string]*model.Post
	MetaSlice []*dto.Meta
}

func InitCache(DB *database.DBPool) (*PostCache, error) {

	// grab all posts from database
	posts, err := model.GetPostsFromDB(DB)
	if err != nil {
		return nil, err
	}

	// Sort posts by date
	model.SortPostsByDate(posts)

	metaSlice := []*dto.Meta{}

	postsMap := map[string]*model.Post{}

	for _, p := range posts {

		// convert raw markdown to html
		err := p.ConvertMarkdownToHTML()
		if err != nil {
			return nil, err
		}

		// sort categories by alphabetical order
		slices.Sort(p.Categories)

		if p.Published {
			metaSlice = append(metaSlice, dto.MetaFromPost(p))
		}

		postsMap[p.Slug] = p
	}

	return &PostCache{
		MetaSlice: metaSlice,
		PostsMap:  postsMap,
	}, nil
}

// func (p *PostCache) UpdateCache() {
// 	p.Mx.Lock()
// }
