package repos

import "github.com/kuriouslabs/godo/models"

type UserRepoAdaptor interface {
}

type UserRepo struct {
	adaptor UserRepoAdaptor
}

func NewUserRepo(adaptor UserRepoAdaptor) *UserRepo {
	return &UserRepo{
		adaptor: adaptor,
	}
}

func (r *UserRepo) AuthenticateUserPassword(uid string, password string) bool {
	return uid == "chase" && password == "pw"
}

func (r *UserRepo) ByID(uid string) (*models.User, error) {
	u := models.NewUser("user123", uid)
	return &u, nil
}
