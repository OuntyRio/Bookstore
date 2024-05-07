package repository

import (
	"context"
	"fmt"

	"github.com/ountyrio/gql-bookstore/internal/service/model"
	"gorm.io/gorm"
)

type AuthorRepository interface {
	GetById(ctx context.Context, ID int) (*model.Author, error)
	Query(ctx context.Context, query *Query) ([]model.Author, int64, error)
	GetMany(ctx context.Context, IDs []int) ([]model.Author, error)
	Create(ctx context.Context, author *model.Author) (*model.Author, error)
	Update(ctx context.Context, author *model.Author) (*model.Author, error)
	Delete(ctx context.Context, ID int) error
}

type AuthorRepositoryImpl struct {
	db *gorm.DB
}

func (r *AuthorRepositoryImpl) GetById(ctx context.Context, ID int) (*model.Author, error) {
	var res *model.Author
	err := r.db.Preload("Books").Where("id = ?", ID).Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *AuthorRepositoryImpl) Query(ctx context.Context, query *Query) ([]model.Author, int64, error) {
	res := []model.Author{}
	var count int64
	offset := (query.Page - 1) * query.Limit

	err := r.db.Preload("Books").Order(fmt.Sprintf("%s %s", query.Field, query.Order)).Limit(query.Limit).Offset(offset).Find(&res).Error
	if err != nil {
		return nil, count, err
	}
	r.db.Table("authors").Count(&count)
	return res, count, nil
}

func (r *AuthorRepositoryImpl) GetMany(ctx context.Context, IDs []int) ([]model.Author, error) {
	res := []model.Author{}
	err := r.db.Find(&res, "id in (?)", IDs).Error
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *AuthorRepositoryImpl) Create(ctx context.Context, author *model.Author) (*model.Author, error) {
	err := r.db.Create(&author).Error
	if err != nil {
		return nil, err
	}

	return author, nil
}

func (r *AuthorRepositoryImpl) Update(ctx context.Context, author *model.Author) (*model.Author, error) {
	var res *model.Author
	err := r.db.Where("id = ?", author.ID).Find(&res).Error
	if err != nil {
		return nil, err
	}

	if res.ID == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	res.Firstname = author.Firstname
	res.Lastname = author.Lastname
	r.db.Save(&res)
	return res, nil
}

func (r *AuthorRepositoryImpl) Delete(ctx context.Context, ID int) error {
	var res *model.Author
	err := r.db.Where("id = ?", ID).Find(&res).Error

	if err != nil {
		return err
	}

	if res.ID == 0 {
		return gorm.ErrRecordNotFound
	}

	r.db.Model(&res).Association("Books").Clear()

	err = r.db.Delete(&res, ID).Error
	if err != nil {
		return err
	}

	return nil
}

func NewAuthorRepository(db *gorm.DB) *AuthorRepositoryImpl {
	return &AuthorRepositoryImpl{
		db: db,
	}
}
