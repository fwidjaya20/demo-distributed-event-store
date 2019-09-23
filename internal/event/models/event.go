package models

import "github.com/payfazz/go-apt/pkg/fazzdb"

// EventStore ...
type EventStore struct {
	fazzdb.Model `json:"-"`
	Id 				string `db:"id" json:"id"`
	Type 			string `db:"type" json:"type"`
	AggregateId 	string `db:"aggregate_id" json:"aggregate_id"`
	AggregateType 	string `db:"aggregate_type" json:"aggregate_type"`
	EventPayload 	string `db:"payload" json:"payload"`
	Channel 		string `db:"channel" json:"channel"`
}

// GeneratePK is a function that used to generate the primary key in table Division table
func (m *EventStore) GeneratePK() {
	m.GenerateId(m)
}

// Payload is a function to read the payload data
func (m *EventStore) Payload() map[string]interface{} {
	return m.MapPayload(m)
}

// Get ...
func (m *EventStore) Get(key string) interface{} {
	return m.Payload()[key]
}

// DivisionModel is function to get User Model
func EventStoreModel() *EventStore {
	return &EventStore{
		Model: fazzdb.UuidModel("event_stores",
			[]fazzdb.Column{
				fazzdb.Col("id"),
				fazzdb.Col("type"),
				fazzdb.Col("aggregate_id"),
				fazzdb.Col("aggregate_type"),
				fazzdb.Col("payload"),
				fazzdb.Col("channel"),
			},
			"id",
			false,
			false,
		),
	}
}