package dto

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/pheynnx/pheynnx.com/internal/model"
)

type Meta struct {
	ID          uuid.UUID      `json:"id"`
	Date        time.Time      `json:"date"`
	Slug        string         `json:"slug"`
	Title       string         `json:"title"`
	Series      string         `json:"series"`
	Categories  pq.StringArray `json:"categories"`
	Published   bool           `json:"published"`
	Featured    bool           `json:"featured"`
	PostSnippet string         `db:"post_snippet" json:"post_snippet"`
}

func MetaFromPost(p *model.Post) *Meta {
	return &Meta{
		ID:          p.ID,
		Date:        p.Date,
		Slug:        p.Slug,
		Title:       p.Title,
		Series:      p.Series,
		Categories:  p.Categories,
		Published:   p.Published,
		Featured:    p.Featured,
		PostSnippet: p.PostSnippet,
	}
}
