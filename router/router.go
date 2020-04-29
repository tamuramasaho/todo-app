package router

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"github.com/tamuramasaho/todo-app/handler"
)

func New() *echo.Echo {
	e := echo.New()
	e.Validator = NewValidator()
	e.Renderer = NewTemplates()

	e.Logger.SetLevel(log.DEBUG)
	e.Pre(middleware.RemoveTrailingSlash())
	e.Pre(middleware.MethodOverrideWithConfig(middleware.MethodOverrideConfig{
		Getter: middleware.MethodFromForm("_method"),
	}))
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	e.GET("/healthcheck", func(c echo.Context) error {
		return c.String(200, "OK")
	})

	e.Static("/static", "public")  

	todo := handler.NewTodo()
	e.GET("/todos", todo.List)
	e.GET("/todos/:id", todo.Get)
	e.GET("/search", todo.Search)
	e.POST("/todos", todo.Create)                                                                                                                                                                                                                                                                                                            
	e.GET("/todos/:id/edit", todo.Edit)
	e.PUT("/todos/:id", todo.Update)
	e.DELETE("todos/:id",todo.Destroy)
	e.PATCH("todos/:id", todo.ChangeActiveness)

	return e
}
