package response

import (
	"github.com/tamuramasaho/todo-app/entity"
	"time"
	// "fmt"
)

type(
	Todo struct {
		ID				int64     
		Title       	string    
		DueDate     	string
		DueDate2     	string
		RemindAt		string 
	}
	Todos []Todo
)

func NewTodo(entity *entity.Todo) *Todo {
	if entity == nil {
		return nil
	} else {
		return &Todo{
			ID:				entity.ID,
			Title:			entity.Title,
			DueDate:		entity.DueDate.Format("1/2"),
			// MEMO: edit.htmlのvalueのためのフォーマット（名前は再考の余地あり）
			DueDate2:		entity.DueDate.Format("2006-01-02"),
			RemindAt:		formatReminderTime(entity.ShouldRemind, entity.RemindAt),
		}
	}
}

func NewTodos(entities *entity.Todos) *Todos {
	ret := &Todos{}
	for _, v := range *entities {
		*ret = append(*ret, *NewTodo(&v))
	}
	return ret
}

// TODO: スマートにしよう
func formatReminderTime(b bool, t time.Time) string {
	var ret string
	if b && t.Format("15:04") == ("20:00") {
			ret = "ON: 前日" + t.Format("15時")
	}
	if b && (t.Format("15:04") == ("12:00") || t.Format("15:04") == ("09:00")){
		ret = "ON: 当日" + t.Format("15時")
	}
	if !b {
		ret = "OFF"
	}
	return ret
}