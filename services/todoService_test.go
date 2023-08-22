package services

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/ykotanli/mocks/repository"
	"github.com/ykotanli/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var mockRepo *repository.MockTodoRepository
var service TodoService

var FakeData = []models.Todo{
	{primitive.NewObjectID(), "Test1", "content 1"},
	{primitive.NewObjectID(), "Test2", "content 2"},
	{primitive.NewObjectID(), "Test3", "content 3"},
}

func setup(t *testing.T) func() {
	ct := gomock.NewController(t)
	defer ct.Finish()

	mockRepo = repository.NewMockTodoRepository(ct)
	service = NewTodoService(mockRepo)
	return func() {
		service = nil
		defer ct.Finish()
	}
}

func TestDefaultTodoService_GetAll(t *testing.T) {
	teardown := setup(t)
	defer teardown()

	mockRepo.EXPECT().Getall().Return(FakeData, nil)

	res, err := service.TodoGetAll()

	if err != nil {
		t.Error(err)
	}
	assert.NotEmpty(t, res)

}
