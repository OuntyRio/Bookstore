package repository

import (
	"context"
	"fmt"

	"github.com/ountyrio/gql-bookstore/internal/service/model"
	"gorm.io/gorm"
)

type BookRepository interface {
	GetById(ctx context.Context, ID int) (*model.Book, error)
	Query(ctx context.Context, query *Query) ([]model.Book, int64, error)
	Create(ctx context.Context, book *model.Book) (*model.Book, error)
	Update(ctx context.Context, Book *model.Book) (*model.Book, error)
	Delete(ctx context.Context, ID int) error
}

type BookRepositoryImpl struct {
	db *gorm.DB
}

type Query struct {
	Limit int
	Page  int
	Field string
	Order string
}

func (r *BookRepositoryImpl) GetById(ctx context.Context, ID int) (*model.Book, error) {
	var res *model.Book
	err := r.db.Preload("Authors").Where("id = ?", ID).Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *BookRepositoryImpl) Query(ctx context.Context, query *Query) ([]model.Book, int64, error) {
	res := []model.Book{}
	var count int64
	offset := (query.Page - 1) * query.Limit

	err := r.db.Preload("Authors").Order(fmt.Sprintf("%s %s", query.Field, query.Order)).Limit(query.Limit).Offset(offset).Find(&res).Error
	if err != nil {
		return nil, count, err
	}
	r.db.Table("books").Count(&count)
	return res, count, nil
}

func (r *BookRepositoryImpl) Create(ctx context.Context, book *model.Book) (*model.Book, error) {
	err := r.db.Create(&book).Error
	if err != nil {
		return nil, err
	}
	return book, nil
}

func (r *BookRepositoryImpl) Update(ctx context.Context, book *model.Book) (*model.Book, error) {
	var res *model.Book
	err := r.db.Where("id = ?", book.ID).Find(&res).Error
	if err != nil {
		return nil, err
	}
	res.Title = book.Title
	res.GenreID = book.GenreID
	res.Authors = book.Authors
	r.db.Save(&res)

	r.db.Model(&res).Association("Authors").Replace(book.Authors)

	return res, nil
}

func (r *BookRepositoryImpl) Delete(ctx context.Context, ID int) error {
	var res *model.Book
	err := r.db.Where("id = ?", ID).Find(&res).Error
	if err != nil {
		return err
	}

	if res.ID == 0 {
		return gorm.ErrRecordNotFound
	}

	r.db.Model(&res).Association("Authors").Clear()
	err = r.db.Delete(&res, ID).Error
	if err != nil {
		return err
	}
	return nil
}

func NewBookRepository(db *gorm.DB) *BookRepositoryImpl {
	return &BookRepositoryImpl{
		db: db,
	}
}
