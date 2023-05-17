package user

import "source-base-go/entity"

type Service struct {
	userRepository UserRepository
}

func NewService(userRepository UserRepository) *Service {
	return &Service{
		userRepository: userRepository,
	}
}

func (s *Service) GetUserProfile(userId int) (*entity.User, error) {
	return s.userRepository.GetUserProfile(userId)
}
