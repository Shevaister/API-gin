package handler

import (
	"API/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("sign-up", h.signUp)
		auth.POST("sign-in", h.signIn)
	}

	api := router.Group("/api", h.userIdentity)
	{
		posts := api.Group("/posts")
		{
			posts.POST("/", h.createPost)
			posts.GET("/", h.getAllPosts)
			posts.GET("/:post_id", h.getPostById)
			posts.PUT("/:post_id", h.updatePost)
			posts.DELETE("/:post_id", h.deletePost)

			comments := posts.Group(":post_id/comments")
			{
				comments.POST("/", h.createComment)
				comments.GET("/", h.getAllComments)
				comments.GET("/:comment_id", h.getCommentById)
				comments.PUT("/:comment_id", h.updateComment)
				comments.DELETE("/:comment_id", h.deleteComment)
			}
		}
	}
	return router
}
