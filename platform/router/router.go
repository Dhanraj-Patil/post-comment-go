package router

import (
	_ "github.com/Dhanraj-Patil/post-comment-go/docs"
	"github.com/Dhanraj-Patil/post-comment-go/platform/handler"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func New() *gin.Engine {

	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.POST("/post", handler.SubmitPostHandler)
	router.GET("/posts", handler.GetPostsHandler)
	router.GET("/post/:id", handler.GetPostWithCommentsHandler)

	return router
}
