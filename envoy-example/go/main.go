package main

import (
    "log"

    "github.com/gofiber/fiber/v2"

	controllers "container-apps-multicontainers-envoy-example-go/controllers"
)

func main() {
    app := fiber.New()

    app.Get("/", controllers.IndexController)
    app.Get("/get/headers", controllers.UpstreamController)

    log.Fatal(app.Listen(":3000"))
}