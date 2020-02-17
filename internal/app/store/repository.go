package store

import "github.com/wmrsmile2018/GG/internal/app/model"

//UserRepository ...
type UserRepository interface {
	CreateUser(*model.User) error
	FindByEmail(string) (*model.User, error)
	Find(string) (*model.User, error)
	FindByChat(id_chat string) (map[*model.User]bool , error)
}