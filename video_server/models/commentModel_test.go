package models

import (
	"fmt"
	"strconv"
	"testing"
)

func TestAddComment(t *testing.T) {
	err := AddComment("333f24a1249311eb880cf0761c81050b", uint(6), "不愧是功夫皇帝")
	if err != nil {
		t.Errorf("addComment error : %v", err)
	}
}

func TestDeleteComment(t *testing.T) {
	err := DeleteComment("6757044524cc11eb8d9df0761c81050b")
	if err != nil {
		t.Errorf("deleteComment error : %v", err)
	}
}

func TestAddComments(t *testing.T) {
	for i := 10; i < 30; i++ {
		content := "好身手" + strconv.FormatInt(int64(i), 10)
		err := AddComment("333f24a1249311eb880cf0761c81050b", uint(i), content)
		if err != nil {
			t.Errorf("addComment error : %v", err)
		}
	}
}

func TestGetCommentsByVideoId(t *testing.T) {
	comments, err := GetCommentsByVideoId("333f24a1249311eb880cf0761c81050b", 0, 100)
	if err != nil {
		t.Errorf("getcomments error : %v", err)
	}
	for _, comment := range comments {
		fmt.Println(comment.AuthorName, " [", comment.createDate.String(), "]")
		fmt.Println(comment.Content)
		fmt.Println("")
	}
}

