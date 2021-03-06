package test

import (
	"context"
	"fmt"
	"testing"

	dssync "github.com/ipfs/go-datastore/sync"
	mocknet "github.com/libp2p/go-libp2p/p2p/net/mock"
	"github.com/stretchr/testify/require"
	ipfslog "p2pdb-log"
	idp "p2pdb-log/identityprovider"
	"p2pdb-log/keystore"
)

func BenchmarkAdd(b *testing.B) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	m := mocknet.New(ctx)
	ipfs, closeNode := NewMemoryServices(ctx, b, m)
	defer closeNode()

	datastore := dssync.MutexWrap(NewIdentityDataStore(b))
	ks, err := keystore.NewKeystore(datastore)
	require.NoError(b, err)

	identity, err := idp.CreateIdentity(&idp.CreateIdentityOptions{
		Keystore: ks,
		ID:       "userA",
		Type:     "orbitdb",
	})

	log, err := ipfslog.NewLog(ipfs, identity, &ipfslog.LogOptions{ID: "A"})
	require.NoError(b, err)

	b.ResetTimer()
	// Start the main loop
	for n := 0; n < b.N; n++ {
		_, err = log.Append(ctx, []byte(fmt.Sprintf("%d", n)), nil)
		require.NoError(b, err)
	}
}
