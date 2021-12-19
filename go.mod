module p2pdb-log

go 1.15

//replace berty.tech/go-ipfs-p2pdblog v0.0.0 => ../

replace github.com/golangci/golangci-lint => github.com/golangci/golangci-lint v1.18.0

replace github.com/go-critic/go-critic v0.0.0-20181204210945-ee9bf5809ead => github.com/go-critic/go-critic v0.3.5-0.20190526074819-1df300866540

require (
	berty.tech/go-ipfs-log v1.5.0 // indirect
	github.com/ipfs/go-ipfs v0.11.0 // indirect
	github.com/libp2p/go-libp2p v0.17.0 // indirect
)
