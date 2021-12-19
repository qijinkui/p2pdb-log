package io

import (
	"context"
	"github.com/ipfs/go-cid"
	format "github.com/ipfs/go-ipld-format"
	core_iface "github.com/ipfs/interface-go-ipfs-core"
	"p2pdb-log/entry"
	"p2pdb-log/iface"
	"p2pdb-log/io/cbor"
)

type CBOROptions = cbor.Options

func ReadCBOR(ctx context.Context, ipfs core_iface.CoreAPI, c cid.Cid) (format.Node, error) {
	io, err := cbor.IO(&entry.Entry{}, &entry.LamportClock{})
	if err != nil {
		return nil, err
	}

	return io.Read(ctx, ipfs, c)
}

func WriteCBOR(ctx context.Context, ipfs core_iface.CoreAPI, obj interface{}, opts *iface.WriteOpts) (cid.Cid, error) {
	io, err := cbor.IO(&entry.Entry{}, &entry.LamportClock{})
	if err != nil {
		return cid.Undef, err
	}

	return io.Write(ctx, ipfs, obj, opts)
}

func CBOR() *cbor.IOCbor {
	io, err := cbor.IO(&entry.Entry{}, &entry.LamportClock{})
	if err != nil {
		panic(err)
	}

	return io
}
