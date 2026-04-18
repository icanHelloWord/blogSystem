package main

import (
	"blogSystem/config"
	database "blogSystem/dataBase"
	"blogSystem/handle"
	"blogSystem/middle"
	"blogSystem/repository"
	"blogSystem/service"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal("加载配置失败:", err)
	}

	database.InitDB(&config.Database)

	db := database.GetDB()

	router := gin.Default()

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo, config)
	userHandler := handle.NewUserHandle(userService)

	postsRepo := repository.NewPostsRepository(db)
	postsService := service.NewPostsService(postsRepo)
	postsHandler := handle.NewPostsHandler(postsService)

	//注册
	router.POST("/register", userHandler.Register)
	//登录
	router.POST("/login", userHandler.Login)
	//文章列表
	router.POST("/postsList", postsHandler.PostsList)

	//验证token
	router.Use(middle.JWTAuthMiddleware(config))

	//新增博客
	router.POST("/addPosts", postsHandler.AddPosts)
	//更新博客
	router.POST("/updatePosts", postsHandler.UpdatePosts)
	//删除博客
	router.DELETE("/deletePosts/:id", postsHandler.DeletePosts)

	router.Run(":8080")

	//4. 文章管理功能
	// 实现文章的创建功能，只有已认证的用户才能创建文章，创建文章时需要提供文章的标题和内容。
	// 实现文章的读取功能，支持获取所有文章列表和单个文章的详细信息。
	// 实现文章的更新功能，只有文章的作者才能更新自己的文章。
	// 实现文章的删除功能，只有文章的作者才能删除自己的文章。
	// 5. 评论功能
	// 实现评论的创建功能，已认证的用户可以对文章发表评论。
	// 实现评论的读取功能，支持获取某篇文章的所有评论列表。
}
