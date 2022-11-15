package audit

import (
	"context"
	"fmt"

	"github.com/pangeacyber/pangea-go/packages/pangea-sdk/internal/signer"
	"github.com/pangeacyber/pangea-go/packages/pangea-sdk/pangea"
)

type Client interface {
	Log(context.Context, Event, bool) (*pangea.PangeaResponse[LogOutput], error)
	Search(context.Context, *SearchInput) (*pangea.PangeaResponse[SearchOutput], error)
	SearchResults(context.Context, *SearchResultInput) (*pangea.PangeaResponse[SearchResultOutput], error)
	Root(context.Context, *RootInput) (*pangea.PangeaResponse[RootOutput], error)
}

type Audit struct {
	*pangea.Client

	SignLogs bool
	Signer   signer.Signer

	VerifyProofs          bool
	SkipEventVerification bool
	rp                    RootsProvider
	lastUnpRootHash       *string
}

func New(cfg *pangea.Config, opts ...Option) (*Audit, error) {
	cli := &Audit{
		Client:                pangea.NewClient("audit", cfg),
		SkipEventVerification: false,
		rp:                    nil,
		lastUnpRootHash:       nil,
	}
	for _, opt := range opts {
		err := opt(cli)
		if err != nil {
			return nil, err
		}
	}
	return cli, nil
}

type Option func(*Audit) error

func WithLogProofVerificationEnabled() Option {
	return func(a *Audit) error {
		a.VerifyProofs = true
		return nil
	}
}

func WithLogSigningEnabled(filename string) Option {
	return func(a *Audit) error {
		a.SignLogs = true
		s, err := signer.NewSignerFromPrivateKeyFile(filename)
		if err != nil {
			return fmt.Errorf("audit: failed signer creation: %w", err)
		}
		a.Signer = s
		return nil
	}
}

func DisableEventVerification() Option {
	return func(a *Audit) error {
		a.SkipEventVerification = true
		return nil
	}
}
