package handler

import (
	"github.com/labstack/echo"
	"github.com/Tech-Design-Inc/sirius/form"
	"github.com/Tech-Design-Inc/sirius/response"
	"github.com/Tech-Design-Inc/sirius/usecase"
	"net/http"
	// "fmt"
	// "time"
	// "github.com/robfig/cron/v3"
	// "github.com/carlescere/scheduler"
)

type (
	ITodo interface {
		List(c echo.Context) error
		Get(c echo.Context) error
		Search(c echo.Context) error
		Create(c echo.Context) error
		Edit(c echo.Context) error
		Update(c echo.Context) error
		Destroy(c echo.Context) error
		ChangeActiveness(c echo.Context) error
	}

	Todo struct {
		TodoUsecase usecase.ITodo
	}
)

func NewTodo() *Todo {
	return &Todo{ TodoUsecase: usecase.NewTodo() }
}

func (h *Todo) List(c echo.Context) error {
	todos, err := h.TodoUsecase.GetList()
	if err != nil {
		return err
	}
	return c.Render(http.StatusOK, "index.html", response.NewTodos(todos))
}

func (h *Todo) Get(c echo.Context) error {
	id, err := form.GetID(c)
	if err != nil {
		return err
	}

	todo, err := h.TodoUsecase.GetByID(id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, response.NewTodo(todo))
}

func (h *Todo) Search(c echo.Context) error {
	word := form.GetWord(c)

	todos, err := h.TodoUsecase.GetListByWord(word)
	if err != nil {
		return err
	}
	return c.Render(http.StatusOK, "index.html", response.NewTodos(todos))
}

func (h *Todo) Create(c echo.Context) error {
	form, err := form.NewTodo(c)
	if err != nil {
		return err
	}
	todo, err := h.TodoUsecase.Create(form)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, response.NewTodo(todo))
}

func (h *Todo) Edit(c echo.Context) error {
	id, err := form.GetID(c)
	if err != nil {
		return err
	}

	todo, err := h.TodoUsecase.GetByID(id)
	if err != nil {
		return err
	}
	return c.Render(http.StatusOK, "edit.html", response.NewTodo(todo))
}

func (h *Todo) Update(c echo.Context) error {
	id, err := form.GetID(c)
	if err != nil {
		return err
	}

	form, err := form.NewTodo(c)
	if err != nil {
		return err
	}

	_, err = h.TodoUsecase.Update(id, form)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, "Updated the ToDo")
}

func (h *Todo) Destroy(c echo.Context) error {
	id, err := form.GetID(c)
	if err != nil {
		return err
	}

	err = h.TodoUsecase.Delete(id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, "Destroyed the ToDo")
}

func (h *Todo) ChangeActiveness(c echo.Context) error {
	id, err := form.GetID(c)
	if err != nil {
		return err
	}
	todo, err := h.TodoUsecase.ChangeActiveness(id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, response.NewTodo(todo))
}
