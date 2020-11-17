package models

import (
	"fmt"
	"testing"
)

func TestAddVideo(t *testing.T) {
	name := "李连杰1"
	author_id := uint(233)
	err := AddVideo(name, author_id)
	if err != nil {
		t.Errorf("addVideo error : %v\n", err)
	}
}

func TestDeleteVideo(t *testing.T) {
	err := DeleteVideo("05df0970249311eba2c4f0761c81050b")
	if err != nil {
		t.Errorf("delete video error : %v\n", err)
	}
}

func TestSetVideoName(t *testing.T) {
	err := SetVideoName("333f24a1249311eb880cf0761c81050b", "李连杰版霍元甲")
	if err != nil {
		t.Errorf("setVidoeName  error : %v\n", err)
	}
}

func TestGetVideoById(t *testing.T) {
	video, err := GetVideoById("1c6700aa249011eb821ef0761c81050b")
	if err != nil {
		t.Errorf("GetVideoById error : %v\n", err)
	}
	fmt.Println(video.Name)
}

func TestGetVideoByAuthorId(t *testing.T) {
	vidoes, err := GetVideoByAuthorId(123)
	if err != nil {
		t.Errorf("GetVideoByAuthorId error : %v\n", err)
	}
	for _, video := range vidoes{
		fmt.Println(video.Name)
	}
}