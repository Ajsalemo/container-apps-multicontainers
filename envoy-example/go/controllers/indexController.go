package controllers

import (
    "github.com/gofiber/fiber/v2"
)

type IndexResponse struct {
	Message string `json:"message"`
}

func IndexController(c *fiber.Ctx) error {
	res := IndexResponse{
		Message: "container-apps-multicontainers-envoy-example-go",
	}

	return c.JSON(res)
}