package handle

import (
	"blogSystem/dto"
	"blogSystem/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandle struct {
	userService service.UserService
}

func NewUserHandle(userService service.UserService) *UserHandle {

	return &UserHandle{
		userService: userService,
	}
}

// 用户登录
func (handle *UserHandle) Login(c *gin.Context) {

	var loginReq dto.UserLoginReq

	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	token, err := handle.userService.Login(c, &loginReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
		"type":  "Bearer",
	})
}

func (handle *UserHandle) Register(c *gin.Context) {
	// 可校验用户名合法性
	var req dto.UserLoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := handle.userService.Register(c, &req); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "用户创建成功，请登录",
	})

}
