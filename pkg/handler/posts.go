package handler

import (
	"API"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) createPost(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	var input API.Posts
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Post.Create(userId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getAllPosts(c *gin.Context) {

}

func (h *Handler) getPostById(c *gin.Context) {

}

func (h *Handler) updatePost(c *gin.Context) {

}

func (h *Handler) deletePost(c *gin.Context) {

}
