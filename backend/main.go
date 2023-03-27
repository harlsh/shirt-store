package main

import (
	"log"
	"os"
	"shirt-store/handlers"
	"shirt-store/middleware"
	"shirt-store/models"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var db *gorm.DB

const secret = "shirtsarecool"
const userkey = "user"

func main() {

	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file")
	}
	port := os.Getenv("PORT")

	r := SetupRouter()

	log.Fatal(r.Run("localhost:" + port))
}

func DbInit() *gorm.DB {
	db, err := models.Setup()
	if err != nil {
		log.Println("Problem setting up database")
	}
	return db
}

func SetupRouter() *gin.Engine {
	r := gin.Default()

	db := DbInit()

	server := handlers.NewServer(db)

	router := r.Group("/api")

	router.POST("/register", server.Register)
	router.POST("/login", server.Login)
	router.GET("/me", server.Me)
	authorized := r.Group("/api/admin")

	authorized.Use(middleware.JwtAuthMiddleware())
	// authorized.GET("/groceries", server.GetGroceries)
	// authorized.POST("/grocery", server.PostGrocery)
	return r

}
