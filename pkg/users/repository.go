package users

import "gorm.io/gorm"

type repo struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) Interface {
	return &repo{
		DB: db,
	}
}

func (r *repo) CreateUserProfile(username string) error {
	panic(0)
}

func (r *repo) DeleteUserProfile(username string) error {
	panic(0)
}
