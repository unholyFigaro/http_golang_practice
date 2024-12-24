package taskService

import (
	"gorm.io/gorm"
)

type TaskRepository interface {
	CreateTask(task Task) (Task, error)
	GetAllTasks() ([]Task, error)
	UpdateTaskByID(id uint, task Task) (Task, error)
	DeleteTaskByID(id uint) (Task, error)
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *taskRepository {
	return &taskRepository{db: db}
}

func (r *taskRepository) GetAllTasks() ([]Task, error) {
	var tasks []Task
	err := r.db.Find(&tasks)
	if err.Error != nil {
		return []Task{}, err.Error
	}
	return tasks, nil
}

func (r *taskRepository) CreateTask(task Task) (Task, error) {
	err := r.db.Create(&task)
	if err.Error != nil {
		return Task{}, err.Error
	}
	return task, nil
}

func (r *taskRepository) UpdateTaskByID(id uint, task Task) (Task, error) {
	var existingTask Task
	err := r.db.Model(&task).Where("id = ?", id).Updates(task)
	if err.Error != nil {
		return Task{}, err.Error
	}
	r.db.Find(&existingTask, id)
	return existingTask, nil
}

func (r *taskRepository) DeleteTaskByID(id uint) (Task, error) {
	var task Task
	if err := r.db.First(&task, id).Error; err != nil {
		return Task{}, err
	}
	if err := r.db.Model(&task).Delete(&task, id).Error; err != nil {
		return Task{}, err
	}
	return task, nil
}
