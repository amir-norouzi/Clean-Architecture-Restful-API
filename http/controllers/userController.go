package controllers

import (
	"RestApi/http/controllers/repositories"
	"github.com/gin-gonic/gin"
)

var UserRepository = repositories.Create()

func GetUsers() gin.HandlerFunc {
	return UserRepository.GetAllUsers
}
