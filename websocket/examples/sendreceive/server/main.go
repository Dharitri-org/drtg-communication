package main

import (
	"time"

	"github.com/Dharitri-org/drtg-communication/websocket/data"
	factoryHost "github.com/Dharitri-org/drtg-communication/websocket/factory"
	"github.com/Dharitri-org/drtg-core/marshal/factory"
	logger "github.com/Dharitri-org/drtg-logger"
)

var (
	marshaller, _ = factory.NewMarshalizer("json")
	log           = logger.GetOrCreate("server")
	url           = "localhost:12345"
)

func main() {
	args := factoryHost.ArgsWebSocketHost{
		WebSocketConfig: data.WebSocketConfig{
			URL:                        url,
			Mode:                       data.ModeServer,
			RetryDurationInSec:         1,
			WithAcknowledge:            true,
			BlockingAckOnError:         false,
			DropMessagesIfNoConnection: false,
		},
		Marshaller: marshaller,
		Log:        log,
	}

	wsServer, err := factoryHost.CreateWebSocketHost(args)
	if err != nil {
		log.Error("cannot create WebSocket server", "error", err)
		return
	}

	for {
		err = wsServer.Send([]byte("message"), "test")
		if err == nil {
			break
		}
		time.Sleep(100 * time.Millisecond)
	}
}
