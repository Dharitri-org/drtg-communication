package integrationTests

import (
	"fmt"
	"net"

	"github.com/Dharitri-org/drtg-communication/websocket"
	"github.com/Dharitri-org/drtg-communication/websocket/client"
	hostFactory "github.com/Dharitri-org/drtg-communication/websocket/factory"
	"github.com/Dharitri-org/drtg-communication/websocket/server"
	"github.com/Dharitri-org/drtg-core/core"
	"github.com/Dharitri-org/drtg-core/marshal/factory"
)

const retryDurationInSeconds = 1

var (
	marshaller, _       = factory.NewMarshalizer("gogo protobuf")
	payloadConverter, _ = websocket.NewWebSocketPayloadConverter(marshaller)
)

func createClient(url string, log core.Logger) (hostFactory.FullDuplexHost, error) {
	return client.NewWebSocketClient(client.ArgsWebSocketClient{
		RetryDurationInSeconds:     retryDurationInSeconds,
		WithAcknowledge:            true,
		URL:                        url,
		PayloadConverter:           payloadConverter,
		Log:                        log,
		DropMessagesIfNoConnection: false,
	})
}

func createServer(url string, log core.Logger) (hostFactory.FullDuplexHost, error) {
	return server.NewWebSocketServer(server.ArgsWebSocketServer{
		RetryDurationInSeconds:     retryDurationInSeconds,
		WithAcknowledge:            true,
		URL:                        url,
		PayloadConverter:           payloadConverter,
		Log:                        log,
		DropMessagesIfNoConnection: false,
	})
}

func getFreePort() string {
	// Listen on port 0 to get a free port
	l, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = l.Close()
	}()

	// Get the port number that was assigned
	addr := l.Addr().(*net.TCPAddr)
	return fmt.Sprintf("%d", addr.Port)
}
