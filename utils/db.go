package pg

import (
	"github.com/ountyrio/gql-bookstore/internal/service/model"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func OpenDBConnection() (*gorm.DB, error) {
	var err error

	db, err := gorm.Open(postgres.Open(viper.GetString("database.url")), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&model.Author{}, &model.Genre{}, &model.Book{})

	return db, nil
}
