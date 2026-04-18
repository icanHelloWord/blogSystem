package handle

import (
	"blogSystem/dto"
	"blogSystem/model"
	"blogSystem/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PostsHandler struct {
	postsService service.PostsService
}

func NewPostsHandler(postsService service.PostsService) *PostsHandler {
	return &PostsHandler{
		postsService: postsService,
	}
}

func (postsHandler *PostsHandler) AddPosts(c *gin.Context) {
	var req dto.AddPosts
	userID := c.GetInt("userID")
	req.UserID = userID
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	if err := postsHandler.postsService.AddPosts(&req); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "添加博客成功",
	})
}

func (postsHandler *PostsHandler) PostsList(c *gin.Context) {
	var req dto.PostsListReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	var posts []model.Posts
	posts = postsHandler.postsService.PostsList(&req)

	c.JSON(http.StatusOK, gin.H{
		"data": posts,
	})
}

func (postsHandler *PostsHandler) DeletePosts(c *gin.Context) {
	postsIdstr := c.Param("id")
	postsId, err := strconv.Atoi(postsIdstr)

	userId := c.GetInt("userID")
	if err == nil {
		deleteErr := postsHandler.postsService.DeletePosts(postsId, userId)

		if deleteErr != nil {
			c.JSON(http.StatusBadRequest, deleteErr.Error())
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": "删除成功",
			})
		}
	}

}

func (postsHandler *PostsHandler) UpdatePosts(c *gin.Context) {
	var req dto.UpdatePostsReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	userId := c.GetInt("userID")
	req.UserId = userId
	if err := postsHandler.postsService.UpdatePosts(&req); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "博客更新成功",
	})
}
