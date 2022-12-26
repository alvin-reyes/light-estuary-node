package core

import "github.com/multiformats/go-multiaddr"

type StorageProviders struct {
	ID        string
	Address   string
	MultiAddr multiaddr.Multiaddr
}
