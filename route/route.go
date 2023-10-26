package route

import (
	"laundromat/config"
	"laundromat/controller"
	"laundromat/repository"
	"laundromat/service"

	"github.com/gofiber/fiber/v2"
)

func Router(app *fiber.App) {
	db := config.InitDatabase()

	api := app.Group("/api")

	laundromatRepository := repository.NewLaundromatRepositoryPG(db)
	laundromatService := service.NewOrderService(laundromatRepository)
	laundromatController := controller.NewLaundromatController(laundromatService)

	api.Get("/laundromat/types", laundromatController.GetAllMachineType)
	api.Get("/laundromat", laundromatController.GetAlleWashingMachine)
	api.Get("/laundromat/working", laundromatController.GetAllCheckWorkingMachine)
	api.Post("/laundromat", laundromatController.CreateWashingMachine)
	api.Post("/laundromat/working", laundromatController.CreateWorkingMachine)

}
