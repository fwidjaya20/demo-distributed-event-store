package nats

import (
	"github.com/fwidjaya20/demo-distributed-event-store/config"
	stan "github.com/nats-io/go-nats-streaming"
	"log"
	"os"
)

func InitNatsConnection() stan.Conn {
	cluster := config.GetEnv(config.CLUSTER_ID)
	clientId := config.GetEnv(config.SERVICE_NAME)
	natsAddr := config.GetEnv(config.NATS_ADDR)

	natsConn, err := stan.Connect(cluster, clientId, stan.NatsURL(natsAddr))

	if nil != err {
		log.Println("transport", "nats", "err", err)
		os.Exit(1)
	}

	log.Println("transport", "nats", "addr", natsAddr)

	return natsConn
}