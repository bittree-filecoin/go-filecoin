package fast

import (
	"context"
	"encoding/json"

	"github.com/ipfs/go-cid"
	"github.com/libp2p/go-libp2p-core/peer"

	"github.com/filecoin-project/go-filecoin/chain"
)

// ChainHead runs the chain head command against the filecoin process.
func (f *Filecoin) ChainHead(ctx context.Context) ([]cid.Cid, error) {
	var out []cid.Cid
	if err := f.RunCmdJSONWithStdin(ctx, nil, &out, "go-filecoin", "chain", "head"); err != nil {
		return nil, err
	}
	return out, nil

}

// ChainLs runs the chain ls command against the filecoin process.
func (f *Filecoin) ChainLs(ctx context.Context) (*json.Decoder, error) {
	return f.RunCmdLDJSONWithStdin(ctx, nil, "go-filecoin", "chain", "ls")
}

// ChainStatus runs the chain status command against the filecoin process.
func (f *Filecoin) ChainStatus(ctx context.Context) (*chain.Status, error) {
	var out *chain.Status
	if err := f.RunCmdJSONWithStdin(ctx, nil, &out, "go-filecoin", "chain", "status"); err != nil {
		return nil, err
	}
	return out, nil
}

// TrustPeer runs the peertrackers trust command against the filecoin process.
func (f *Filecoin) TrustPeer(ctx context.Context, p peer.ID) error {
	if _, err := f.RunCmdWithStdin(ctx, nil, "go-filecoin", "chain", "trust", p.Pretty()); err != nil {
		return err
	}
	return nil
}
