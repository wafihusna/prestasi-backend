package service

import (
	"projectbase/app/repository"
)

type RBACService struct {
	roleRepo repository.RoleRepository
}

func NewRBACService(roleRepo repository.RoleRepository) *RBACService {
	return &RBACService{roleRepo}
}

// CheckPermission memverifikasi apakah user memiliki izin tertentu
func (s *RBACService) CheckPermission(roleID string, requiredPermission string) (bool, error) {
	perms, err := s.roleRepo.GetPermissions(roleID)
	if err != nil {
		return false, err
	}

	for _, p := range perms {
		permString := p.Resource + ":" + p.Action
		if permString == requiredPermission {
			return true, nil
		}
	}

	return false, nil
}