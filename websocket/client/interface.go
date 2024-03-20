package client

import (
	"github.com/Dharitri-org/drtg-communication/websocket"
)

// Transceiver defines what a WebSocket transceiver should be able to do
type Transceiver interface {
	Send(payload []byte, topic string, connection websocket.WSConClient) error
	SetPayloadHandler(handler websocket.PayloadHandler) error
	Listen(connection websocket.WSConClient) (closed bool)
	Close() error
}