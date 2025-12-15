package service

import (
	"errors"
	"projectbase/app/model"
	"projectbase/app/repository"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	userRepo repository.UserRepository
	roleRepo repository.RoleRepository
}

func NewUserService(
	userRepo repository.UserRepository,
	roleRepo repository.RoleRepository,
) UserService {
	return &userService{
		userRepo: userRepo,
		roleRepo: roleRepo,
	}
}

func (s *userService) GetUsers(limit, page int) ([]model.User, int64, error) {
	offset := (page - 1) * limit
	return s.userRepo.FindAll(limit, offset)
}

func (s *userService) GetUserByID(id string) (*model.User, error) {
	return s.userRepo.FindByID(id)
}

func (s *userService) CreateUser(user *model.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(user.PasswordHash),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return err
	}

	user.PasswordHash = string(hashedPassword)
	user.ID = uuid.NewString()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	user.IsActive = true

	return s.userRepo.Create(user)
}

func (s *userService) UpdateUser(id string, user *model.User) error {
	existing, err := s.userRepo.FindByID(id)
	if err != nil {
		return err
	}

	existing.FullName = user.FullName
	existing.Email = user.Email
	existing.IsActive = user.IsActive
	existing.UpdatedAt = time.Now()

	return s.userRepo.Update(existing)
}

func (s *userService) DeleteUser(id string) error {
	return s.userRepo.Delete(id)
}

func (s *userService) AssignRole(userID, roleID string) error {
	_, err := s.roleRepo.GetByID(roleID)
	if err != nil {
		return errors.New("role not found")
	}

	return s.userRepo.UpdateRole(userID, roleID)
}
