package user

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Repository interface {
	Save(user User) (User, error)
	FindByEmail(email string) (User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Save(user User) (User, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindByEmail(email string) (User, error) {
	user := User{}

	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		logrus.Error(err)
		return user, err
	}

	return user, nil
}
