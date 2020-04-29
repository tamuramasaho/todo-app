package usecase

import (
  "github.com/tamuramasaho/todo-app/entity"
  "github.com/tamuramasaho/todo-app/form"
  "github.com/tamuramasaho/todo-app/repository"
  "time"
  "strconv"
  _"fmt"
)

type (
  ITodo interface {
	GetList() (*entity.Todos, error)
	GetListByWord(word string) (*entity.Todos, error)
	Create(form *form.Todo) (*entity.Todo, error)
	GetByID(id int64) (*entity.Todo, error)
	Update(id int64, form *form.Todo) (*entity.Todo, error)
	Delete(id int64) error
	SendReminder() error
	ChangeActiveness(id int64) (*entity.Todo, error)
  }
  Todo struct {
    TodoRepo repository.ITodo
  }
)

func NewTodo() ITodo {
  return &Todo{
	TodoRepo: repository.NewTodo(),
  }
}

func (usecase *Todo)GetList() (*entity.Todos, error) {
  entities, err := usecase.TodoRepo.All(&entity.Todos{})
  if err != nil {
	return nil, err
  }
  return entities, nil
}

func (usecase *Todo)GetListByWord(word string) (*entity.Todos, error) {
  entities, err := usecase.TodoRepo.ByWord(&entity.Todos{}, word)
  if err != nil {
	return nil, err
  }
  return entities, nil
}

func (usecase *Todo)Create(form *form.Todo) (*entity.Todo, error) {
  entity := &entity.Todo{
	Title: 			form.Title,
	DueDate:		form.DueDate,
	RemindAt:		form.DueDate.Add(-4 * time.Hour),
	ShouldRemind: 	true,
  }

  todo, err := usecase.TodoRepo.Create(entity)
  if err != nil {
	  return nil, err
  }
  return todo, nil
}

func (usecase *Todo)GetByID(id int64) (*entity.Todo, error) {
	entity, err := usecase.TodoRepo.ByID(&entity.Todo{}, id)
	if err != nil {
		return nil, err
	}
	return entity, nil
}

func (usecase *Todo)Update(id int64, form *form.Todo) (*entity.Todo, error) {
	entity, err := usecase.TodoRepo.ByID(&entity.Todo{}, id)
	if err != nil {
		return nil, err
	}
	entity.Title = form.Title
	entity.DueDate = form.DueDate
	entity.RemindAt = form.RemindAt
	entity.ShouldRemind = true

	todo, err := usecase.TodoRepo.Update(entity)
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func (usecase *Todo)Delete(id int64) error {
	entity, err := usecase.TodoRepo.ByID(&entity.Todo{}, id)
	if err != nil {
		return  err
	}

	err = usecase.TodoRepo.Delete(entity)
	return err
}


func (usecase *Todo)SendReminder() error {
	// 通知するtodoをフェッチしてあればSend to slack
	entities, err := usecase.TodoRepo.ByTime(&entity.Todos{}, time.Now())
	if err != nil {
		return  err
	}

	if len(*entities) > 0 {
		err = usecase.TodoRepo.SendToSlack(makeMessage(entities), entities)
	}
	return err
}

func (usecase *Todo)ChangeActiveness(id int64) (*entity.Todo, error) {
	entity, err := usecase.TodoRepo.ByID(&entity.Todo{}, id)
	if err != nil {
		return  nil, err
	}
	
	if entity.ShouldRemind {
		entity.ShouldRemind = false
	} else {
		entity.ShouldRemind = true
	}
	todo, err := usecase.TodoRepo.Update(entity)
	if err != nil {
		return nil, err
	}
	return todo, nil
}


func makeMessage(e *entity.Todos) string {
	var msg string
	if (*e)[0].RemindAt.Format("15:04") == ("20:00") {
		msg = "明日の予定が" + strconv.Itoa(len(*e)) + "件あります"
	} else {
		msg = "今日の予定が" + strconv.Itoa(len(*e)) + "件あります"
	}
	return msg
}