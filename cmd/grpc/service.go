package grpc

import (
	"github.com/fwidjaya20/demo-distributed-event-store/config"
	"github.com/fwidjaya20/demo-distributed-event-store/internal/event"
	grpcTransport "github.com/fwidjaya20/demo-distributed-event-store/internal/event/transports/grpc"
	pb "github.com/fwidjaya20/demo-distributed-event-store/pkg/protobuf/eventstore"
	"github.com/nats-io/go-nats"

	//stan "github.com/nats-io/go-nats-streaming"
	"github.com/oklog/oklog/pkg/group"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

func InitGrpcServer(g *group.Group, natsConn *nats.Conn) {
	GRPCAddress := config.GetEnv(config.GRPC_ADDR)

	list, err := net.Listen("tcp", GRPCAddress)

	if nil != err {
		log.Println("transport", "gRPC", "during", "Listen", "err", err)
		os.Exit(1)
	}

	g.Add(func() error {
		log.Println("transport", "gRPC", "addr", GRPCAddress)
		s := grpc.NewServer()

		grpcService := grpcTransport.EventGrpcService{
			EventService: event.NewEventService(config.GetDb(), natsConn),
		}

		pb.RegisterEventStoreServer(s, grpcService)

		return s.Serve(list)
	}, func(err error) {
		if err != nil {
			_ = list.Close()
		}
	})
}
