package search

import (
	"context"
	"net/url"

	"github.com/gary23w/metasearch_api/internal/models"
)

type Searcher interface {
	Search(ctx context.Context, req Request) ResultIterator
	ContinueSearch(ctx context.Context, tok Token) ResultIterator
}

type Service interface {
	models.Provider
	Languages(ctx context.Context) ([]Language, error)
	Regions(ctx context.Context) ([]Region, error)

	Searcher
}

type Request struct {
	Query  string
	Lang   LangCode
	Region RegionCode
	Safe   bool
}

type ResultIterator interface {
	models.PagedIterator
	Result() Result
	Token() Token
}

var _ ResultIterator = Empty{}

type Empty struct {
	models.Empty
}

func (Empty) Result() Result {
	return nil
}

func (Empty) Token() Token {
	return nil
}

type Result interface {
	GetURL() *url.URL
	GetTitle() string
	GetDesc() string
}

type ThumbnailResult interface {
	Result
	GetThumbnail() *Image
}

type Token []byte
