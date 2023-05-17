package user

import (
	"source-base-go/entity"
	"source-base-go/infrastructure/repository"

	"gorm.io/gorm"
)

type UserRepository interface {
	WithTrx(trxHandle *gorm.DB) repository.UserRepository
	GetUserProfile(userId int) (*entity.User, error)
}

type UseCase interface {
	GetUserProfile(userId int) (*entity.User, error)
}
