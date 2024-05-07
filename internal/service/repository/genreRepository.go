package repository

import (
	"context"
	"fmt"

	"github.com/ountyrio/gql-bookstore/internal/service/model"
	"gorm.io/gorm"
)

type GenreRepository interface {
	GetById(ctx context.Context, ID int) (*model.Genre, error)
	Query(ctx context.Context, query *Query) ([]model.Genre, int64, error)
	Create(ctx context.Context, genre *model.Genre) (*model.Genre, error)
	Update(ctx context.Context, genre *model.Genre) (*model.Genre, error)
	Delete(ctx context.Context, ID int) error
}

type GenreRepositoryImpl struct {
	db *gorm.DB
}

func (r *GenreRepositoryImpl) GetById(ctx context.Context, ID int) (*model.Genre, error) {
	var res *model.Genre
	err := r.db.Where("id = ?", ID).Find(&res).Error
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *GenreRepositoryImpl) Query(ctx context.Context, query *Query) ([]model.Genre, int64, error) {
	res := []model.Genre{}
	var count int64
	offset := (query.Page - 1) * query.Limit

	err := r.db.Order(fmt.Sprintf("%s %s", query.Field, query.Order)).Limit(query.Limit).Offset(offset).Find(&res).Error
	if err != nil {
		return nil, count, err
	}
	r.db.Table("genres").Count(&count)
	return res, count, nil
}

func (r *GenreRepositoryImpl) Create(ctx context.Context, genre *model.Genre) (*model.Genre, error) {
	err := r.db.Create(&genre).Error
	if err != nil {
		return nil, err
	}
	return genre, nil
}

func (r *GenreRepositoryImpl) Update(ctx context.Context, genre *model.Genre) (*model.Genre, error) {
	var res *model.Genre
	err := r.db.Where("id = ?", genre.ID).Find(&res).Error
	if err != nil {
		return nil, err
	}
	res.Name = genre.Name
	r.db.Save(&res)
	return res, nil
}

func (r *GenreRepositoryImpl) Delete(ctx context.Context, ID int) error {
	query := r.db.Delete(&model.Genre{}, ID)

	if query.Error != nil {
		return query.Error
	}

	if query.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func NewGenreRepository(db *gorm.DB) *GenreRepositoryImpl {
	return &GenreRepositoryImpl{
		db: db,
	}
}
