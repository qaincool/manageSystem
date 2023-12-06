package repository

import (
	"fmt"
	"manageSystem/model"

	"gorm.io/gorm"
)

type CategoryRepository struct {
	DB *gorm.DB
}

type CategoryRepoInterface interface {
	Get(Category model.Category) (*model.Category, error)
	ExistByName(Category model.Category) *model.Category
	Add(Category model.Category) (*model.Category, error)
	EditDescByName(Category model.Category) (bool, error)
}

func (repo *CategoryRepository) Get(category model.Category) (*model.Category, error) {
	if err := repo.DB.Where(&category).Find(&category).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (repo *CategoryRepository) ExistByName(category model.Category) *model.Category {
	// var count int
	repo.DB.Find(&category).Where("category_name = ?", category.CategoryName)
	// if count > 0 {
	// 	return &category
	// }
	if &category != nil {
		return &category
	}
	return nil
}

func (repo *CategoryRepository) Add(category model.Category) (*model.Category, error) {
	if exist := repo.ExistByName(category); exist != nil {
		return nil, fmt.Errorf("分类已存在")
	}
	err := repo.DB.Create(&category).Error
	if err != nil {
		return nil, fmt.Errorf("分类添加失败")
	}
	return &category, nil
}

func (repo *UserRepository) EditDescByName(category model.Category) (bool, error) {
	err := repo.DB.Model(&category).Where("category_name=?", category.CategoryName).Updates(map[string]interface{}{"category_desc": category.CategoryDesc}).Error
	//err := repo.DB.Save(&user).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
