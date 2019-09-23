package grpc

import (
	"context"
	"github.com/fwidjaya20/demo-distributed-event-store/internal/event"
	"github.com/fwidjaya20/demo-distributed-event-store/internal/event/models"
	pb "github.com/fwidjaya20/demo-distributed-event-store/pkg/protobuf/eventstore"
	"log"
)

type EventGrpcService struct {
	EventService event.ServiceInterface
}

func (s EventGrpcService) CreateEvent(context context.Context, payload *pb.Event) (*pb.EventResponse, error) {
	log.Println("[EVENT_GRPC_SERVICE] CreateEvent", payload.String())

	m := &models.EventStore{
		Id:            payload.EventId,
		Type:          payload.EventType,
		AggregateId:   payload.AggregateId,
		AggregateType: payload.AggregateType,
		EventPayload:  payload.EventData,
		Channel:       payload.Channel,
	}

	result, err := s.EventService.Create(context, m)

	if nil != err {
		return nil, err
	}

	return &pb.EventResponse{
		Event: &pb.Event{
			EventId:       result.Id,
			EventType:     result.Type,
			AggregateId:   result.AggregateId,
			AggregateType: result.AggregateType,
			EventData:     result.EventPayload,
			Channel:       result.Channel,
		},
	}, nil
}
