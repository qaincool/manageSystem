package service

import (
	"errors"
	"fmt"
	"manageSystem/model"
	"manageSystem/repository"
	"time"
)

type VideoRepoSrv interface {
	List() (videos []*model.Video, err error)
	GetTotal() (total int64, err error)
	Get(video *model.Video) (*model.Video, error)
	Exist(video *model.Video) *model.Video
	Add(video *model.Video) (*model.Video, error)
	Edit(video model.Video) (*model.Video, error)
	Delete(video model.Video) (bool, error)
	GetVideoByTag(tags []string) ([]*model.Video, error)
	GetVideoByCategory(categories []string) ([]*model.Video, error)
}

type VideoService struct {
	Repo repository.VideoRepoInterface
}

func (srv *VideoService) List() (Videos []*model.Video, err error) {
	return srv.Repo.List()
}

func (srv *VideoService) GetTotal() (total int64, err error) {
	return srv.Repo.GetTotal()
}

func (srv *VideoService) Get(video *model.Video) (*model.Video, error) {
	return srv.Repo.Get(video)
}

func (srv *VideoService) GetVideoByTag(tags []string) ([]*model.Video, error) {
	if len(tags) > 0 {
		return srv.Repo.GetVideoByTag(tags)
	} else {
		return nil, errors.New("tag信息为空")
	}
}

func (srv *VideoService) GetVideoByCategory(categories []string) ([]*model.Video, error) {
	if len(categories) > 0 {
		return srv.Repo.GetVideoByCategory(categories)
	} else {
		return nil, errors.New("类别信息为空")
	}
}

func (srv *VideoService) Exist(video *model.Video) *model.Video {
	return srv.Repo.Exist(video)
}

func (srv *VideoService) Add(video *model.Video) (*model.Video, error) {
	if video.VideoPath == "" || video.VideoName == "" {
		return nil, errors.New("请输入视频名称或存放地址")
	}
	nameResult := srv.Repo.Exist(video)
	if nameResult != nil {
		return nil, errors.New("视频名称或地址已经存在")
	}
	video.CreateTime = time.Now()

	return srv.Repo.Add(video)
}

func (srv *VideoService) Edit(video model.Video) (*model.Video, error) {
	if video.VideoId == 0 {
		return nil, fmt.Errorf("参数错误")
	}
	exist := srv.Repo.ExistByID(video.VideoId)
	if exist == nil {
		return nil, errors.New("参数错误")
	}
	exist.VideoName = video.VideoName
	exist.VideoDetail = video.VideoDetail
	exist.VideoIntro = video.VideoIntro
	exist.VideoPath = video.VideoPath
	exist.VideoTag = video.VideoTag
	return srv.Repo.Edit(exist)
}

func (srv *VideoService) Delete(video model.Video) (bool, error) {
	if video.VideoId == 0 {
		return false, errors.New("参数错误")
	}
	v := srv.Exist(&video)
	if v == nil {
		return false, errors.New("参数错误")
	}
	return srv.Repo.Delete(video)
}
