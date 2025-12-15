package service

import "projectbase/app/model"

type UserService interface {
	GetUsers(limit, page int) ([]model.User, int64, error)
	GetUserByID(id string) (*model.User, error)
	CreateUser(user *model.User) error
	UpdateUser(id string, user *model.User) error
	DeleteUser(id string) error
	AssignRole(userID, roleID string) error
}
