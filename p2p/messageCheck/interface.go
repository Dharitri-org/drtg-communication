package messagecheck

import "github.com/Dharitri-org/drtg-core/core"

type p2pSigner interface {
	Verify(payload []byte, pid core.PeerID, signature []byte) error
	IsInterfaceNil() bool
}
