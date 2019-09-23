package event

import (
	"context"
	"github.com/fwidjaya20/demo-distributed-event-store/internal/event/models"
	"github.com/fwidjaya20/demo-distributed-event-store/internal/event/repository"
	pb "github.com/fwidjaya20/demo-distributed-event-store/pkg/protobuf/eventstore"
	"github.com/jmoiron/sqlx"
	"github.com/nats-io/go-nats"
	"log"
)

type ServiceInterface interface {
	Create(context context.Context, payload *models.EventStore) (*models.EventStore, error)
}

type eventService struct {
	repository repository.EventRepositoryInterface
	nats *nats.Conn
}

func NewEventService(db *sqlx.DB, natsConn *nats.Conn) ServiceInterface {
	return &eventService{
		repository: repository.NewEventRepository(db),
		nats: natsConn,
	}
}

func (es *eventService) Create(context context.Context, payload *models.EventStore) (*models.EventStore, error) {
	uuid, err := es.repository.Insert(context, payload)

	if nil != err {
		return nil, err
	}

	result, err := es.repository.FindById(context, *uuid)

	if nil != err {
		return nil, err
	}

	es.publishEvent(pb.Event{
		EventId:              payload.Id,
		EventType:            payload.Type,
		AggregateId:          payload.AggregateId,
		AggregateType:        payload.AggregateType,
		EventData:            payload.EventPayload,
		Channel:              payload.Channel,
	})

	return result, nil
}

func (es *eventService) publishEvent(event pb.Event) {
	eventMsg := []byte(event.EventData)

	_ = es.nats.Publish(event.Channel, eventMsg)

	log.Println("Event Published on Channel", event.Channel)
}