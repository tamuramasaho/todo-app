package form

import (
  "github.com/labstack/echo"
  "strconv"
  "time"
  "fmt"
)

type(
	Todo struct {
		Title       	string    	`validate:"required"`
		DueDate     	time.Time    
		RemindAt		time.Time 
	}
	todoForBind struct {
		Title       	string    `json:"title"`
		DueDate     	string    `json:"due_date"`
		RemindAt		string    `json:"remind_at"`
	}
)

func InitTodo() *Todo {
	return &Todo{}
}

func InitTodoForBind() *todoForBind{
	return &todoForBind{}
}

func NewTodo(c echo.Context) (*Todo, error) {
	todo := InitTodoForBind()
	if err := c.Bind(todo); err != nil {
		return nil, err
	}
	var td time.Time
	if todo.DueDate != ""{
		td, _ = time.Parse("2006-01-02", todo.DueDate)
	} else {
		td, _ = time.Parse("2006-01-02", time.Now().Format("2006-01-02"))
	}
	tr := distinguishRemindTime(todo.RemindAt, td)
	form := InitTodo()
	form.Title = todo.Title
	form.DueDate = td
	form.RemindAt = tr
	
	if err := c.Validate(form); err != nil {
		return nil, err
	}
	return form, nil
}

func GetID(c echo.Context) (id int64, err error) {
	id, err = strconv.ParseInt(c.Param("id"), 10, 64)
	return id, err
}

func GetWord(c echo.Context) (word string) {
	word = c.FormValue("word")
	return word
}

// MEMO:　これでいいのか
func distinguishRemindTime(remindAt string, td time.Time) time.Time {
	var tr time.Time
	switch remindAt {
		case "the_day_before":
			tr = td.Add(-4 * time.Hour)
		case "the_day_morning":
			tr = td.Add(9 * time.Hour)
		case "the_day_noon":
			tr = td.Add(12 * time.Hour)
	}
	return tr
} 

var _ = fmt.Print