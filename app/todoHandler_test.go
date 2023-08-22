package app

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/ykotanli/mocks/services"
	"github.com/ykotanli/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var td TodoHandler
var mockService *services.MockTodoService

func setup(t *testing.T) func() {
	ctrl := gomock.NewController(t)
	mockService = services.NewMockTodoService(ctrl)

	td = TodoHandler{
		Service: mockService,
	}

	return func() {
		defer ctrl.Finish()
	}
}

func TestTodoHandler_GetAllTodo(t *testing.T) {
	teardown := setup(t)
	defer teardown()
	router := fiber.New()

	router.Get("/api/todo", td.GetAllTodo)

	var FakeDataForHandler = []models.Todo{
		{Id: primitive.NewObjectID(), Title: "Test1", Content: "content 1"},
		{Id: primitive.NewObjectID(), Title: "Test2", Content: "content 2"},
		{Id: primitive.NewObjectID(), Title: "Test3", Content: "content 3"},
	}

	mockService.EXPECT().TodoGetAll().Return(FakeDataForHandler, nil)

	req := httptest.NewRequest("GET", "/api/todo", nil)

	resp, _ := router.Test(req, 1)

	assert.Equal(t, 200, resp.StatusCode)
}
