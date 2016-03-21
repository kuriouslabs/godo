package repos

import (
	"errors"

	"github.com/kuriouslabs/godo/models"
)

var (
	ErrUserNotFound      = errors.New("Cannot find user")
	ErrUserAlreadyExists = errors.New("User already exists")
)

type UserRepoAdaptor interface {
	UserByID(id string) (*models.User, error)
	UserByUsername(username string) (*models.User, error)
	UserExists(id string) bool
	CreateUser(user *models.User) error
	UpdateUser(user *models.User)
}

type UserRepo struct {
	adaptor UserRepoAdaptor
}

func NewUserRepo(adaptor UserRepoAdaptor) *UserRepo {
	return &UserRepo{
		adaptor: adaptor,
	}
}

func (r *UserRepo) AuthenticateUserPassword(username string, password string) *models.User {
	//TODO: Validate password
	if password != "pw" {
		return nil
	}

	user, _ := r.adaptor.UserByUsername(username)
	return user
}

func (r *UserRepo) CreateUser(user *models.User) error {
	return r.adaptor.CreateUser(user)
}

func (r *UserRepo) ByID(uid string) (*models.User, error) {
	user, err := r.adaptor.UserByID(uid)
	return user, err
}
