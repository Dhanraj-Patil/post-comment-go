package service

import (
	"time"

	"github.com/Dhanraj-Patil/post-comment-go/platform/model"
	"github.com/Dhanraj-Patil/post-comment-go/platform/repository"
	"github.com/gomarkdown/markdown"
)

func SavePost(post model.Post) error {
	post.CreateAt = time.Now()
	post.PostData = string(markdown.ToHTML([]byte(post.PostData), nil, nil))
	if err := repository.Save(post); err != nil {
		return err
	}
	return nil
}

func GetPosts() ([]model.Post, error) {
	posts, err := repository.GetAll()
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func GetThread(postId string) (model.Thread, error) {
	var thread model.Thread
	post, err := repository.GetById(postId)
	if err != nil {
		return model.Thread{}, err
	}
	comments, err := repository.GetCommentsByThread(postId)
	if err != nil {
		return model.Thread{}, err
	}

	thread.Post = post
	thread.Comments = comments
	return thread, nil
}
