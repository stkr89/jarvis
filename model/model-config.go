package model

type Tasks struct {
	Tasks []*Task `yaml:"tasks" validate:"gte=1,dive"`
}

type Task struct {
	Name     string    `yaml:"name" validate:"required"`
	TaskType *TaskType `yaml:"task_type" validate:"required"`
}

type TaskType struct {
	Http *TaskTypeHttp `yaml:"http" validate:"dive"`
}
