package form

import (
	"testing"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"strings"
	"net/http"
	"net/http/httptest"
	"net/url"
	// "fmt"
	"gopkg.in/go-playground/validator.v9"
	"time"
)

func TestNewTodo(t *testing.T) {
	e := echo.New()
	e.Validator = NewValidator()
	req := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(todoJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
 	c := e.NewContext(req, rec)
	actual, err := NewTodo(c) 
	
	dueDate := time.Date(1989, 11, 5, 0, 0, 0, 0, time.UTC)
	remindAt := time.Date(1989, 11, 5, 9, 0, 0, 0, time.UTC)
	expected := &Todo {
		Title: "Work",
		DueDate: dueDate,
		RemindAt: remindAt,
	}
	
  	if assert.NoError(t, err) {
		assert.Equal(t, expected, actual)
	}
}

func TestGetID(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	actual, err := GetID(c)
	expected := int64(1)

	if assert.NoError(t, err) {
		assert.Equal(t, expected, actual)
	}
}

func TestGetWord(t *testing.T) {
	e := echo.New()
	q := make(url.Values)
	q.Set("word", "work")
	req := httptest.NewRequest(http.MethodGet, "/search?"+q.Encode(), nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	actual := GetWord(c)
	expected := "work"

	assert.Equal(t, expected, actual)
}

var todoJSON = 
`{
	"title": "Work",
	"due_date": "1989-11-05",
	"remind_at": "the_day_morning"
}`


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