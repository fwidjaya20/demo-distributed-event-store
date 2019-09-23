package config

const (
	SERVICE_NAME = "SERVICE_NAME"
	NATS_ADDR = "NATS_ADDR"
	CLUSTER_ID = "CLUSTER_ID"
	CLIENT_ID = "CLIENT_ID"
	GRPC_ADDR = "GRPC_ADDR"
	DB_HOST = "DB_HOST"
	DB_PORT = "DB_PORT"
	DB_NAME = "DB_NAME"
	DB_USER = "DB_USER"
	DB_PASS = "DB_PASS"
)

var defaultConfig = map[string]string {
	SERVICE_NAME:	"EventStore",
	NATS_ADDR: 		"nats://localhost:4222",
	CLUSTER_ID: 	"test-cluster",
	CLIENT_ID: 		"event-store-api",
	GRPC_ADDR:		":5000",
	DB_HOST:		"127.0.0.1",
	DB_PORT:		"5432",
	DB_NAME:		"distributed_sys",
	DB_USER:		"postgres",
	DB_PASS: 		"postgres",
}

func GetEnv(key string) string {
	return defaultConfig[key]
}
