package model

import (
	"os"
	"testing"
	"time"
)

var comment = Comment{
	FatherId:   0,
	ToUserId:   10091,
	VideoId:    13736,
	FromUserId: 31263,
	Content:    "6666,太厉害了！",
	CreateAt:   time.Now(),
}

func TestMain(m *testing.M) {
	InitDB()
	code := m.Run()
	os.Exit(code)
}

func TestCommentModel_TableName(t *testing.T) {
	tableName := commentModel.TableName()
	if "tb_comment" != tableName {
		t.FailNow()
	}
}

func TestCommentModel_AddCommentAndUpdateCount(t *testing.T) {
	commentModel := NewCommentModel()
	if err := commentModel.AddCommentAndUpdateCount(&comment); err != nil {
		t.Errorf("AddCommentAndUpdateCount is fail, %v", err)
	}
}

func TestCommentModel_QueryCommentById(t *testing.T) {
	commentModel := NewCommentModel()
	var com Comment
	if err := commentModel.QueryCommentById(comment.Id, &com); err != nil {
		t.Errorf("QueryCommentById is fail, %v", err)
	}
	if "6666,太厉害了！" != com.Content {
		t.Errorf("QueryCommentById is fail, no query!")
	}
}

func TestCommentModel_QueryCommentListByVideoId(t *testing.T) {
	commentModel := NewCommentModel()
	var commentList []*Comment
	if err := commentModel.QueryCommentListByVideoId(comment.VideoId, &commentList); err != nil {
		t.Errorf("QueryCommentListByVideoId is fail, %v", err)
	}

	if 0 == len(commentList) {
		t.Errorf("QueryCommentListByVideoId is fail, no query !")
	}

	if len(commentList) > 1 {
		t.Errorf("QueryCommentListByVideoId is fail, query too many !")
	}
}

func TestCommentModel_DeleteCommentAndUpdateCountById(t *testing.T) {
	commentModel := NewCommentModel()
	if err := commentModel.DeleteCommentAndUpdateCountById(comment.Id, 13736); err != nil {
		t.Errorf("DeleteCommentAndUpdateCountById is fail, %v", err)
	}
}
