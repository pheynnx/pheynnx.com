package blog

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/pheynnx/pheynnx.com/internal/cache"
	"github.com/pheynnx/pheynnx.com/internal/orbit"
)

// blog router state
//
// extends chi router and orbit
type Router struct {
	postCache *cache.PostCache
	*chi.Mux
	*orbit.Orbit
}

func newRouter(postCache *cache.PostCache) *Router {
	return &Router{
		postCache: postCache, Mux: chi.NewRouter(),
	}
}

func Routes(postCache *cache.PostCache) *Router {
	router := newRouter(postCache)

	h := Handler{router}

	router.Get("/", h.getBlogIndex())

	router.Route("/blog", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, "/", http.StatusMovedPermanently)
		})
		r.Get("/{slug}", h.getBlogBySlug())
	})

	return router
}
