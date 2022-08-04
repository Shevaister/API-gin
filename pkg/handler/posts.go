package handler

import (
	"API"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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

type getAllPostsResponse struct {
	Data []API.Posts `json:"data"`
}

func (h *Handler) getAllPosts(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	posts, err := h.services.Post.GetAll(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllPostsResponse{
		Data: posts,
	})
}

func (h *Handler) getPostById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("post_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	post, err := h.services.Post.GetById(userId, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, post)
}

func (h *Handler) updatePost(c *gin.Context) {

}

func (h *Handler) deletePost(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("post_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.services.Post.Delete(userId, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
