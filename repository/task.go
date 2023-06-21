package repository

import (
	"a21hc3NpZ25tZW50/model"

	"gorm.io/gorm"
)

type TaskRepository interface {
	Store(task *model.Task) error
	Update(id int, task *model.Task) error
	Delete(id int) error
	GetByID(id int) (*model.Task, error)
	GetList() ([]model.Task, error)
	GetTaskCategory(id int) ([]model.TaskCategory, error)
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepo(db *gorm.DB) *taskRepository {
	return &taskRepository{db}
}

func (t *taskRepository) Store(task *model.Task) error {
	err := t.db.Create(task).Error
	if err != nil {
		return err
	}

	return nil
}

func (t *taskRepository) Update(id int, task *model.Task) error {
	// Query Update Data
	salah := t.db.Where("id = ?",
		task.ID).Updates(task).Error
	if salah != nil {
		return salah
	}

	return nil // TODO: replace this
}

func (t *taskRepository) Delete(id int) error {
	// Query Delete Data
	salah := t.db.Delete(&model.Task{},
		id).Error
	if salah != nil {
		return salah // return Error Message
	}

	return nil // TODO: replace this
}

func (t *taskRepository) GetByID(id int) (*model.Task, error) {
	var task model.Task
	err := t.db.First(&task, id).Error
	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (t *taskRepository) GetList() ([]model.Task, error) {
	// Buat variabel baru students
	var tasks []model.Task

	// Query data students
	hasil := t.db.Find(&tasks)
	if hasil.Error != nil {
		return nil,
			hasil.Error
	}

	return tasks, nil // TODO: replace this
}

func (t *taskRepository) GetTaskCategory(id int) ([]model.TaskCategory, error) {
	// Buat variabel baru studentClasses
	var tasksCategory []model.TaskCategory

	// Query Join
	hasil := t.db.Table("tasks").
		Select("tasks.id as id, tasks.title as title, categories.name as category").
		Joins("LEFT JOIN categories ON tasks.category_id = categories.id").
		Where("tasks.id = ?",
			id).
		Scan(&tasksCategory)

	if hasil.Error != nil {
		return nil,
			hasil.Error
	}

	return tasksCategory,
		nil // TODO: replace this
}
