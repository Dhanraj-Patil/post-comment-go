package handler

import (
	"github.com/Dhanraj-Patil/post-comment-go/platform/model"
	"github.com/Dhanraj-Patil/post-comment-go/platform/service"
	"github.com/gin-gonic/gin"
)

// SubmitPostHandler godoc
// @Summary Create a new post
// @Description Submits a new top-level post or comment (if thread is set). Warning: Set thread to "" if creating new post, else put post id as thread if is a comment.
// @Tags posts
// @Accept json
// @Produce json
// @Param post body model.PostRequest true "Post Body"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /post [post]
func SubmitPostHandler(c *gin.Context) {
	var post model.Post

	if err := c.ShouldBindBodyWithJSON(&post); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}

	if err := service.SavePost(post); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"status": "Post created."})
}

// GetPostsHandler godoc
// @Summary Get all top-level posts
// @Description Retrieves all posts where thread is empty (i.e., original posts)
// @Tags posts
// @Produce json
// @Success 200 {array} model.Post
// @Failure 500 {object} map[string]string
// @Router /posts [get]
func GetPostsHandler(c *gin.Context) {

	post, err := service.GetPosts()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, post)
}

// GetPostWithCommentsHandler godoc
// @Summary Get post with comments
// @Description Returns a post and all its comment posts (where thread == post ID)
// @Tags threads
// @Produce json
// @Param id path string true "Post ID"
// @Success 200 {object} model.Thread
// @Failure 500 {object} map[string]string
// @Router /post/{id} [get]
func GetPostWithCommentsHandler(c *gin.Context) {
	postId := c.Param("id")

	thread, err := service.GetThread(postId)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, thread)
}
