package controller

import (
	"blogSystem/dto"
	"blogSystem/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CommentController struct {
	commentService service.CommentService
}

func NewCommentHandler(commentService service.CommentService) *CommentController {
	return &CommentController{commentService: commentService}
}

func (co *CommentController) AddComment(c *gin.Context) {
	var req dto.AddCommentReq
	userID := c.GetInt("userID")
	req.UserID = userID
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	if err := co.commentService.AddComment(&req); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "添加评论成功",
	})
}

func (co *CommentController) CommentList(c *gin.Context) {
	var req dto.CommentListReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	comments := co.commentService.CommentList(&req)
	c.JSON(http.StatusOK, gin.H{
		"message":  "查询评论成功",
		"comments": comments,
	})
}

func (co *CommentController) Delete(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	userID := c.GetInt("userID")
	if err := co.commentService.Delete(idInt, userID); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "删除评论成功",
	})
}
