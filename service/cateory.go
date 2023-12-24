package service

import (
	"errors"
	"manageSystem/model"
	"manageSystem/repository"
)

type CategorySrv interface {
	Get(Category model.Category) (*model.Category, error)
	ExistByName(Category model.Category) *model.Category
	Add(Category model.Category) (*model.Category, error)
}

type CategoryService struct {
	Repo repository.CategoryRepoInterface
}

func (srv *CategoryService) Get(category model.Category) (*model.Category, error) {
	if category.CategoryId == 0 {
		return nil, errors.New("参数错误")
	}
	return srv.Repo.Get(category)
}

func (srv *CategoryService) ExistByName(category model.Category) *model.Category {
	return srv.Repo.ExistByName(category)
}

func (srv *CategoryService) Add(category model.Category) (*model.Category, error) {
	result := srv.ExistByName(category)
	if result != nil {
		return nil, errors.New("分类已经存在")
	}
	return srv.Repo.Add(category)
}
