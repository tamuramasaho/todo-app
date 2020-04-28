package usecase

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/Tech-Design-Inc/sirius/usecase/mock_todos"
	"github.com/Tech-Design-Inc/sirius/entity"
	"github.com/golang/mock/gomock"
	"github.com/Tech-Design-Inc/sirius/form"
	"time"
)

func TestGetList(t *testing.T) {
	TodoRepo.EXPECT().All(gomock.Any()).Return(&todos, nil)
	
	entities, err := h.GetList()
	if assert.NoError(t, err) {
		assert.Equal(t, entities, &todos)
	}
}

func TestGetListByWord(t *testing.T) {
	TodoRepo.EXPECT().ByWord(gomock.Any(), "work").Return(&todos, nil)
	
	entities, err := h.GetListByWord("work")
	if assert.NoError(t, err) {
		assert.Equal(t, entities, &todos)
	}
}

func TestCreate(t *testing.T) {
	TodoRepo.EXPECT().Create(gomock.Any()).Return(&todos[0], nil)
	
	entity, err := h.Create(&formTodo)
	if assert.NoError(t, err) {
		assert.Equal(t, entity, &todos[0])
	}
}

func TestGetByID(t *testing.T) {
	TodoRepo.EXPECT().ByID(gomock.Any(), gomock.Any()).Return(&todos[0], nil)
	
	entity, err := h.GetByID(1)
	if assert.NoError(t, err) {
		assert.Equal(t, entity, &todos[0])
	}
}

func TestUpdate(t *testing.T) {
	TodoRepo.EXPECT().ByID(gomock.Any(), gomock.Any()).Return(&todos[0], nil)
	TodoRepo.EXPECT().Update(gomock.Any()).Return(&todos[0], nil)
	
	entity, err := h.Update(1, &formTodo)
	if assert.NoError(t, err) {
		assert.Equal(t, entity, &todos[0])
	}
}

func TestDelete(t *testing.T) {
	TodoRepo.EXPECT().ByID(gomock.Any(), gomock.Any()).Return(&todos[0], nil)
	TodoRepo.EXPECT().Delete(gomock.Any())

	err := h.Delete(1)
	assert.NoError(t, err)
}

func TestCheck(t *testing.T) {
	TodoRepo.EXPECT().ByTime(gomock.Any(), gomock.Any()).Return(&todos, nil)
	TodoRepo.EXPECT().SendToSlack(gomock.Any(), gomock.Any())
	
	err := h.SendReminder()
	assert.NoError(t, err)
}



var todos = entity.Todos{
	entity.Todo{
		ID: 1,
		Title: "Work",
		DueDate: time.Now(),
		RemindAt: time.Now(),
		ShouldRemind: true,
	},
	entity.Todo{
		ID: 2,
		Title: "Music",
		DueDate: time.Now(),
		RemindAt: time.Now(),
		ShouldRemind: true,
	},
}

var formTodo = form.Todo{
	Title: "Work",
	DueDate: time.Now(),
	RemindAt: time.Now(),
}

var t testing.T
var ctrl = gomock.NewController(&t)
var TodoRepo = mock_todos.NewMockITodo(ctrl)
var h = &Todo { TodoRepo: TodoRepo }