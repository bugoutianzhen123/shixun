package dao

import "project/domain"

type UserDao interface {
	CreateUser(user domain.User) error
	GetUserById(userId uint) (domain.User, error)
	GetUserByName(username string) (domain.User, error)
	ChangePassword(user domain.User) error
	ChangePermission(user domain.User) error
	ChangeName(user domain.User) error
}

func (dao *GORMDAO) CreateUser(user domain.User) error {
	err := dao.db.Create(&user).Error
	return err
}

func (dao *GORMDAO) GetUserById(userId uint) (domain.User, error) {
	user := domain.User{}
	err := dao.db.First(&user, "id = ?", userId).Error
	return user, err
}

func (dao *GORMDAO) GetUserByName(username string) (domain.User, error) {
	user := domain.User{}
	err := dao.db.First(&user, "Name = ?", username).Error
	return user, err
}

func (dao *GORMDAO) ChangePassword(user domain.User) error {
	err := dao.db.Model(&user).Where("Id = ?", user.ID).Update("password", user.Password).Error
	return err
}

func (dao *GORMDAO) ChangePermission(user domain.User) error {
	err := dao.db.Model(&user).Where("Id = ?", user.ID).Update("permission", user.Permission).Error
	return err
}

func (dao *GORMDAO) ChangeName(user domain.User) error {
	err := dao.db.Model(&user).Where("Id = ?", user.ID).Update("name", user.Name).Error
	return err
}
