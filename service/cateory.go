package service

import (
	"errors"
	"fmt"
	"manageSystem/model"
	"manageSystem/repository"

	uuid "github.com/satori/go.uuid"
)

type CategorySrv interface {
	Get(Category model.Category) (*model.Category, error)
	ExistByName(Category model.Category) *model.Category
	Add(Category model.Category) (*model.Category, error)
	EditDescByName(Category model.Category) (bool, error)
}

type CategoryService struct {
	Repo repository.CategoryRepoInterface
}

func (srv *CategoryService) Get(category model.Category) (*model.Category, error) {
	if category.CategoryId == "" {
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
	category.CategoryId = uuid.NewV4().String()
	if category.CategoryId == "" {
		category.CategoryDesc = "请填写描述信息"
	}
	return srv.Repo.Add(category)
}

func (srv *CategoryService) EditDescByName(category model.Category) (bool, error) {
	if category.CategoryName == "" {
		return false, fmt.Errorf("参数错误")
	}
	exist := srv.Repo.ExistByName(category)
	if exist == nil {
		return false, errors.New("参数错误")
	}
	exist.CategoryDesc = category.CategoryDesc
	return srv.Repo.EditDescByName(category)
}
