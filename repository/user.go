package repository

import (
	"errors"
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
	Get(User model.User) (*model.User, error)
	Exist(User model.User) *model.User
	ExistByUserID(id string) *model.User
	ExistByMobile(mobile string) *model.User
	Add(User model.User) (*model.User, error)
	Edit(User model.User) (bool, error)
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
	var total int64
	repo.DB.Where(&user).Find(&user).Count(&total)
	if total > 0 {
		return &user, nil
	} else {
		return nil, errors.New("用户不存在")
	}
}

func (repo *UserRepository) Exist(user model.User) *model.User {
	var total int64
	repo.DB.Where(&user).Find(&user).Count(&total)
	if total > 0 {
		return &user
	}
	return nil
}

func (repo *UserRepository) ExistByMobile(mobile string) *model.User {
	var user model.User
	var total int64
	repo.DB.Find(&user).Where("mobile = ?", mobile).Count(&total)
	if total > 0 {
		return &user
	}
	return nil
}

func (repo *UserRepository) ExistByUserID(id string) *model.User {
	var user model.User
	var total int64
	repo.DB.Where("user_id = ?", id).First(&user).Count(&total)
	if total > 0 {
		return &user
	}
	return nil
}

func (repo *UserRepository) Add(user model.User) (*model.User, error) {
	if exist := repo.Exist(user); exist != nil {
		return nil, fmt.Errorf("用户已存在")
	}
	err := repo.DB.Create(&user).Error
	if err != nil {
		return nil, fmt.Errorf("用户注册失败")
	}
	return &user, nil
}

func (repo *UserRepository) Edit(user model.User) (bool, error) {
	err := repo.DB.Model(&user).Where("user_id = ?", user.UserID).Updates(map[string]interface{}{
		"username":  user.Username,
		"mobile":    user.Mobile,
		"address":   user.Address,
		"password":  user.Password,
		"role_name": user.RoleName,
	}).Error
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
