package main

import (
  "log"
  "github.com/gofiber/fiber/v2"
  "github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/a-h/templ"
  components "todoer/src"
  fns "todoer/fns"
)

func main() {
  app := fiber.New()

  db := fns.Conn()

  app.Static("/dist", "./dist")
	
  component := components.HomePage()
  app.Get("/", adaptor.HTTPHandler(templ.Handler(component)))
  
  component = components.Todos(fns.GetAllTodos(db))
  app.Get("/api/todos", adaptor.HTTPHandler(templ.Handler(component)))
  
  log.Fatal(app.Listen(":3000"))
}
