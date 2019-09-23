package repository

import (
	"context"
	"github.com/fwidjaya20/demo-distributed-event-store/internal/event/models"
	"github.com/jmoiron/sqlx"
	"github.com/payfazz/go-apt/pkg/fazzcommon/formatter"
	"github.com/payfazz/go-apt/pkg/fazzdb"
)

type repository struct {
	Q 		*fazzdb.Query
	Event 	*models.EventStore
}

func NewEventRepository(db *sqlx.DB) EventRepositoryInterface {
	q := fazzdb.QueryDb(db, fazzdb.DEFAULT_QUERY_CONFIG)

	return &repository{
		Q: q,
		Event: models.EventStoreModel(),
	}
}

func (r *repository) Insert(context context.Context, payload *models.EventStore) (*string, error) {
	result, err := r.Q.Use(payload).InsertCtx(context, false)

	if nil != err {
		return nil, err
	}

	uuid := formatter.SliceUint8ToString(result.([]uint8))

	return &uuid, nil
}

func (r *repository) FindById(context context.Context, id string) (*models.EventStore, error) {
	result, err := r.Q.Use(r.Event).Where("id", id).FirstCtx(context)

	if nil != err {
		return nil, err
	}

	return result.(*models.EventStore), nil
}