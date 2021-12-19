// Package accesscontroller defines a default access controller for IPFS Log, it won't actually check anything.
package accesscontroller // import "berty.tech/go-ipfs-p2pdblog/accesscontroller"

import (
	"p2pdb-log/identityprovider"
)

type Default struct {
}

// CanAppend Checks whether a given identity can append an entry to the p2pdblog.
// This implementation allows anyone to write to the p2pdblog.
func (d *Default) CanAppend(LogEntry, identityprovider.Interface, CanAppendAdditionalContext) error {
	return nil
}

var _ Interface = &Default{}
