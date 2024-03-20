package main

import (
	"sync"

	"github.com/Dharitri-org/drtg-communication/testscommon"
	"github.com/Dharitri-org/drtg-communication/websocket/data"
	factoryHost "github.com/Dharitri-org/drtg-communication/websocket/factory"
	"github.com/Dharitri-org/drtg-core/marshal/factory"
	logger "github.com/Dharitri-org/drtg-logger"
)

var (
	marshaller, _ = factory.NewMarshalizer("json")
	log           = logger.GetOrCreate("client")
	url           = "localhost:12345"
)

func main() {
	args := factoryHost.ArgsWebSocketHost{
		WebSocketConfig: data.WebSocketConfig{
			URL:                        url,
			Mode:                       data.ModeClient,
			RetryDurationInSec:         1,
			WithAcknowledge:            true,
			BlockingAckOnError:         false,
			DropMessagesIfNoConnection: false,
		},
		Marshaller: marshaller,
		Log:        log,
	}

	wsClient, err := factoryHost.CreateWebSocketHost(args)
	if err != nil {
		log.Error("cannot create WebSocket client", "error", err)
		return
	}

	defer func() {
		_ = wsClient.Close()
	}()

	wg := sync.WaitGroup{}
	wg.Add(1)
	_ = wsClient.SetPayloadHandler(&testscommon.PayloadHandlerStub{
		ProcessPayloadCalled: func(payload []byte, topic string) error {
			log.Info("received", "topic", topic, "payload", string(payload))
			wg.Done()
			return nil
		},
	})

	wg.Wait()
}
