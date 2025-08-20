package repo

import (
	"errors"
	"github.com/diiineeei/mvc-sample/model"
)

type UserRepo struct {
	// Simula um banco de dados em memória
	users map[int]model.User
}

func NewUserRepo() *UserRepo {
	return &UserRepo{
		users: make(map[int]model.User),
	}
}

func (r *UserRepo) Get(id string) (model.User, error) {
	// Simulação: retorna um usuário fixo para id "1"
	if id == "1" {
		return model.User{ID: 1, Name: "John Doe"}, nil
	}
	return model.User{}, errors.New("user not found")
}

func (r *UserRepo) Save(user model.User) error {
	if user.ID == 0 {
		return errors.New("invalid user ID")
	}
	r.users[user.ID] = user
	return nil
}
