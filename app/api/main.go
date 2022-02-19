package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/viniciusarambul/go-flight/app/api/handler"
	"github.com/viniciusarambul/go-flight/app/api/presenter"
	"github.com/viniciusarambul/go-flight/app/infrastructure/repository"
	"github.com/viniciusarambul/go-flight/app/usecase/pilot"
	"github.com/viniciusarambul/go-flight/app/usecase/plane"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	engine := gin.Default()

	engine.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "estamos on"})
	})
	dsn := "root:root@tcp(docker.for.mac.localhost:3306)/goflight?&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("errou")
	}

	planeRepository := repository.NewPlaneMySQLRepository(db)
	planePresenter := presenter.NewPlanePresenter()
	planeUseCase := plane.NewPlaneUseCase(planeRepository, planePresenter)

	handler.NewPlaneHandler(engine, planeUseCase)

	pilotRepository := repository.NewPilotMySQLRepository(db)
	pilotPresenter := presenter.NewPilotPresenter()
	pilotUseCase := pilot.NewPilotUseCase(pilotRepository, pilotPresenter)

	handler.NewPilotHandler(engine, pilotUseCase)

	engine.Run()
}
