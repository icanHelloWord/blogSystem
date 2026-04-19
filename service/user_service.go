package service

import (
	"blogSystem/config"
	"blogSystem/dto"
	"blogSystem/middle"
	"blogSystem/model"
	"blogSystem/repository"
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	userRepo repository.UserRepository
	config   *config.Config
}

func NewUserService(userRepo repository.UserRepository, config *config.Config) UserService {
	return &userService{
		userRepo: userRepo,
		config:   config,
	}
}

func (userService *userService) Login(c *gin.Context, req *dto.UserLoginReq) (string, error) {

	userRepo := userService.userRepo

	user, err := userRepo.GetUserByUserName(req.Username)

	if user == nil || err != nil {

		return "", errors.New("用户不存在")
	}
	password := user.Password

	if err := bcrypt.CompareHashAndPassword([]byte(password), []byte(req.Password)); err != nil {
		return "", errors.New("密码错误")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  user.ID,
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, err := token.SignedString([]byte(userService.config.JWT.Secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (userService *userService) Register(c *gin.Context, userRegister *dto.UserLoginReq) error {
	db := middle.GetDBFromContext(c)
	tx := db.Begin()
	//有多表操作的话，开启事务
	userRepo := userService.userRepo
	userCheck, err := userRepo.GetUserByUserName(userRegister.Username)
	if err != nil {
		tx.Rollback()
		return err
	}

	if userCheck != nil {
		tx.Rollback()
		return errors.New("用户名已经存在")
	}

	var user model.User
	user.Username = userRegister.Username
	user.Email = userRegister.Email
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userRegister.Password), bcrypt.DefaultCost)
	if err != nil {
		tx.Rollback()
		return err
	}
	user.Password = string(hashedPassword)
	errCreate := userRepo.Create(&user)
	if err != nil {
		tx.Rollback()
		return errCreate
	}
	tx.Commit()
	return nil
}
