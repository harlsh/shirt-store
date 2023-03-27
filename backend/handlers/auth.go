package handlers

import (
	"fmt"
	"net/http"
	"shirt-store/models"
	"shirt-store/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RegisterInput struct {
	Email     string
	FirstName string
	LastName  string
	Password  string
}

type Login struct {
	Email    string
	Password string
}

type Server struct {
	db *gorm.DB
}

func NewServer(db *gorm.DB) *Server {
	return &Server{db}
}

func (server *Server) Login(c *gin.Context) {

	var login Login

	if err := c.BindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if login.Email == "" || login.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required fields"})
		return
	}

	var dbUser models.User

	result := server.db.Model(&models.User{}).Where("email = ?", login.Email).First(&dbUser)
	if result.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid credentials"})
		return
	}

	err := dbUser.VerifyPassword(login.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid credentials"})
		return
	}

	type AuthKeyUser struct {
		models.User
		AuthKey string
	}

	token, err := utils.GenerateToken(dbUser)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't create a token"})
		return
	}

	c.SetCookie("jwt", token, 3600*24*30, "", "", false, true)
	userWithAuthKey := AuthKeyUser{dbUser, token}
	c.JSON(http.StatusOK, userWithAuthKey)

}



func (server *Server) Register(c *gin.Context) {

	var newUser RegisterInput

	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if newUser.Email == "" || newUser.FirstName == "" || newUser.LastName == "" || newUser.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required fields"})
		return
	}

	if len(newUser.Email) < 5 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email is too short!"})
		return
	}

	if len(newUser.Password) < 8 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password is too short!"})
		return
	}

	dbUser := models.User{
		Email:     newUser.Email,
		FirstName: newUser.FirstName,
		LastName:  newUser.LastName,
		Password:  newUser.Password,
		Role:      0,
	}

	if dbUser.Exists() {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already exists!"})
		return
	}

	err := dbUser.HashPassword()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password!"})
		return
	}

	dbUser.Save()
	c.JSON(http.StatusOK, dbUser)

}

func (server *Server) Me(c *gin.Context) {
	fmt.Println("Hello")
	user, err := utils.CurrentUser(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
}

func status(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "You are logged in"})
}
