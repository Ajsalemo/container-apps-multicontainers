package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type UpstreamControllerResponse struct {
	Message string `json:"message"`
}

func UpstreamController(c *fiber.Ctx) error {
	c.Request().Header.VisitAll(func(key, value []byte) {
		fmt.Println("Request Header", string(key), "value", string(value))
	})

	r := UpstreamControllerResponse{
		Message: "Printing request headers to console..",
	}

	return c.JSON(r)
}