package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type User struct {
	gorm.Model
	Email     string
	FirstName string
	LastName  string
	Password  string `json:"-"`
	Role      uint   `json:"-"`
}


func (u *User) UnmarshalJSON(data []byte) error {
    type userAlias User // Define an alias of the struct to avoid infinite recursion
    aux := &struct {
        *userAlias
        Password string `json:"password"`
    }{
        userAlias: (*userAlias)(u),
    }
    if err := json.Unmarshal(data, &aux); err != nil {
        return err
    }
    if aux.Password != "" {
        u.Password = aux.Password
    }
    return nil
}

var db *gorm.DB

func main() {

	var err error
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&User{})

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/register", RegisterHandle)
	r.POST("/login", LoginHandle)
	r.GET("/users", UsersHandle)
	r.Run() // listen and serve on 0.0.0.0:8080
}

func UsersHandle(c *gin.Context) {
	var users []User
	db.Find(&users)
	c.JSON(http.StatusOK, users)
}

func LoginHandle(c *gin.Context) {
	type Login struct {
		Email    string
		Password string
	}

	var login Login

	if err := c.BindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if login.Email == "" || login.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required fields"})
		return
	}

	var dbUser User

	result := db.Model(&User{}).Where("email = ?", login.Email).First(&dbUser)
	if result.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid credentials"})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(login.Password))
	if dbUser.Email != login.Email || err != nil {
		fmt.Println("Wrong password")
		fmt.Print(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid credentials"})
		return
	}

	//save the login session somehow
	c.JSON(http.StatusOK, dbUser)

}

func RegisterHandle(c *gin.Context) {
	var newUser User

	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("Password is ", newUser.Password)
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

	email := newUser.Email

	var count int64
	db.Model(&User{}).Where("email = ?", email).Count(&count)
	if count > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already exists!"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password!"})
		return
	}

	dbUser := User{
		Email:     newUser.Email,
		FirstName: newUser.FirstName,
		LastName:  newUser.LastName,
		Password:  string(hashedPassword),
		Role:      0,
	}

	db.Create(&dbUser)

	c.JSON(http.StatusOK, dbUser)

}
