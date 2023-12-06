package repository

import (
	"fmt"
	"manageSystem/model"
	"manageSystem/query"
	"manageSystem/utils"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

type UserRepoInterface interface {
	List(req *query.ListQuery) (Users []*model.User, err error)
	GetTotal(req *query.ListQuery) (total int64, err error)
	Get(Banner model.User) (*model.User, error)
	Exist(Banner model.User) *model.User
	ExistByUserID(id string) *model.User
	ExistByMobile(mobile string) *model.User
	Add(Banner model.User) (*model.User, error)
	Edit(Banner model.User) (bool, error)
	Delete(id string) (bool, error)
}

func (repo *UserRepository) List(req *query.ListQuery) (users []*model.User, err error) {
	db := repo.DB
	limit, offset := utils.Page(req.PageSize, req.Page) // 分页

	if err := db.Order("user_id desc").Limit(limit).Offset(offset).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (repo *UserRepository) GetTotal(req *query.ListQuery) (total int64, err error) {
	var users []model.User
	db := repo.DB

	if err := db.Find(&users).Count(&total).Error; err != nil {
		return total, err
	}
	return total, nil
}

func (repo *UserRepository) Get(user model.User) (*model.User, error) {
	if err := repo.DB.Where(&user).Find(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *UserRepository) Exist(user model.User) *model.User {
	repo.DB.Find(&user).Where("username = ?", user.Username)
	if &user != nil {
		return &user
	}
	return nil
}

func (repo *UserRepository) ExistByMobile(mobile string) *model.User {
	var count int
	var user model.User
	repo.DB.Find(&user).Where("mobile = ?", mobile)
	if count > 0 {
		return &user
	}
	return nil
}

func (repo *UserRepository) ExistByUserID(id string) *model.User {
	var user model.User
	repo.DB.Where("user_id = ?", id).First(&user)
	return &user
}

func (repo *UserRepository) Add(user model.User) (*model.User, error) {
	if exist := repo.Exist(user); exist != nil {
		return nil, fmt.Errorf("用户注册已存在")
	}
	err := repo.DB.Create(&user).Error
	if err != nil {
		return nil, fmt.Errorf("用户注册失败")
	}
	return &user, nil
}

func (repo *UserRepository) Edit(user model.User) (bool, error) {
	err := repo.DB.Model(&user).Where("user_id=?", user.UserID).Updates(map[string]interface{}{"username": user.Username, "mobile": user.Mobile, "address": user.Address}).Error
	//err := repo.DB.Save(&user).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (repo *UserRepository) Delete(id string) (bool, error) {
	var u = &model.User{}
	err := repo.DB.Model(&u).Where("user_id=?", id).Delete(&u).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
