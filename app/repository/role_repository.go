package repository

import (
	"projectbase/app/model"

	"gorm.io/gorm"
)

type roleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) RoleRepository {
	return &roleRepository{db}
}

func (r *roleRepository) GetByID(id string) (*model.Role, error) {
	var role model.Role
	err := r.db.First(&role, "id = ?", id).Error
	return &role, err
}

func (r *roleRepository) GetPermissions(roleID string) ([]model.Permission, error) {
	var perms []model.Permission
	err := r.db.Raw(`
		SELECT p.* 
		FROM role_permissions rp
		JOIN permissions p ON p.id = rp.permission_id
		WHERE rp.role_id = ?
	`, roleID).Scan(&perms).Error

	return perms, err
}
