package user

import (
	"source-base-go/config"
	"source-base-go/entity"
	"source-base-go/infrastructure/repository/util"
)

type Service struct {
	userRepository UserRepository
	roleRepo       RoleRepository
	verifier       Verifier
}

func NewService(userRepository UserRepository, roleRepo RoleRepository, verifier Verifier) *Service {
	return &Service{
		userRepository: userRepository,
		roleRepo:       roleRepo,
		verifier:       verifier,
	}
}

func (s *Service) GetUserProfile(userId int) (*entity.User, error) {
	return s.userRepository.GetUserProfile(userId)
}

func (s *Service) Login(username string, password string) (*entity.TokenPair, *entity.User, error) {
	//Get user by user name
	user, err := s.userRepository.FindByUsername(username)
	if err != nil {
		return nil, nil, entity.ErrInternalServerError
	}

	if user == nil {
		return nil, nil, entity.ErrUsernameNotExists
	}

	//Vefiry password
	passwordValidate := user.ValidatePassword(password)
	if !passwordValidate {
		return nil, nil, entity.ErrInvalidPassword
	}

	//Find all roles of user
	listRole, err := s.roleRepo.FindAllRolesOfUser(user.Id)
	listRoleStr := []string{}
	for _, role := range listRole {
		listRoleStr = append(listRoleStr, role.Code)
	}
	//Cache userdata into redis
	err = s.verifier.CacheUserData(user, listRoleStr, config.GetInt("jwt.accessMaxAge"))
	if err != nil {
		return nil, nil, entity.ErrInternalServerError
	}

	//Generate token
	acessToken, err := util.GenerateAccessToken(user)
	if err != nil {
		return nil, nil, entity.ErrInternalServerError
	}
	refreshToken, err := util.GenerateRefreshToken(user)
	if err != nil {
		return nil, nil, entity.ErrInternalServerError
	}

	tokenPair := &entity.TokenPair{
		Token:        acessToken,
		RefreshToken: refreshToken,
	}

	return tokenPair, user, nil
}
