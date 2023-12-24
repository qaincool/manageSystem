package repository

import (
	"fmt"
	"gorm.io/gorm"
	"manageSystem/model"
)

type VideoRepository struct {
	DB *gorm.DB
}

type VideoRepoInterface interface {
	List() (videos []*model.Video, err error)
	GetTotal() (total int64, err error)
	Get(video *model.Video) (*model.Video, error)
	Exist(video *model.Video) *model.Video
	ExistByID(id int) *model.Video
	Add(video *model.Video) (*model.Video, error)
	Edit(video *model.Video) (*model.Video, error)
	Delete(video model.Video) (bool, error)
	GetVideoByTag(tags []string) ([]*model.Video, error)
	GetVideoByCategory(categories []string) ([]*model.Video, error)
}

func (repo *VideoRepository) List() (videos []*model.Video, err error) {
	db := repo.DB
	//limit, offset := utils.Page(req.PageSize, req.Page) // 分页

	if err := db.Preload("Category").Find(&videos).Error; err != nil {
		return nil, err
	}
	return videos, nil
}

func (repo *VideoRepository) GetTotal() (total int64, err error) {
	var videos []model.Video
	db := repo.DB
	if err := db.Preload("Category").Find(&videos).Count(&total).Error; err != nil {
		return total, err
	}
	return total, nil
}

func (repo *VideoRepository) Get(video *model.Video) (*model.Video, error) {
	if err := repo.DB.Preload("Category").Find(&video).Error; err != nil {
		return nil, err
	}
	return video, nil
}

// GetVideoByTag 根据tag信息查询视频
func (repo *VideoRepository) GetVideoByTag(tags []string) ([]*model.Video, error) {
	var tagsVideos []*model.Video
	for _, tag := range tags {
		var video *model.Video
		if err := repo.DB.Preload("Category").Find(&video).Where("video_tag LIKE ?", "%"+tag+"%").Error; err != nil {
			continue
		}
		tagsVideos = append(tagsVideos, video)
	}
	return tagsVideos, nil
}

// GetVideoByCategory 根据视频所属类别查询
// TODO: 视频类别查询
func (repo *VideoRepository) GetVideoByCategory(categories []string) ([]*model.Video, error) {
	var categoryVideos []*model.Video
	for _, category := range categories {
		var video *model.Video
		if err := repo.DB.Preload("Category", "category_name = ?", category).Find(&video).Error; err != nil {
			continue
		}
		categoryVideos = append(categoryVideos, video)
	}
	return categoryVideos, nil

}

func (repo *VideoRepository) Exist(video *model.Video) *model.Video {
	var total int64
	repo.DB.Where(&video).Find(&video).Count(&total)
	if total > 0 {
		return video
	}
	return nil
}

func (repo *VideoRepository) ExistByID(id int) *model.Video {
	var video model.Video
	var total int64
	repo.DB.First(&video).Where("video_id = ?", id).Count(&total)
	if total > 0 {
		return &video
	}
	return nil
}

func (repo *VideoRepository) Add(video *model.Video) (*model.Video, error) {
	if exist := repo.Exist(video); exist != nil {
		return nil, fmt.Errorf("视频已存在")
	}
	err := repo.DB.Create(&video).Error
	if err != nil {
		return nil, fmt.Errorf("视频创建失败")
	}
	return video, nil
}

func (repo *VideoRepository) Edit(video *model.Video) (*model.Video, error) {
	err := repo.DB.Model(&video).Select("*").Where("video_id=?", video.VideoId).Updates(map[string]interface{}{
		"video_name":   video.VideoName,
		"video_intro":  video.VideoIntro,
		"video_path":   video.VideoPath,
		"video_detail": video.VideoDetail,
		"video_tag":    video.VideoTag,
	}).Error
	//err := repo.DB.Save(&user).Error
	if err != nil {
		return nil, err
	}
	return video, nil
}

func (repo *VideoRepository) Delete(video model.Video) (bool, error) {
	err := repo.DB.Model(&video).Where("video_id=?", video.VideoId).Delete(&video).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
