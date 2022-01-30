package users

import (
	"errors"
	"gorm.io/gorm"
)

type UserStorage interface {
	Save(dto *UserCreateDTO) (*User, error)
	Lookup(username string) ([]User, error)

	// Login using basic auth method.
	Login(loginDTO LoginDTO) (*User, error)
}

type userStorage struct {
	*gorm.DB
}

func newUserStorage(db *gorm.DB) *userStorage {
	return &userStorage{
		DB: db,
	}
}

// method implementations.

// Save creates a new user.
func (us *userStorage) Save(dto *UserCreateDTO) (*User, error) {
	u := dto.ToUser()
	err := us.Create(u).Error

	if err != nil {
		return nil, err
	}
	return u, nil
}

func (us *userStorage) Lookup(username string) ([]User, error) {
	var res []User
	err := us.Where("name LIKE ?", "%"+username+"%").Find(&res).Error

	if err != nil {
		return nil, err
	}
	return res, nil
}

func (us *userStorage) Login(dto LoginDTO) (*User, error) {
	var u User
	var err error
	if dto.Username != "" {
		err = us.Where("username=?", dto.Username).First(&u).Error
	}

	if dto.Email != "" {
		err = us.Where("email=?", dto.Email).First(&u).Error
	}

	if err != nil {
		return nil, err
	}

	if u.ID == 0 {
		return nil, errors.New("user not found")
	}

	if !doPasswordsMatch(u.Password, dto.Password) {
		return nil, errors.New("incorrect password")
	}

	return &u, nil
}
