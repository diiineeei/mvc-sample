package providers

import (
	"github.com/diiineeei/mvc-sample/model"
	"github.com/diiineeei/mvc-sample/repo"
)

type UserProvider struct {
	repo *repo.UserRepo
}

func NewUserProvider(repo *repo.UserRepo) *UserProvider {
	return &UserProvider{repo: repo}
}

func (p *UserProvider) GetUser(id string) (model.User, error) {
	user, err := p.repo.Get(id)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (p *UserProvider) CreateUser(user model.User) error {
	return p.repo.Save(user)
}
