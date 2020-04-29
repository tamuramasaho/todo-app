package handler

import (
	"net/http/httptest"
	"testing"
	"strings"
	"html/template"
	// "fmt"
	"net/http"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/golang/mock/gomock"
	"github.com/tamuramasaho/todo-app/handler/mock_todos"
	"github.com/tamuramasaho/todo-app/entity"
	"io"
	"errors"
	"time"
	"gopkg.in/go-playground/validator.v9"
	"os"
)


func TestMain(m *testing.M) {
	setUp()
	code := m.Run()
	os.Exit(code)
}


func TestList(t *testing.T) {
	TodoUsecase.EXPECT().SendReminder()
	TodoUsecase.EXPECT().GetList().Return(&todos, nil)
	
	req := httptest.NewRequest(http.MethodGet, "/todos", nil)
	rec := httptest.NewRecorder()	
	c := e.NewContext(req, rec)
	
	err := h.List(c)
	body := rec.Body.String()	
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Containsf(t, body, "Work", "Body does not containe Work")
		assert.Containsf(t, body, "Music", "Body does not containe Music")
	}
}

func TestGet(t *testing.T) {
	TodoUsecase.EXPECT().GetByID(gomock.Any()).Return(&todos[0], nil)
	
	req := httptest.NewRequest(http.MethodGet, "/todos/1", nil)
	rec := httptest.NewRecorder()	
	c := e.NewContext(req, rec)
	c.SetPath("/todos/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")
	
	err := h.Get(c)
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestSearch(t *testing.T) {
	TodoUsecase.EXPECT().GetListByWord(gomock.Any()).Return(&todos, nil)
	
	req := httptest.NewRequest(http.MethodGet, "/search?title=Work", nil)
	rec := httptest.NewRecorder()	
	c := e.NewContext(req, rec)
	
	err := h.Search(c)
	body := rec.Body.String()
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Containsf(t, body, "Work", "Body does not containe Work")
	}
}

func TestCreate(t *testing.T) {
	TodoUsecase.EXPECT().Create(gomock.Any()).Return(&todos[0], nil)
	
	req := httptest.NewRequest(http.MethodPost, "/todos", strings.NewReader(todoJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()	
	c := e.NewContext(req, rec)
	
	err := h.Create(c)
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusCreated, rec.Code)
	}
}

func TestEdit(t *testing.T) {
	TodoUsecase.EXPECT().GetByID(gomock.Any()).Return(&todos[0], nil)
	
	req := httptest.NewRequest(http.MethodGet, "/todos/1/edit", nil)
	rec := httptest.NewRecorder()	
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")
	
	err := h.Edit(c)
	body := rec.Body.String()
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Containsf(t, body, "Work", "Body does not containe Work")
	}
}

func TestUpdate(t *testing.T) {
	TodoUsecase.EXPECT().Update(gomock.Any(), gomock.Any())
	
	req := httptest.NewRequest(http.MethodPut, "/tasks/1", strings.NewReader(todoJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()	
	c := e.NewContext(req, rec)
	c.SetPath("/todos/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	err := h.Update(c)	
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestDestroy(t *testing.T) {
	TodoUsecase.EXPECT().Delete(gomock.Any())
	
	req := httptest.NewRequest(http.MethodDelete, "/todos/1", nil)
	rec := httptest.NewRecorder()	
	c := e.NewContext(req, rec)
	c.SetPath("/todos/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	err := h.Destroy(c)
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

// 以下はテストをするための準備
var t testing.T
var ctrl *gomock.Controller
var TodoUsecase *mock_todos.MockITodo
var h *Todo
var e *echo.Echo

func setUp() {
	ctrl = gomock.NewController(&t) 
	TodoUsecase = mock_todos.NewMockITodo(ctrl) 
	h = &Todo { TodoUsecase: TodoUsecase } 

	e = echo.New() 
	e.Renderer = &TemplateRegistry{ templates: templates }
	e.Validator = NewValidator()
}

func NewValidator() *Validator {
	return &Validator{
		validator: validator.New(),
	}
}
type Validator struct {
	validator *validator.Validate
}
func (v *Validator) Validate(i interface{}) error {
	return v.validator.Struct(i)
}

type TemplateRegistry struct {
    templates map[string]*template.Template
}
var templates = map[string]*template.Template{
	"index.html": template.Must(template.ParseFiles("../public/views/index.html", 
								"../public/views/layout.html")),
	"edit.html": template.Must(template.ParseFiles("../public/views/edit.html",
								 "../public/views/layout.html")),
}
func (t *TemplateRegistry) Render(w io.Writer, name string, data interface{},
									 					c echo.Context) error {
	tmpl, ok := t.templates[name]
	if !ok {
	  err := errors.New("Template not found -> " + name)
	  return err
	}
	return tmpl.ExecuteTemplate(w, "layout.html", data)
}

var dueDate, _ = time.Parse("2006-01-02", "1989-11-05")
var remindAt, _ = time.Parse("2006-01-02 15:04", "1989-11-04 20:00")

var todos = entity.Todos{
	entity.Todo{
		ID: 1,
		Title: "Work",
		DueDate: dueDate,
		RemindAt: remindAt,
		ShouldRemind: true,
	},
	entity.Todo{
		ID: 2,
		Title: "Music",
		DueDate: dueDate,
		RemindAt: remindAt,
		ShouldRemind: true,
	},
}
var todoJSON = 
`{
	"ID": "1", 
	"title": "Work",
	"due_date": "1989-11-05",
	"remind_at": "the_day_morning",
	"should_remind": "true"
}`

