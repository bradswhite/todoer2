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
  
  component = components.AllTodos(fns.GetAllTodos(db))
  app.Get("/api/todos", adaptor.HTTPHandler(templ.Handler(component)))
  
  app.Get("/api/todos/:username", func (c *fiber.Ctx) error {
    component = components.Todos(fns.GetTodosForUser(db, c.Params("username")))
    return adaptor.HTTPHandler(templ.Handler(component))(c)
  })
  
  app.Post("/api/add-todo", func (c *fiber.Ctx) error {
    component = components.Toast(fns.AddTodo(db, c.FormValue("username"), c.FormValue("text")))
    return adaptor.HTTPHandler(templ.Handler(component))(c)
    //c.SendString("Successfully added todo")
  })

  /* Page Not Found Management */
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).SendFile("./dist/404.html")
	})
  
  log.Fatal(app.Listen(":3000"))
}
