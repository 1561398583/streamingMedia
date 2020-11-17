package models

import (
	"gorm.io/gorm"
	"video_server/utils"
	"time"
)

type VideoInfo struct {
	ID        string `gorm:"primarykey"`
	AuthorId uint
	Name string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func AddVideo(name string, authorId uint)  error{
	id, err := utils.GetUUID()
	if err != nil {
		return err
	}
	videoInfo := VideoInfo{ID: id, Name: name, AuthorId: authorId}
	r := DB.Create(&videoInfo)
	if r.Error != nil {
		return r.Error
	}
	return nil
}

func DeleteVideo(id string)  error{
	result := DB.Where("id=?", id).Delete(&VideoInfo{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func SetVideoName(id string, newName string) error {
	r := DB.Model(&VideoInfo{}).Where("id = ?", id).Update("name", newName)
	if r.Error != nil {
		return r.Error
	}
	return nil
}

func GetVideoById(id string) (*VideoInfo,error) {
	videoInfo := VideoInfo{}
	r := DB.Where("id = ?", id).First(&videoInfo)
	if r.Error != nil {
		return nil, r.Error
	}
	return &videoInfo, nil
}

func GetVideoByAuthorId(author_id uint) ([]VideoInfo,error) {
	var videoInfos []VideoInfo
	r := DB.Where("author_id = ?", author_id).Find(&videoInfos)
	if r.Error != nil {
		return nil, r.Error
	}
	return videoInfos, nil
}


