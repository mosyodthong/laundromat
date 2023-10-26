package controller

import (
	"laundromat/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type laundromatController struct {
	laundromatService service.LaundromatService
}

func NewLaundromatController(laundromatService service.LaundromatService) laundromatController {
	return laundromatController{laundromatService: laundromatService}
}

func (h laundromatController) GetAllMachineType(c *fiber.Ctx) error {

	machineType, err := h.laundromatService.GetAllMachineType()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"data": machineType,
	})
}

func (h laundromatController) CreateWashingMachine(c *fiber.Ctx) error {

	request := service.WashingMachineRequest{}

	err := c.BodyParser(&request)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}

	response, err := h.laundromatService.CreateWashingMachine(request)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(&fiber.Map{
		"data": response,
	})

}

func (h laundromatController) GetAlleWashingMachine(c *fiber.Ctx) error {

	limit, _ := strconv.Atoi(c.Query("limit"))
	page, _ := strconv.Atoi(c.Query("page"))

	params := service.ParamsFilter{
		Q:     c.Query("q"),
		Limit: limit,
		Page:  page,
	}

	response, err := h.laundromatService.GetAlleWashingMachine(params)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"data": response,
	})

}

func (h laundromatController) CreateWorkingMachine(c *fiber.Ctx) error {

	request := service.WorkingMachineRequest{}

	err := c.BodyParser(&request)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}

	response, err := h.laundromatService.CreateWorkingMachine(request)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(&fiber.Map{
		"data": response,
	})

}

func (h laundromatController) GetAllCheckWorkingMachine(c *fiber.Ctx) error {

	limit, _ := strconv.Atoi(c.Query("limit"))
	page, _ := strconv.Atoi(c.Query("page"))

	params := service.ParamsFilter{
		Q:     c.Query("q"),
		Limit: limit,
		Page:  page,
	}

	response, err := h.laundromatService.GetAllCheckWorkingMachine(params)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"data": response,
	})

}
