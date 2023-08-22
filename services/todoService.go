package services

import (
	"github.com/ykotanli/dto"
	"github.com/ykotanli/models"
	"github.com/ykotanli/repository"
)

type DefaultTodoService struct {
	Repo repository.TodoRepository
}

type TodoService interface {
	TodoInsert(todo models.Todo) (*dto.TodoDTO, error)
	TodoGetAll() ([]models.Todo, error)
	TodoDelete(id string) (*dto.TodoDTO, error)
	TodoDeleteAll() (*dto.TodoDTO, error)
}

func (t DefaultTodoService) TodoInsert(todo models.Todo) (*dto.TodoDTO, error) {
	var res dto.TodoDTO
	if len(todo.Title) <= 2 {
		res.Status = false
		return &res, nil
	}
	result, err := t.Repo.Insert(todo)

	if err != nil || result == false {
		res.Status = false
		return &res, err
	}

	res = dto.TodoDTO{Status: result}
	return &res, nil
}

func (t DefaultTodoService) TodoGetAll() ([]models.Todo, error) {
	result, err := t.Repo.Getall()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (t DefaultTodoService) TodoDelete(id string) (*dto.TodoDTO, error) {
	var res dto.TodoDTO
	result, err := t.Repo.Delete(id)
	if err != nil || result == false {
		res.Status = false
		return &res, err
	}
	res = dto.TodoDTO{Status: result}
	return &res, nil
}

func (t DefaultTodoService) TodoDeleteAll() (*dto.TodoDTO, error) {
	var res dto.TodoDTO
	result, err := t.Repo.DeleteAll()
	if err != nil || result == false {
		res.Status = false
		return &res, err
	}
	res = dto.TodoDTO{Status: result}
	return &res, nil
}

func NewTodoService(Repo repository.TodoRepository) DefaultTodoService {
	return DefaultTodoService{Repo: Repo}
}
