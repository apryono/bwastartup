package user

import (
	"errors"
	"fmt"

	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
	Login(input LoginInput) (User, error)
	IsEmailAvailable(input CheckEmailInput) (bool, error)
	SaveAvatar(ID int, fileLocation string) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{
		repository: repository,
	}
}

func (s *service) RegisterUser(input RegisterUserInput) (User, error) {
	user := User{}
	user.Name = input.Name
	user.Occupation = input.Occupation
	user.Email = input.Email
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, nil
	}

	user.PasswordHash = string(passwordHash)
	user.Role = "user"

	response, err := s.repository.Save(user)
	if err != nil {
		return response, err
	}

	return response, nil
}

func (s *service) Login(input LoginInput) (User, error) {
	email := input.Email
	pwd := input.Password

	fmt.Println("email", email)

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		logrus.Error(err)
		return user, err
	}

	fmt.Println("usrEmail", user.Email)

	if user.ID == 0 {
		return user, errors.New("Data Not Found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(pwd))
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *service) IsEmailAvailable(input CheckEmailInput) (bool, error) {
	email := input.Email

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return false, err
	}

	if user.ID == 0 {
		return true, nil
	}

	return false, nil
}

func (s *service) SaveAvatar(ID int, fileLocation string) (User, error){
	user, err := s.repository.FindByID(ID)
	if err != nil {
		return user, nil
	}

	user.AvatarFileName = fileLocation

	updatedUser, err := s.repository.Update(user)
	if err != nil {
		return updatedUser, err
	}

	return updatedUser, nil
}