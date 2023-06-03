package user

import (
	"source-base-go/entity"
	"source-base-go/infrastructure/repository"

	"gorm.io/gorm"
)

type Verifier interface {
	CacheUserData(user *entity.User, listRole []string, expiredAt int) error
}

type UserRepository interface {
	WithTrx(trxHandle *gorm.DB) repository.UserRepository
	GetUserProfile(userId int) (*entity.User, error)
	FindByUsername(userName string) (*entity.User, error)
}

type RoleRepository interface {
	FindAllRolesOfUser(userId int) ([]*entity.Role, error)
}

type UseCase interface {
	GetUserProfile(userId int) (*entity.User, error)
	Login(username string, password string) (*entity.TokenPair, *entity.User, error)
}
