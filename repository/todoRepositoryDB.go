package repository

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/ykotanli/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TodoRepositoryDB struct {
	TodoCollection *mongo.Collection
}

type TodoRepository interface {
	Insert(todo models.Todo) (bool, error)
	Getall() ([]models.Todo, error)
	Delete(id string) (bool, error)
	DeleteAll() (bool, error)
}

func (t TodoRepositoryDB) Insert(todo models.Todo) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	todo.Id = primitive.NewObjectID()
	result, err := t.TodoCollection.InsertOne(ctx, todo)
	if result.InsertedID == nil || err != nil {
		errors.New("Error inserting todo")
		return false, err
	}
	return true, nil
}

func (t TodoRepositoryDB) Getall() ([]models.Todo, error) {
	var todo models.Todo
	var todos []models.Todo

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := t.TodoCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	for result.Next(ctx) {
		if err := result.Decode(&todo); err != nil {
			log.Fatalln(err)
			return nil, err
		}
		todos = append(todos, todo)
	}
	return todos, nil
}

func (t TodoRepositoryDB) Delete(id string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return false, err
	}

	_, err = t.TodoCollection.DeleteOne(ctx, bson.M{"_id": objectId})
	if err != nil {
		return false, err
	}
	return true, nil
}

func (t TodoRepositoryDB) DeleteAll() (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := t.TodoCollection.DeleteMany(ctx, bson.M{})
	if err != nil {
		return false, err
	}
	return true, nil
}
func NewTodoRepositoryDB(dbClient *mongo.Collection) TodoRepositoryDB {
	return TodoRepositoryDB{TodoCollection: dbClient}
}
