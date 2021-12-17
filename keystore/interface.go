package keystore

import crypto "github.com/libp2p/go-libp2p-core/crypto"

type Interface interface {
	HasKey(id string) (bool, error)

	CreateKey(id string) (crypto.PrivKey, error)

	GetKey(id string) (crypto.PrivKey, error)

	Sign(pubKey crypto.PrivKey, bytes []byte) ([]byte, error)

	Verify(signature []byte, publicKey crypto.PubKey, data []byte) error
}
