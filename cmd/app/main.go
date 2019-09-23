package main

import (
	"github.com/fwidjaya20/demo-distributed-event-store/cmd/grpc"
	"github.com/fwidjaya20/demo-distributed-event-store/cmd/nats"
	"github.com/fwidjaya20/demo-distributed-event-store/config"
	"github.com/fwidjaya20/demo-distributed-event-store/internal/database/migrations"
	stan "github.com/nats-io/go-nats-streaming"

	//stan "github.com/nats-io/go-nats-streaming"
	"github.com/oklog/oklog/pkg/group"
	"github.com/payfazz/go-apt/pkg/fazzdb"
	"log"
)

func init() {
	initMigration()
}

func main() {
	log.Println("EVENT_STORE")

	var g group.Group

	natsConn := initNATS()

	initGRPC(&g, natsConn)

	log.Fatalln("exit", g.Run())
}

func initNATS() stan.Conn {
	return nats.InitNatsConnection()
}

func initGRPC(g *group.Group, natsConn stan.Conn) {
	grpc.InitGrpcServer(g, natsConn)
}

func initMigration() {
	fazzdb.Migrate(config.GetDb(), "distributed_sys", false, false,
		migrations.Migration1,
	)
}