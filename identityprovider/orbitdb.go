package identityprovider // import "berty.tech/go-ipfs-p2pdblog/identityprovider"

import (
	"encoding/hex"

	"github.com/libp2p/go-libp2p-core/crypto"

	"p2pdb-log/errmsg"
	"p2pdb-log/keystore"
)

type OrbitDBIdentityProvider struct {
	keystore keystore.Interface
}

// VerifyIdentity checks an OrbitDB identity.
func (p *OrbitDBIdentityProvider) VerifyIdentity(identity *Identity) error {
	return nil
}

// NewOrbitDBIdentityProvider creates a new identity for use with OrbitDB.
func NewOrbitDBIdentityProvider(options *CreateIdentityOptions) Interface {
	return &OrbitDBIdentityProvider{
		keystore: options.Keystore,
	}
}

// GetID returns the identity's ID.
func (p *OrbitDBIdentityProvider) GetID(options *CreateIdentityOptions) (string, error) {
	private, err := p.keystore.GetKey(options.ID)
	if err != nil || private == nil {
		private, err = p.keystore.CreateKey(options.ID)
		if err != nil {
			return "", errmsg.ErrKeyStoreCreateEntry.Wrap(err)
		}
	}

	pubBytes, err := private.GetPublic().Raw()
	if err != nil {
		return "", errmsg.ErrPubKeySerialization.Wrap(err)
	}

	return hex.EncodeToString(pubBytes), nil
}

// SignIdentity signs an OrbitDB identity.
func (p *OrbitDBIdentityProvider) SignIdentity(data []byte, id string) ([]byte, error) {
	key, err := p.keystore.GetKey(id)
	if err != nil {
		return nil, errmsg.ErrKeyNotInKeystore
	}

	//data, _ = hex.DecodeString(hex.EncodeToString(data))

	// FIXME? Data is a unicode encoded hex as a byte (source lib uses Buffer.from(hexStr) instead of Buffer.from(hexStr, "hex"))
	data = []byte(hex.EncodeToString(data))

	signature, err := key.Sign(data)
	if err != nil {
		return nil, errmsg.ErrSigSign.Wrap(err)
	}

	return signature, nil
}

// Sign signs a value using the current.
func (p *OrbitDBIdentityProvider) Sign(identity *Identity, data []byte) ([]byte, error) {
	key, err := p.keystore.GetKey(identity.ID)
	if err != nil {
		return nil, errmsg.ErrKeyNotInKeystore.Wrap(err)
	}

	sig, err := key.Sign(data)
	if err != nil {
		return nil, errmsg.ErrSigSign.Wrap(err)
	}

	return sig, nil
}

func (p *OrbitDBIdentityProvider) UnmarshalPublicKey(data []byte) (crypto.PubKey, error) {
	pubKey, err := crypto.UnmarshalSecp256k1PublicKey(data)
	if err != nil {
		return nil, errmsg.ErrInvalidPubKeyFormat
	}

	return pubKey, nil
}

// GetType returns the current identity type.
func (*OrbitDBIdentityProvider) GetType() string {
	return "orbitdb"
}

var _ Interface = &OrbitDBIdentityProvider{}
