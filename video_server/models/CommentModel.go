package models

import (
	"gorm.io/gorm"
	"time"
	"video_server/utils"
)

type Comment struct {
	ID      string `gorm:"primarykey"`
	VideoId string
	AuthorId uint
	Content string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type CommentAndAuthor struct {
	AuthorName string
	Content string
	createDate time.Time
}

func AddComment(videoId string, authorId uint, content string)  error{
	id, err := utils.GetUUID()
	if err != nil {
		return err
	}
	comment := Comment{ID : id, VideoId: videoId, AuthorId: authorId, Content: content}
	r := DB.Create(&comment)
	if r.Error != nil {
		return r.Error
	}
	return nil
}

func DeleteComment(id string)  error{
	result := DB.Where("id=?", id).Delete(&Comment{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetCommentsByVideoId(videoId string, offeet, limit int) ([]CommentAndAuthor, error) {
	var comments []Comment
	r := DB.Limit(limit).Offset(offeet).Where("video_id=?", videoId).Order("created_at desc").Find(&comments)
	if r.Error != nil {
		return nil, r.Error
	}
	//获取所有comment的author_id并用map去重
	authorMap := make(map[uint]bool)
	for _, comment := range comments {
		authorId := comment.AuthorId
		authorMap[authorId] = true
	}
	//获取这些author的信息
	authorIds := make([]uint, len(authorMap))	// 数组默认长度为map长度,后面append时,不需要重新申请内存和拷贝,效率较高
	j := 0
	for k, _ := range authorMap {
		authorIds[j] = k
		j ++
	}
	//查询这些author的信息
	var authors []User
	r = DB.Where("id IN ?", authorIds).Find(&authors)
	if r.Error != nil {
		return nil, r.Error
	}
	//酱author放入map
	authorMap2 := make(map[uint]User)
	for _, a := range authors {
		authorMap2[a.ID] = a
	}
	//组装CommentAndAuthor
	caa := make([]CommentAndAuthor, len(comments))
	for i:= 0; i < len(comments); i++ {
		authorId := comments[i].AuthorId
		var authorName string
		if a, ok := authorMap2[authorId]; ok {
			authorName = a.Name
		}else {
			authorName = "(已注销)"
		}
		caa[i].AuthorName = authorName
		caa[i].Content = comments[i].Content
		caa[i].createDate = comments[i].CreatedAt
	}
	return caa, nil
}



