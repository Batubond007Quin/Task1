package api

import (
	"gateway/pkg/api/protobuf"
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var Connection *grpc.ClientConn

func ConnectService() {
	var err error
	Connection, err = grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Client couldn't connect: %s", err)
	}
}

func DisconnectService() {
	Connection.Close()
}

func SendMessageToService(body string) string {
	client := protobuf.NewGatewayServiceClient(Connection)
	message := protobuf.Message{
		Body: body,
	}

	response, err := client.ChangeText(context.Background(), &message)
	if err != nil {
		log.Fatalf("Error when client call ChangeText: %s", err)
	}

	return response.String()
}
