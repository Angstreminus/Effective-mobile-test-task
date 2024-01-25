package service

import (
	"github.com/Angstreminus/Effective-mobile-test-task/config"
	"github.com/Angstreminus/Effective-mobile-test-task/internal/apperrors"
	"github.com/Angstreminus/Effective-mobile-test-task/internal/dto"
	"github.com/Angstreminus/Effective-mobile-test-task/internal/entity"
	"github.com/Angstreminus/Effective-mobile-test-task/internal/repository"
	uuid "github.com/satori/go.uuid"
)

type UserService struct {
	Repo        *repository.UserRepository
	Config *config.Config
	// TODO repository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{
		Repo: userRepo,
	}
}

func (us *UserService) GetAllUsers(cursor string, limit int, filters map[string]string) ([]entity.User, string, error) {
	return us.Repo.GetAllUsers(cursor, limit, filters)
}

func (us *UserService) DeleteUser(uuid uuid.UUID) apperrors.AppError {
	return us.Repo.DeleteUser(uuid)
}

func (us *UserService) EditUser(uuid uuid.UUID, toEdit dto.UserEditRequest) (*entity.User, apperrors.AppError) {
	var user entity.User
	user.ID = uuid
	user.Name = toEdit.Name
	user.Surname = toEdit.Surname
	user.Patronymic = toEdit.Patronymic
	user.Gender = toEdit.Gender
	user.Nationality = toEdit.Nationality
	user.Age = toEdit.Age
	return us.Repo.EditUser(&user)
}

func (us *UserService) CreateUser(toCreate dto.UserRequest) (*entity.User, apperrors.AppError) {
	var user entity.User

	user.Name = toCreate.Name
	user.Surname = toCreate.Surname
	user.Patronymic = toCreate.Patronymic
	user.Age, err := GetAge(user.Name, us.Repo.)

	return us.Repo.CreateUser(&user)
}
