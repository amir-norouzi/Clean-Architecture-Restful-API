package repositories

import (
	"RestApi/database/mysql"
	"RestApi/models/mysql/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type UserRepository struct {
	DB *gorm.DB
}

// create repository and start connection with database
func Create() *UserRepository {
	db := mysql.Init()
	return &UserRepository{DB: db}
}

// create a user
func (repo *UserRepository) CreateUser(c *gin.Context) {
	var user models.User
	err := c.BindJSON(&user)
	if err != nil {
		return
	}
	err = models.CreateUser(repo.DB, &user)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, user)
}

// get all users
func (repo *UserRepository) GetAllUsers(c *gin.Context) {
	var user []models.User
	err := models.GetUsers(repo.DB, &user)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, user)
}
