package service

import (
	"errors"
	"projectbase/app/repository"
	"projectbase/utils"
	"projectbase/app/model"
)

type AuthService struct {
	userRepo repository.UserRepository
	roleRepo repository.RoleRepository
}

func NewAuthService(userRepo repository.UserRepository, roleRepo repository.RoleRepository) *AuthService {
	return &AuthService{
		userRepo: userRepo,
		roleRepo: roleRepo,
	}
}

func (s *AuthService) Login(email, password string) (string, *model.User, error) {
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return "", nil, errors.New("user tidak ditemukan")
	}

	if !user.IsActive {
		return "", nil, errors.New("akun tidak aktif")
	}

	if !utils.CheckPasswordHash(password, user.PasswordHash) {
		return "", nil, errors.New("password salah")
	}

	role, err := s.roleRepo.GetByID(user.RoleID)
	if err != nil {
		return "", nil, errors.New("role tidak ditemukan")
	}

	// 🔑 SIMPLE RBAC (SESUI TUGAS)
	var permissions []string

	switch role.Name {
	case "Admin":
		permissions = []string{
			"user:manage",
			"achievement:read",
			"achievement:create",
			"achievement:update",
			"achievement:delete",
			"achievement:verify",
		}

	case "Dosen Wali":
		permissions = []string{
			"achievement:read",
			"achievement:verify",
		}

	case "Mahasiswa":
		permissions = []string{
			"achievement:read",
			"achievement:create",
			"achievement:update",
			"achievement:delete",
		}
	}

	token, err := utils.GenerateJWT(user.ID, role.Name, permissions)
	if err != nil {
		return "", nil, err
	}

	return token, user, nil
}
