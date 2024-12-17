// Code generated by mirip; DO NOT EDIT.
// github.com/gmhafiz/mirip

package repository

import (
	"context"
	"github.com/gmhafiz/go8/internal/domain/author"
)

// AuthorMock is a mock implementation of Author.
type AuthorMock struct {
	CreateFunc func(ctx context.Context, a *author.CreateRequest) (*author.Schema, error)
	DeleteFunc func(ctx context.Context, authorID uint64) error
	ListFunc   func(ctx context.Context, f *author.Filter) ([]*author.Schema, int, error)
	ReadFunc   func(ctx context.Context, id uint64) (*author.Schema, error)
	UpdateFunc func(ctx context.Context, toAuthor *author.UpdateRequest) (*author.Schema, error)
}

func (m *AuthorMock) Create(ctx context.Context, a *author.CreateRequest) (*author.Schema, error) {
	return m.CreateFunc(ctx, a)
}

func (m *AuthorMock) Delete(ctx context.Context, authorID uint64) error {
	return m.DeleteFunc(ctx, authorID)
}

func (m *AuthorMock) List(ctx context.Context, f *author.Filter) ([]*author.Schema, int, error) {
	return m.ListFunc(ctx, f)
}

func (m *AuthorMock) Read(ctx context.Context, id uint64) (*author.Schema, error) {
	return m.ReadFunc(ctx, id)
}

func (m *AuthorMock) Update(ctx context.Context, toAuthor *author.UpdateRequest) (*author.Schema, error) {
	return m.UpdateFunc(ctx, toAuthor)
}

// SearcherMock is a mock implementation of Searcher.
type SearcherMock struct {
	SearchFunc func(ctx context.Context, f *author.Filter) ([]*author.Schema, int, error)
}

func (m *SearcherMock) Search(ctx context.Context, f *author.Filter) ([]*author.Schema, int, error) {
	return m.SearchFunc(ctx, f)
}