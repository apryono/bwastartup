package user

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Repository interface {
	Save(user User) (User, error)
	FindByEmail(email string) (User, error)
	FindByID(ID int) (User, error)
	Update(user User) (User, error)
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

	if err := r.db.Where("email = ?", email).Find(&user).Error; err != nil {
		logrus.Error(err)
		return user, err
	}

	return user, nil
}

func (r *repository) FindByID(ID int) (User, error) {
	user := User{}

	if err := r.db.Where("id = ?", ID).Find(&user).Error; err != nil {
		logrus.Error(err)
		return user, err
	}

	return user, nil
}

func (r *repository) Update(user User) (User, error) {
	if err := r.db.Save(&user).Error; err != nil {
		logrus.Error(err)
		return user, err
	}

	return user, nil
}
