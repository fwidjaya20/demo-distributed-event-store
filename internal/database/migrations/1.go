package migrations

import (
	"github.com/payfazz/go-apt/pkg/fazzdb"
)

// Migration1 .
var Migration1 = fazzdb.MigrationVersion{
	Tables: []*fazzdb.MigrationTable{
		fazzdb.CreateTable("event_stores", func(table *fazzdb.MigrationTable) {
			table.Field(fazzdb.CreateUuid("id").Primary())
			table.Field(fazzdb.CreateString("type"))
			table.Field(fazzdb.CreateString("aggregate_id"))
			table.Field(fazzdb.CreateString("aggregate_type"))
			table.Field(fazzdb.CreateString("payload"))
			table.Field(fazzdb.CreateString("channel"))
		}),
	},
}