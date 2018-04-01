package sdk

import (
	"errors"

	"github.com/elastos/Elastos.ELA.SPV/bloom"
	"github.com/elastos/Elastos.ELA.SPV/p2p"
	"github.com/elastos/Elastos.ELA.SPV/msg"
)

type SPVClient interface {
	SetMessageHandler(SPVMessageHandler)
	Start()
	PeerManager() *p2p.PeerManager
}

type SPVMessageHandler interface {
	OnPeerEstablish(*p2p.Peer)
	OnInventory(*p2p.Peer, *msg.Inventory) error
	OnMerkleBlock(*p2p.Peer, *bloom.MerkleBlock) error
	OnTxn(*p2p.Peer, *msg.Txn) error
	OnNotFound(*p2p.Peer, *msg.NotFound) error
}

func GetSPVClient(netType string, clientId uint64, seeds []string) (SPVClient, error) {
	var magic uint32
	switch netType {
	case TypeMainNet:
		magic = MainNetMagic
	case TypeTestNet:
		magic = TestNetMagic
	default:
		return nil, errors.New("Unknown net type ")
	}
	return NewSPVClientImpl(magic, clientId, seeds)
}