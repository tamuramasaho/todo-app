package entity

import "time"

type (
	Todo struct {
		ID          	int64     `gorm:"column:id;primary_key"`
		Title       	string    `gorm:"title"`
		DueDate       	time.Time `gorm:"due_date"`
		RemindAt		time.Time `gorm:"remind_at"`
		ShouldRemind	bool	  `gorm:"should_remind"`
		CreatedAt   	time.Time `gorm:"column:created_at"`
		UpdatedAt   	time.Time `gorm:"column:updated_at`
	}
	Todos []Todo
)