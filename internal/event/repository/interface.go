package repository

import (
	"context"
	"github.com/fwidjaya20/demo-distributed-event-store/internal/event/models"
)

type EventRepositoryInterface interface {
	Insert(context context.Context, payload *models.EventStore) (*string, error)
	FindById(context context.Context, id string) (*models.EventStore, error)
}
