package repository

import (
	"github.com/tamuramasaho/todo-app/db"
	"github.com/tamuramasaho/todo-app/entity"
	"time"
	"strconv"
  	"github.com/ashwanthkumar/slack-go-webhook"
	"os"
	// "fmt"
)

type (
	ITodo interface {
		All(e *entity.Todos) (*entity.Todos, error)
		ByWord(e *entity.Todos, word string) (*entity.Todos, error)
		Create(e *entity.Todo) (*entity.Todo, error)
		ByID(e *entity.Todo, id int64) (*entity.Todo, error)
		Update(e *entity.Todo) (*entity.Todo, error)
		Delete(e *entity.Todo) error
		ByTime(e *entity.Todos, now time.Time) (*entity.Todos, error)
		SendToSlack(msg string, entities *entity.Todos) error
	}
	Todo struct {}
)

func NewTodo() ITodo {
	return &Todo{}
}

func (repo *Todo)All(e *entity.Todos) (*entity.Todos, error) {
	err := db.New().Order("remind_at").Find(e).Error
	if err != nil {
		return nil, err
	}
	return e, err
}

func (repo *Todo)ByWord(e *entity.Todos, word string) (*entity.Todos, error) {
	err := db.New().Where("title LIKE ?", "%"+word+"%").Find(e).Error
	if err != nil {
		return nil, err
	}
	return e, err
}

func (repo *Todo)Create(e *entity.Todo) (*entity.Todo, error) {
	err := db.New().Create(e).Error
	if err != nil {
		return nil, err
	}
	return e, nil
}

func (repo *Todo)ByID(e *entity.Todo, id int64) (*entity.Todo, error) {
	err := db.New().Find(e, id).Error
	if err != nil {
		return nil, err
	}
	return e, nil
}

func (repo *Todo)Update(e *entity.Todo) (*entity.Todo, error) {
	err := db.New().Save(e).Error
	if err != nil {
		return nil, err
	}
	return e, nil
}

func (repo *Todo)Delete(e *entity.Todo) error {
	err := db.New().Delete(e).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *Todo)ByTime(e *entity.Todos, t time.Time) (*entity.Todos, error) {
	layout := "2006-1-2 15:04"
	err := db.New().Where("remind_at = ? AND should_remind = ?", 
								t.Format(layout), true).Find(e).Error
	if err != nil {
		return nil, err
	}
	return e, err
}

const (
    WEBHOOKURL = "yourWebhookUrl"
    CHANNEL    = "yourchannell"
    USERNAME   = "ToDoAppBot"
)

func (repo *Todo)SendToSlack(msg string, entities *entity.Todos) error {
	field1 := slack.Field{Title: "メッセージ", Value: msg}
    attachment := slack.Attachment{}
	attachment.AddField(field1)
    for i:=0; i < len(*entities); i++ {
        title := strconv.Itoa(i + 1) + "件目"
        attachment.AddField(slack.Field{Title: title, Value: (*entities)[i].Title})
    }
	color := "good"
    attachment.Color = &color
    payload := slack.Payload{
        Username:    USERNAME,
        Channel:     CHANNEL,
        Attachments: []slack.Attachment{attachment},
    }
    err := slack.Send(WEBHOOKURL, "", payload)
    if err != nil {
        os.Exit(1)
	}
	return nil
}