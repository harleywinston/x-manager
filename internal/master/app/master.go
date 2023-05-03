package master

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/harleywinston/x-manager/internal/master/database"
	"github.com/harleywinston/x-manager/internal/master/transport"
)

func initDB() error {
	postgres_host := os.Getenv("POSTGRES_HOST")
	postgres_user := os.Getenv("POSTGRES_USER")
	postgres_password := os.Getenv("POSTGRES_PASSWORD")
	postgres_db := os.Getenv("POSTGRES_DB")
	postgres_port := os.Getenv("POSTGRES_PORT")
	dsn := fmt.Sprintf(
		`host=%v user=%v password=%v dbname=%v port=%v sslmode=disable`,
		postgres_host,
		postgres_user,
		postgres_password,
		postgres_db,
		postgres_port,
	)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	if err != nil {
		return err
	}
	database.DB = db
	if err := database.InitModels(); err != nil {
		return err
	}

	return nil
}

func registerHandlers() error {
	r := gin.Default()

	usersHandlers := transport.UsersHandler{}
	r.GET("/user", usersHandlers.GetUserHandler)
	r.POST("/user", usersHandlers.AddUserHandler)
	r.DELETE("/user", usersHandlers.DeleteUserHandler)

	resourcesHandlers := transport.ResourcesHandler{}
	r.GET("/resource", resourcesHandlers.GetResourcesHandler)
	r.POST("/resource", resourcesHandlers.AddResourcesHandler)
	r.DELETE("/resource", resourcesHandlers.DeleteResourcesHandler)

	return r.Run()
}

func InitApp() error {
	if err := initDB(); err != nil {
		return err
	}

	if err := registerHandlers(); err != nil {
		return err
	}

	return nil
}
