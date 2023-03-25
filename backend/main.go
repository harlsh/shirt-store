package main

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)



var db *gorm.DB

const secret = "shirtsarecool"
const userkey = "user"


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

	r.Use(sessions.Sessions("shirt-store-session", cookie.NewStore([]byte(secret))))

	r.POST("/register", RegisterHandle)
	r.POST("/login", LoginHandle)
	r.GET("/logout", LogoutHandle)
	r.GET("/users", UsersHandle)

	private := r.Group("/private")
	private.Use(AuthRequired)
	{
		private.GET("/me", me)
		private.GET("/status", status)
	}

	r.Run() 
}

func UsersHandle(c *gin.Context) {
	var users []User
	db.Find(&users)
	c.JSON(http.StatusOK, users)
}