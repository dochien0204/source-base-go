package repository

import (
	"log"
	"source-base-go/entity"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepostory(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r UserRepository) WithTrx(trxHanlde *gorm.DB) UserRepository {
	if trxHanlde == nil {
		log.Print("Transaction DB not found")
		return r
	}

	r.db = trxHanlde
	return r
}

func (r UserRepository) GetUserProfile(userId int) (*entity.User, error) {
	user := entity.User{}

	result := r.db.
		Where("id = ?", userId).
		Preload("Status").
		Find(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
